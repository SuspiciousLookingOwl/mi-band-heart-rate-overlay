package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HeartRateData struct {
	Timestamp time.Time `json:"timestamp"`
	HeartRate int       `json:"heartRate"`
}

type Payload struct {
	HeartRate int `json:"heartRate"`
}

var latestHeartRateData HeartRateData = HeartRateData{
	HeartRate: 0,
	Timestamp: time.Now(),
}

// For logging
// var f, _ = os.OpenFile("heart_rate.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// f.Write([]byte(timestamp + "," + strconv.Itoa(payload.HeartRate) + "\n"))

func logHeartRate(c echo.Context) error {
	payload := new(HeartRateData)
	if err := c.Bind(payload); err != nil {
		return err
	}

	timestamp := time.Now()
	latestHeartRateData = HeartRateData{
		Timestamp: timestamp,
		HeartRate: payload.HeartRate,
	}
	fmt.Printf("[%s] New Heart Rate %d\n", timestamp.Format(time.RFC3339), payload.HeartRate)

	return c.String(http.StatusOK, "")
}

func retrieveHeartRate(c echo.Context) error {
	return c.JSON(http.StatusOK, latestHeartRateData)
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
