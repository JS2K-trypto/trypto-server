package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

// config.toml 파일 기준으로 구조체 선언
type Work struct {
	Name     string
	Desc     string
	Excute   string
	Duration int
	Args     string
}

type Config struct {
	Server struct {
		Mode string
		Port string
	}

	DB map[string]map[string]interface{}

	Work []Work

	Log struct {
		Level   string
		Fpath   string
		Msize   int
		Mage    int
		Mbackup int
	}
}

// 파일 오픈
func GetConfig(fpath string) *Config {
	c := new(Config)

	//fmt.Println("c", c)
	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		//toml 파일 디코딩
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			fmt.Println(c)
			return c
		}

	}
}
