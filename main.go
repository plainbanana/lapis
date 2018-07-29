package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/plainbanana/lapis/router"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	envLoad()

	r := gin.Default()

	r.GET("/discover.json", router.Discover)
	r.GET("/lineup_status.json", router.LineupStatus)
	r.GET("/lineup.json", router.GetLineup)
	r.GET("/device.xml", router.DeviceInfo)
	r.GET("/ConnectionManager.xml", router.ConnectionManager)
	r.GET("/ContentDirectory.xml", router.ContentDirectory)
	r.POST("/lineup.post", router.PostLineup)

	r.Run(":" + os.Getenv("LAPIS_PORT"))
}
