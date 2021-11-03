package main

import (
	"context"

	"github.com/kevinanthony/bzen/config"
	"github.com/kevinanthony/bzen/http"
	"github.com/kevinanthony/bzen/http/encoder"
	"github.com/kevinanthony/bzen/rest"
	"github.com/kevinanthony/bzen/server"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	native := http.NewNativeClient()
	factory := encoder.NewFactory()

	client := http.NewClient(native, factory)

	upc := rest.NewGameUPC(cfg, client)

	_, err = upc.GetIDFromUPC(context.Background(), "019962194719")
	if err != nil {
		panic(err)
	}

	server.NewServer().Run()
}
