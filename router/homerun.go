package router

import (
	"github.com/gin-gonic/gin"
	"github.com/plainbanana/lapis/app"
	"github.com/plainbanana/lapis/entities"
)

// Discover : GET /discover.json
func Discover(c *gin.Context) {
	device := entities.NewDevice()
	c.BindJSON(&device)

	c.JSON(200, device)
}

// LineupStatus : GET /lineup_status.json
func LineupStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"ScanInProgress": 0,
		"ScanPossible":   0,
		"Source":         "Cable",
		"SourceList":     []string{"Cable"},
	})
}

// GetLineup : GET /lineup.json
func GetLineup(c *gin.Context) {
	lines := app.SetLineup()
	c.JSON(200, lines)
}

// PostLineup : POST /lineup.post
// I am not sure this function is useful
func PostLineup(c *gin.Context) {
	c.JSON(200, gin.H{})
}

// DeviceInfo : GET /device.xml
func DeviceInfo(c *gin.Context) {
	v := entities.NewDeviceXML()
	c.XML(200, v)
}

// ConnectionManager : GET /ConnectionManager.xml
func ConnectionManager(c *gin.Context) {

}

// ContentDirectory : GET /ContentDirectory.xml
func ContentDirectory(c *gin.Context) {

}
