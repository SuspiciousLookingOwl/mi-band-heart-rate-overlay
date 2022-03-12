package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var latestHeartRate int = 0

// For logging
// var f, _ = os.OpenFile("heart_rate.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// f.Write([]byte(timestamp + "," + strconv.Itoa(payload.HeartRate) + "\n"))

type Payload struct {
	HeartRate int `json:"heartRate"`
}

func logHeartRate(c echo.Context) error {
	payload := new(Payload)
	if err := c.Bind(payload); err != nil {
		return err
	}

	timestamp := time.Now().Format(time.RFC3339)
	latestHeartRate = payload.HeartRate
	fmt.Printf("[%s] New Heart Rate %d\n", timestamp, payload.HeartRate)

	return c.String(http.StatusOK, "")
}

func retrieveHeartRate(c echo.Context) error {
	return c.JSON(http.StatusOK, Payload{
		HeartRate: latestHeartRate,
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.HideBanner = true

	e.Static("/", "static")
	e.POST("/heart-rate", logHeartRate)
	e.GET("/heart-rate", retrieveHeartRate)

	e.Logger.Fatal(e.Start(":27893"))
}
