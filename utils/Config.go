package utils

import (
	"os"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"encoding/xml"
	"reflect"
"errors"
)

const (
	YMAL = 0
	JSON = 1
	XML = 2
)

func NewConfig(filename string, ctype int, data map[string]interface{}) *Config {
	config := &Config{filepath:filename, ctype:ctype, data:data}
	config.load()
	return config
}

type Config struct {
	filepath string
	ctype int
	data map[string]interface{}
	file *os.File
	isnew bool
}

func (config *Config) load() {
	_, err := os.Stat(config.filepath)
	file, ferr := os.OpenFile(config.filepath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0700)
	if ferr != nil {
		panic(ferr)
	}
	config.file = file
	config.isnew = os.IsNotExist(err)
	if config.isnew {
		bd := config.Marshal(config.data)
		file.WriteString(string(bd))
		file.Sync()
		file.Close()
	}else if err == nil{
		buffer := make([]byte, 102400)
		file.Read(buffer)
		config.Unmarshal(buffer, config.data)
	}
}

func (config *Config) Get(key string, _default interface{}) interface{} {
	if v, exist := config.data[key]; exist && reflect.TypeOf(v) == reflect.TypeOf(_default) {
		return v
	}else {
		return _default
	}
}

func (config *Config) Set(key string, value interface{}) {
	config.data[key] = value
}

func (config *Config) Save() {
	config.file.Truncate(0)
	config.file.Write(config.Marshal(config.data))
	config.file.Sync()
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
		d, err = xml.Marshal(data)
	default:
		err = errors.New("Type not found")
	}
if err != nil {
GetLogger().Error("Marshal file: " + config.file.Name() + " fail. message: " + err.Error())
}
	return d
}

func (config *Config) Unmarshal(buffer []byte, out interface{}) {
	switch config.ctype {
	case YMAL:
		yaml.Unmarshal(buffer, out)
	case JSON:
		json.Unmarshal(buffer, out)
	case XML:
		xml.Unmarshal(buffer, out)
	default:
		GetLogger().Error("Config: " + config.file.Name() + " cannot be parsed ")
	}
}