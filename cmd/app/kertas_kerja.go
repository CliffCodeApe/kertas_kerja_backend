package app

import (
	"kertas_kerja/config"
	"kertas_kerja/internal/server"
	"kertas_kerja/pkg/token"
)

func Start() {
	config.Load()
	token.Load()
	server.Run()
}
