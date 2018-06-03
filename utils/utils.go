package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

func DecodeJwtPayload(v string, t interface{}) {
	v = strings.Split(v, ".")[1]
	str, err := base64.RawURLEncoding.DecodeString(v)
	if err != nil {
		GetLogger().PacicError(err)
		return
	}
	json.Unmarshal(str, t)
}

func DecodeJwt(v string) []string {
	var splits = strings.Split(v, ".")
	var jwt []string
	for _, split := range splits {
		str, err := base64.RawURLEncoding.DecodeString(split)

		if err != nil {
			GetLogger().PacicError(err)
			continue
		}

		jwt = append(jwt, string(str))
	}
	return jwt
}
