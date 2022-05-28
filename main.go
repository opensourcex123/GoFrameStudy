package main

import (
	_ "goframe-study/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-study/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
