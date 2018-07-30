package router

import (
	"github.com/gin-gonic/gin"
	"github.com/plainbanana/lapis/app"
)

// EPG : GET /epg.xml
func EPG(c *gin.Context) {
	body := app.ConvertEpgToXML()
	c.XML(200, body)
}
