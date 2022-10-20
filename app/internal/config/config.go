package config

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/rest"
	"io/ioutil"
	"os"
)

type Config struct {
	rest.RestConf
}

type AccountConf struct {
	Accounts []struct {
		Uid    string `json:"uid"`
		Cookie string `json:"cookie"`
	} `json:"accounts"`
}

func GetAccountConf() (*AccountConf, error) {
	f, err := os.Open("./account.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var a AccountConf
	if err := json.Unmarshal(b, &a); err != nil {
		return nil, err
	}
	return &a, err
}

func CheckAccountConf() error {
	_, err := GetAccountConf()
	return err
}
