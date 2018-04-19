package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/ghetzel/warp"
	"github.com/ghetzel/warp/daemon"
	"github.com/ghetzel/warp/lib/errors"
	"github.com/ghetzel/warp/lib/logging"
)

var lstFlag string
var prfFlag string
var crtFlag string
var keyFlag string

func init() {
	flag.StringVar(&lstFlag, "listen",
		":4242", "Address to listen on ([ip]:port), default: `:4242`")
	flag.StringVar(&prfFlag, "cpuprofile",
		"", "Enalbe CPU profiling and write to specified file")
	flag.StringVar(&crtFlag, "cert",
		"", "Use the specified cert file to accetpt connections over TLS")
	flag.StringVar(&keyFlag, "key",
		"", "Use the specified key file to accept connections over TLS")

	if fl := log.Flags(); fl&log.Ltime != 0 {
		log.SetFlags(fl | log.Lmicroseconds)
	}
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if prfFlag != "" {
		f, err := os.Create(prfFlag)
		if err != nil {
			log.Fatal(errors.Details(err))
		}

		go func() {
			pprof.StartCPUProfile(f)
			time.Sleep(10 * time.Second)
			pprof.StopCPUProfile()
			log.Fatal("OUT!")
		}()
	}

	ctx := context.Background()

	srv := daemon.NewSrv(
		ctx,
		lstFlag,
		crtFlag,
		keyFlag,
	)

	logging.Logf(ctx, "Started warpd: version=%s", warp.Version)

	err := srv.Run(ctx)
	if err != nil {
		log.Fatal(errors.Details(err))
	}
}
