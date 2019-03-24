package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/plainbanana/lapis/entities"
)

// ConvertEpgToXML : GET /api/schedule/programs.json -> epg.xml
func ConvertEpgToXML() *entities.Guide {
	const xmldateformat = "20060102150400 -0700"
	const originalairdateformat = "2006-01-02 15:04:05"

	base := "http://" + os.Getenv("CHINACHU_IP") + ":" + os.Getenv("CHINACHU_PORT") + "/api/schedule/programs.json"
	if os.Getenv("MIRAKURUN_HTTPS") == "true" {
		base = "https://" + os.Getenv("CHINACHU_IP") + ":" + os.Getenv("CHINACHU_PORT") + "/api/schedule/programs.json"
	}

	res, err := http.Get(base) // get program guides as json
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var s entities.Schedules
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&s)

	xml := &entities.Guide{
		GeneratorInfoName: "Chinachu",
	}

	lineups := SetLineup() // for get unique values
	for _, v := range lineups {
		var c entities.ChannelGuide
		c.ID = v.GuideNumber
		c.DisplayName.DisplayName = v.GuideName
		c.DisplayName.Lang = "ja_JP"
		xml.Channel = append(xml.Channel, c)
	}
	for _, v := range s {
		var p entities.ProgrammeGuide
		p.Start = time.Unix(0, v.Start*1000000).Format(xmldateformat)
		p.Stop = time.Unix(0, v.End*1000000).Format(xmldateformat)
		p.Channel = strconv.Itoa(v.Channel.SID)
		p.Category.Category = v.Category
		p.Desc.Desc = v.Detail
		p.Title.Title = v.Title
		p.Category.Lang, p.Desc.Lang, p.Title.Lang = "ja_JP", "ja_JP", "ja_JP"
		// Plex DVR recognize tvseries from whitch programme has episode-num or not
		if v.Category != "movie" {
			p.EpisodeNum.System = "original-air-date"
			p.EpisodeNum.EpisodeNum = time.Unix(0, v.Start*1000000).Format(originalairdateformat)
		}
		xml.Programme = append(xml.Programme, p)
	}

	return xml
}
