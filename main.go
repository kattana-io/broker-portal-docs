package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
)

func main() {

	doc := redoc.Redoc{
		Title:       "Broker Trading API",
		Description: "Katana Broker API Description",
		SpecFile:    "./docs/broker-portal-api.swagger.json",
		SpecPath:    "/swagger.yaml",
		DocsPath:    "/docs",
	}

	e := echo.New()
	e.Use(echoredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(e.Start(":8000"))
}
