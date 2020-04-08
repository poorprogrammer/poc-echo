package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func main() {
	e := echo.New()

	e.GET("/nodes", func(c echo.Context) error {
		strB, _ := json.Marshal(randomString(10))
		return c.JSON(http.StatusOK, strB)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
