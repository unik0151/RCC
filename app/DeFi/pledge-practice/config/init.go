package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"path"
	"path/filepath"
	"runtime"
)

/*
*
config中全局变量初始化
*/
func init() {
	dir, err := getCallerDir()
	if err != nil {
		panic(err)
		return
	}
	abs, err := filepath.Abs(dir + "/configV1.toml")
	if err != nil {
		panic(err)
		return
	}
	if _, err := toml.DecodeFile(abs, &GlobalConfig); err != nil {
		panic("read toml file error : " + err.Error())
		return
	}

	err = godotenv.Load(dir + "/.env")
	if err != nil {
		panic("load .env file error : " + err.Error())
	}
}

func getCallerDir() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if ok {
		return path.Dir(file), nil
	}
	return "", errors.New("no caller directory")
}
