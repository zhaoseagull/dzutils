package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var sections = make(map[string]interface{})

func NewSetting(cfgPath, cfgName, cfgType string) (*Setting, error) {
	v := viper.New()
	v.AddConfigPath(cfgPath)
	v.SetConfigName(cfgName)
	v.SetConfigType(cfgType)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("read config error.", err)
		fmt.Println("read config error.", err.Error())
		return nil, err
	}
	return &Setting{vp: v}, nil

}
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

//重新加载
func (s *Setting) RealoadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
