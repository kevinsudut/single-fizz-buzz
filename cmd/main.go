package main

import (
	"errors"
	"net/http"

	"github.com/kevinsudut/single-fizz-buzz/app"
	"github.com/kevinsudut/single-fizz-buzz/pkg/lib/log"
)

func main() {
	log.Init()

	err := app.Init()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("app.Init", err)
	}
}
