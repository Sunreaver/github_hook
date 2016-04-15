package ctl

import (
	"log"
	"os/exec"

	"gopkg.in/macaron.v1"
)

// Hook 钩子，自动启动一些lua脚本，来处理事宜
func Hook(ctx *macaron.Context) {
	luaStr := ""
	repo := ctx.Params(":url")
	if repo == "" {
		ctx.JSON(200, map[string]interface{}{
			"status": 200,
			"msg":    "bad url",
		})
		return
	}
	if repo == "github_hook" {
		luaStr += "./lua/hook.lua"
	}

	go func(c string) {
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
