package main_test

import (
	"testing"

	"github.com/godzillaframework/godzilla"
)

func maintest(t *testing.T) {
	gz := godzilla.New()

	gz.Get("/test", func(ctx godzilla.Context) {
		ctx.SendString("Test passed")
	})

	gz.Start(":1234")
}
