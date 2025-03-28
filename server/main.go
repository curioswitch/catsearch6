// Copyright (c) Choko (choko@curioswitch.org)
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"github.com/curioswitch/go-curiostack/server"

	"github.com/curioswitch/catsearch/server/internal/config"
)

var confFiles embed.FS // Currently empty

func main() {
	os.Exit(server.Main(&config.Config{}, confFiles, setupServer))
}

func setupServer(ctx context.Context, _ *config.Config, s *server.Server) error {
	mux := server.Mux(s)

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("I'm a cat!"))
	})

	return server.Start(ctx, s)
}
