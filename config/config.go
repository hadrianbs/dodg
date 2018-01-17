package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PATConfig struct {
	TokenName   string `json:"do_token_name"`
	TokenSecret string `json:"do_token_secret"`
}

type ServiceConfig struct {
	PAT PATConfig
}

func Init() ServiceConfig {
	var pat PATConfig
	var sc ServiceConfig
	if len(os.Getenv("DO_TOKEN_NAME")) == 0 || len(os.Getenv("DO_TOKEN_SECRET")) == 0 {
		rawJson, err := ioutil.ReadFile("/etc/dodg/pat_config.json")
		if err != nil {
			panic(err)
		}
		json.Unmarshal(rawJson, &pat)
	} else {
		pat.TokenName = os.Getenv("DO_TOKEN_NAME")
		pat.TokenSecret = os.Getenv("DO_TOKEN_SECRET")
	}
	sc.PAT = pat
	fmt.Println(sc)
	return sc
}
