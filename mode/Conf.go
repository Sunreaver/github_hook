package mode

import (
	"encoding/json"
)

// Conf 配置信息
type Conf struct {
	Name  string `json:"name"`
	DoLua string `json:"dolua"`
}

// MakeConf CreateConf
func MakeConf(data []byte) []Conf {
	var c []Conf
	e := json.Unmarshal(data, &c)
	if e != nil {
		c = []Conf{}
	}
	return c
}
