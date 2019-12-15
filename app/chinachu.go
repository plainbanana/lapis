package app

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/plainbanana/lapis/entities"
)

// Xmldateformat is xmltv
const Xmldateformat = "20060102150400 -0700"

// Originalairdateformat chinachu format
const Originalairdateformat = "2006-01-02 15:04:05"

// Episodedateformat episodedate
const Episodedateformat = "0102"

func hashMod(value string) int {
	str := fmt.Sprintf("%x", sha256.Sum256([]byte(value)))
	numHash, err := strconv.ParseInt(str[:10], 16, 64)
	if err != nil {
		log.Fatal("Error:", err)
	}
	numReduce := 10000000
	if err != nil {
		log.Fatal("Error:", err)
	}
	fHash := float64(numHash)
	fReduce := float64(numReduce)
	return int(math.Mod(fHash, fReduce))
}

// ConvertEpgToXML : GET /api/schedule/programs.json -> epg.xml
func ConvertEpgToXML() *entities.Guide {

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
		c.Icon.Src = setChannelIconURL(c.ID)
		xml.Channel = append(xml.Channel, c)
	}
	for _, v := range s {
		var p entities.ProgrammeGuide
		p.Start = time.Unix(0, v.Start*1000000).Format(Xmldateformat)
		p.Stop = time.Unix(0, v.End*1000000).Format(Xmldateformat)
		p.Channel = strconv.Itoa(v.Channel.SID)
		p.Category.Category = v.Category
		p.Desc.Desc = v.Detail
		titleSlice := strings.Split(v.Title, "▽")
		p.Title.Title = titleSlice[0]
		if len(titleSlice) >= 2 {
			p.SubTitle.Title = v.SubTitle + " " + titleSlice[1]
		} else {
			p.SubTitle.Title = v.SubTitle
		}
		p.Category.Lang, p.Desc.Lang, p.Title.Lang, p.SubTitle.Lang = "ja_JP", "ja_JP", "ja_JP", "ja_JP"
		// Plex DVR recognize tvseries from whitch programme has episode-num or not
		if v.Category != "cinema" {
			p.EpisodeNum.System = "dd_progid"
			var tail string
			if v.Episode == 0 {
				tail = time.Unix(0, v.Start*1000000).Format(Episodedateformat)
			} else {
				tail = fmt.Sprintf("%04d", v.Episode)
			}
			p.EpisodeNum.EpisodeNum = "EP" + fmt.Sprintf("%08d", hashMod(strings.Split(v.Title, "　")[0])) + "." + tail
			// p.EpisodeNum.System = "original-air-date"
			// p.EpisodeNum.EpisodeNum = time.Unix(0, v.Start*1000000).Format(originalairdateformat)
		}
		xml.Programme = append(xml.Programme, p)
	}

	return xml
}
