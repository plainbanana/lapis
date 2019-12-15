package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/plainbanana/lapis/app"
	"github.com/plainbanana/lapis/entities"
	"github.com/plainbanana/lapis/router"
)

func init() {
	if os.Getenv("DOTENV") == "true" {
		envLoad()
	}

	app.MirakurunBase = "http://" + os.Getenv("MIRAKURUN_IP") + ":" + os.Getenv("MIRAKURUN_PORT")
	if os.Getenv("MIRAKURUN_HTTPS") == "true" {
		app.MirakurunBase = "https://" + os.Getenv("MIRAKURUN_IP") + ":" + os.Getenv("MIRAKURUN_PORT")
	}
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/discover.json", router.Discover)
	r.GET("/lineup_status.json", router.LineupStatus)
	r.GET("/lineup.json", router.GetLineup)
	r.GET("/device.xml", router.DeviceInfo)
	r.GET("/ConnectionManager.xml", router.ConnectionManager)
	r.GET("/ContentDirectory.xml", router.ContentDirectory)
	r.GET("/epg.xml", router.EPG)
	r.POST("/lineup.post", router.PostLineup)
	r.GET("/stream/:OriginURL", router.Stream)

	entities.LapisPort = os.Getenv("LAPIS_PORT")
	if entities.LapisPort == "" {
		entities.LapisPort = "8080"
	}

	r.Run(":" + entities.LapisPort)
}
