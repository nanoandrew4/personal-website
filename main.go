package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
	"personal-website/internal/fragment/calculator"
	"personal-website/internal/fragment/nsteg"
	"personal-website/internal/handler"
	"personal-website/internal/logging"
	"personal-website/internal/pages"
)

const (
	RFC3339Millis = "2006-01-02T15:04:05.000Z07:00"
)

//go:generate go-localize -input locales -output locales_generated
func main() {
	r := gin.New()
	r.Use(logging.NewGinLogger(), gin.Recovery(), gzip.DefaultHandler().Gin)

	r.Static("/assets", "./assets")

	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.GET("/", gin.WrapH(handler.NewStaticHandler(pages.Home())))
	r.GET("/about-me", gin.WrapH(handler.NewStaticHandler(pages.AboutMe())))
	r.POST("/calculator/operation", gin.WrapH(handler.NewPostRequestHandler(calculator.CalculateOperationResult)))
	r.GET("/calculator", gin.WrapH(handler.NewStaticHandler(pages.CalculatorProject())))
	r.GET("/nsteg", gin.WrapH(handler.NewStaticHandler(pages.NstegProject())))
	r.POST("/nsteg/encode-image", gin.WrapH(handler.NewPostRequestHandler(nsteg.EncodeImage)))
	r.POST("/nsteg/decode-image", gin.WrapH(handler.NewPostRequestHandler(nsteg.DecodeFiles)))

	r.Run(":8080")
}
