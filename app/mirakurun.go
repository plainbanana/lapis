package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/plainbanana/lapis/entities"
)

// SetLineup : get ALL channel
func SetLineup() []entities.Lineup {
	base := "http://" + os.Getenv("MIRAKURUN_IP") + ":" + os.Getenv("MIRAKURUN_PORT") + "/api/channels/"
	if os.Getenv("MIRAKURUN_HTTPS") == "true" {
		base = "https://" + os.Getenv("MIRAKURUN_IP") + ":" + os.Getenv("MIRAKURUN_PORT") + "/api/channels/"
	}

	res, err := http.Get(base) // get channel lists as json
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var s entities.Channels
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&s)
	if err != nil {
		log.Println(err)
	}

	var lineups []entities.Lineup
	for _, v := range s {
		for _, vv := range v.Services {
			var items entities.Lineup
			items.GuideName = vv.Name
			items.GuideNumber = strconv.Itoa(vv.ServiceID)
			items.URL = base + v.Type + "/" + v.Channel + "/services/" + strconv.Itoa(vv.ServiceID) + "/stream/"
			lineups = append(lineups, items)
		}
	}

	return lineups
}
