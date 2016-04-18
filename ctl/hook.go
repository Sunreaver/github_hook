package ctl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/sunreaver/gotools/system"

	"github.com/sunreaver/github_hook/mode"

	"gopkg.in/macaron.v1"
)

// Hook 钩子，自动启动一些lua脚本，来处理事宜
func Hook(ctx *macaron.Context) {
	luaStr := ""
	repo := ctx.Params(":url")

	// github_hook 特殊处理
	if repo == "github_hook" {
		req, e := ctx.Req.Body().Bytes()
		if e == nil {
			var form mode.GithubHook
			e = json.Unmarshal(req, &form)
			if e == nil && verification(form) {
				//pull & restart
				luaStr = "./lua/hook.lua"
			} else if e == nil {
				//pull
				luaStr = "./lua/hookpull.lua"
			}
		} else {
			log.Println(e)
		}
	} else {
		data, e := ioutil.ReadFile(system.CurPath() + system.SystemSep() + "github_hook.conf")
		if e == nil {
			conf := mode.MakeConf(data)
			log.Println("Conf : ", conf)
			for _, item := range conf {
				if item.Name == repo {
					luaStr = item.DoLua
					log.Println("Will do: ", luaStr)
					break
				}
			}
		} else {
			log.Println(e)
		}
	}

	go func(c string) {
		if c == "" {
			return
		}
		cmd := exec.Command("nohup", "lua", c, "&!")
		if err := cmd.Run(); err == nil {
			log.Println("OK: " + c)
		} else {
			log.Println("Faile: " + c)
			log.Println(err)
		}
	}(luaStr)

	ctx.JSON(200, map[string]interface{}{
		"status": 200,
		"msg":    luaStr,
	})
}

func verification(v mode.GithubHook) (r bool) {
	r = false
	for _, item := range v.Commits {
		for _, added := range item.Added {
			if strings.HasSuffix(added, ".go") {
				r = true
				log.Println("restart for : " + added)
				return
			}
		}
		for _, removed := range item.Removed {
			if strings.HasSuffix(removed, ".go") {
				r = true
				log.Println("restart for : " + removed)
				return
			}
		}
		for _, modified := range item.Modified {
			if strings.HasSuffix(modified, ".go") {
				r = true
				log.Println("restart for : " + modified)
				return
			}
		}

	}
	return
}
