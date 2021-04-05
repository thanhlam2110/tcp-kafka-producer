package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thanhlam/tcp-kafka-producer/service"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "KAFKA SERVICE")
	})
	//<---------------------------SSO ------------------------------>
	e.POST("/api/kafka/produce", service.PushKafkaMessage)
	e.Logger.Fatal(e.Start(":1323"))
}
