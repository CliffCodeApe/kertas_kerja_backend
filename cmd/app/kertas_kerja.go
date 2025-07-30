package app

import (
	"kertas_kerja/config"
	"kertas_kerja/internal/server"
)

func Start() {
	config.Load()
	server.Run()
}
