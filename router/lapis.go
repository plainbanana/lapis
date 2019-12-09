package router

import (
	"encoding/base64"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/plainbanana/lapis/entities"
)

// Stream is ffmpeg streamer
func Stream(c *gin.Context) {
	b64URL := c.Param("OriginURL")
	byteURL, err := base64.StdEncoding.DecodeString(b64URL)
	if err != nil {
		log.Println(err)
	}

	ffmpeg := os.Getenv("FFMPEG_BIN")
	if ffmpeg == "" {
		ffmpeg = "ffmpeg"
	}

	input := string(byteURL)
	if test := os.Getenv("TEST_INPUT"); test != "" {
		input = test
	}

	cmd := exec.Command(ffmpeg, "-re", "-analyzeduration", "2MB",
		"-probesize", "2MB",
		"-fix_sub_duration",
		"-i", input, "-c:v", "copy",
		"-acodec", "ac3", "-b:a", "192k",
		"-c:s", "webvtt",
		"-f", "matroska", "pipe:1")

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	log.Println(cmd.String())

	header := c.Writer.Header()
	header["Content-type"] = []string{"video"}
	header["Server"] = []string{"lapis/" + entities.LapisVersion}
	header["Content-Disposition"] = []string{`attachment; filename=out.mkv`}

	reader, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("post process start")
		cmd.Process.Kill()
		gpid, err := syscall.Getpgid(cmd.Process.Pid)
		if err == nil {
			syscall.Kill(-gpid, 15)
		}
		cmd.Wait()
		c.Request.Body.Close()
		log.Println("post process end")
	}()

	io.Copy(c.Writer, reader)
}
