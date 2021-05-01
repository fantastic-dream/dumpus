package main

import (
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"github.com/pengcheng789/dumpus/daemon"
	"github.com/pengcheng789/dumpus/log"
	"github.com/pkg/errors"
	"os"
)

const (
	helpMessage string = `
Dumpus, a collector for system performance information.
Usage: dumpus [-irh]

Options:
`
)

var (
	logger = log.Logger

	install bool
	remove  bool
	help    bool
)

func init() {
	flag.BoolVar(&install, "i", false, "Install service.")
	flag.BoolVar(&remove, "r", false, "Remove service.")
	flag.BoolVar(&help, "h", false, "This help.")

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, helpMessage)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
	}

	if install {
		logger.Info("Installing service.")
		s, err := daemon.NewService()
		if err != nil {
			logger.Fatal(err)
		}

		err = s.Install()
		if err != nil {
			logger.Fatal(errors.Wrap(err, "Install service failure."))
		}
		logger.Info("Install service finished.")

		logger.Info("Staring service.")
		err = s.Start()
		if err != nil {
			logger.Fatal(errors.Wrap(err, "Start service failure."))
		}
		logger.Info("Start service finished.")
	}

	if remove {
		s, err := daemon.NewService()
		if err != nil {
			logger.Fatal(err)
		}

		status, err := s.Status()
		if err != nil {
			logger.Fatal(errors.Wrap(err, "Get service status failure, remove service failure."))
		}

		if status == service.StatusRunning {
			logger.Info("Stopping service.")
			err = s.Stop()
			if err != nil {
				logger.Fatal(errors.Wrap(err, "Stop service failure, remove service failure."))
			}
			logger.Info("Stop service finished.")
		}

		logger.Info("Removing service.")
		err = s.Uninstall()
		if err != nil {
			logger.Fatal(errors.Wrap(err, "Remove service failure."))
		}
		logger.Info("Remove service finished.")
	}

	if !help && !install && !remove {
		s, err := daemon.NewService()
		if err != nil {
			logger.Fatal(err)
		}

		err = s.Run()
		if err != nil {
			logger.Fatal(err)
		}
	}
}
