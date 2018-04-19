package main

import (
	"os"

	"github.com/ghetzel/warp/client"
	_ "github.com/ghetzel/warp/client/command"
	"github.com/ghetzel/warp/lib/out"
)

func main() {
	cli, err := cli.New(os.Args[1:])
	if err != nil {
		out.Errof("[Error] %s\n", err.Error())
	}

	err = cli.Run()
	if err != nil {
		out.Errof("[Error] %s\n", err.Error())
	}
}
