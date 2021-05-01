package daemon

import (
	"github.com/kardianos/service"
	"github.com/pengcheng789/dumpus/conf"
	"github.com/pengcheng789/dumpus/log"
	"github.com/pengcheng789/dumpus/trans"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"syscall"
)

var logger = log.Logger

type program struct {
	exit chan os.Signal
}

func (p *program) Start(s service.Service) error {
	logger.Info("Service starting")
	p.exit = make(chan os.Signal, 1)
	signal.Notify(p.exit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go p.run()
	logger.Info("Service started.")
	return nil
}

func (p *program) run() {
	conf.Init()
	trans.Init(p.exit)

	<-p.exit
}

func (p *program) Stop(s service.Service) error {
	logger.Info("Service stop.")
	close(p.exit)

	return nil
}

func NewService() (service.Service, error) {
	svcConfig := &service.Config{
		Name:        "Dumpus",
		DisplayName: "Dumbo Octopus",
		Description: "A collector for system performance information.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Create service failure.")
	}

	return s, nil
}
