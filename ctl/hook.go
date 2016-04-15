package ctl

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/sunreaver/github_hook/mode"

	"gopkg.in/macaron.v1"
)

// Hook 钩子，自动启动一些lua脚本，来处理事宜
func Hook(ctx *macaron.Context) {
	luaStr := ""
	repo := ctx.Params(":url")

	if repo == "github_hook" {
		req, e := ctx.Req.Body().Bytes()
		if e == nil {
			var form mode.GithubHook
			e = json.Unmarshal(req, &form)
			if e == nil && verification(form) {
				//更新
				luaStr = "./lua/hook.lua"
			}
		} else {
			log.Println(e)
		}
	} else if repo == "warms" {
		luaStr = "./lua/warm.lua"
	} else if repo == "myserver" {
		luaStr = "./lua/my80server.lua"
	} else {
		ctx.JSON(200, map[string]interface{}{
			"status": 200,
			"msg":    luaStr,
		})
		return
	}

	go func(c string) {
		if c == "" {
			return
		}
		cmd := exec.Command("lua", c)
		if err := cmd.Run(); err == nil {
			log.Println("OK: " + luaStr)
		} else {
			log.Println("Faile: " + luaStr)
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
