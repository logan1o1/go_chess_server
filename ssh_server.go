package main

import (
	"charm.land/log/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	"charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"
	"github.com/logan1o1/go_chess_server/config"
	"github.com/logan1o1/go_chess_server/interfaces"
)

func InitSshServer(db interfaces.IDatabase) *ssh.Server {
	app, err := wish.NewServer(
		wish.WithAddress(":"+config.EnvVars.AppPort),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatal("Unable to initialize ssh server")
	}

	return app
}
