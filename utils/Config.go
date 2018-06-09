package utils

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

const (
	YMAL = 0
	JSON = 1
	XML  = 2
)

func NewConfig(file string, ctype int, data map[string]interface{}) *Config {
	config := &Config{filepath: file, ctype: ctype, data: data}
	config.load()
	return config
}

type KV struct {
	Key   string
	Value interface{}
}

type Config struct {
	filepath string
	filename string
	ctype    int
	data     map[string]interface{}
	isnew    bool
}

func (config *Config) load() {
	_, err := os.Stat(config.filepath)
	file, ferr := os.OpenFile(config.filepath, os.O_CREATE|os.O_RDWR, 0700)
	defer file.Close()
	if ferr != nil {
		panic(ferr)
	}
	config.filename = file.Name()
	config.isnew = os.IsNotExist(err)
	if config.isnew {
		bd := config.Marshal(config.data)
		file.WriteString(string(bd))
		file.Sync()
		file.Close()
	} else if err == nil {
		buffer := make([]byte, 102400)
		n, err := file.Read(buffer)
		if err != nil {
			panic(err)
		}
		olddata := config.data
		config.data = map[string]interface{}{}
		config.Unmarshal(buffer[:n], config.data)

		for k, v := range olddata {
			if _, has := config.data[k]; !has {
				config.data[k] = v
			}
		}

	}
}

func (config *Config) Get(key string, _default interface{}) interface{} {
	if v, exist := config.data[key]; exist && reflect.TypeOf(v) == reflect.TypeOf(_default) {
		return v
	} else {
		return _default
	}
}

func (config *Config) Set(key string, value interface{}) {
	config.data[key] = value
}

func (config *Config) Save() {
	file, _ := os.OpenFile(config.filepath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0700)
	file.Write(config.Marshal(config.data))
	file.Sync()
	defer file.Close()
}

func (config *Config) IsNew() bool {
	return config.isnew
}

func (config *Config) Marshal(data interface{}) []byte {
	var d []byte
	var err error
	switch config.ctype {
	case YMAL:
		d, err = yaml.Marshal(data)
	case JSON:
		d, err = json.Marshal(data)
	case XML:
		sts := make([]KV, 0)
		for k, v := range data.(map[string]interface{}) {
			sts = append(sts, KV{k, v})
		}
		d, err = xml.MarshalIndent(sts, "", "\n")
	default:
		err = errors.New("Type not found")
	}
	if err != nil {
		GetLogger().Error("Marshal file: " + config.filename + " fail. message: " + err.Error())
	}
	return d
}

func (config *Config) Unmarshal(buffer []byte, out map[string]interface{}) {
	switch config.ctype {
	case YMAL:
		yaml.Unmarshal(buffer, out)
	case JSON:
		json.Unmarshal(buffer, out)
	case XML:
		xml.Unmarshal(buffer, out)
	default:
		GetLogger().Error("Config: " + config.filename + " cannot be parsed ")
	}
}
