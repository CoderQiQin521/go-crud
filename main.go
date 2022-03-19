package main

import (
	_ "goframe-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"goframe-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
