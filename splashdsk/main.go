package main

import (
	"flag"
	"fmt"
	"github.com/vcraescu/go-xrandr"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const (
	resolutionFallback = "1024x768"
)

var (
	query      = "wallpaper"
	resolution = "auto"

	infoLog *log.Logger
)

func init() {
	flag.StringVar(&query, "q", query, "unsplash.com query keyword")
	flag.StringVar(&resolution, "r", resolution, "set resolution manually; automatic resolution detection is attempted otherwise. Fallback value is 1024x768")

	infoLog = log.New(os.Stderr, "INFO: ", 0)
}

func main() {
	log.SetPrefix("splashdsk: ")
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	flag.Parse()

	if resolution == "auto" || resolution == "" {
		resolution = checkResolution()
	}

	url := fmt.Sprintf("https://source.unsplash.com/%s/?%s", resolution, query)
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	imgPath, err := os.CreateTemp("", "*.jpg")
	checkErr(err)
	defer imgPath.Close()
	//defer func() {
	//	name := imgPath.Name()
	//	imgPath.Close()
	//	if err := os.Remove(name); err != nil {
	//		log.Printf("coldn't remove the temporary file %s", name)
	//	}
	//}()

	size, err := io.Copy(imgPath, resp.Body)
	checkErr(err)

	infoLog.Printf("wallpaper with size %d bytes downloaded as %s", size, imgPath.Name())

	fehExe, err := exec.LookPath("feh")
	if err != nil {
		log.Fatal("couldn't find the feh utility: %v", err)
	}

	cmd := exec.Command(fehExe, "--bg-max", imgPath.Name())
	checkErr(cmd.Run())
}

func checkResolution() string {
	screens, err := xrandr.GetScreens()
	if err != nil {
		infoLog.Printf("checkResolution: couldn't get screens: %v", err)

		return resolutionFallback
	}

	if len(screens) != 0 {
		size := screens[0].CurrentResolution
		return fmt.Sprintf("%dx%d", int(size.Width), int(size.Height))
	}

	infoLog.Printf("checkResolution: xrandr couldn't determine the screen's resolution")

	return resolutionFallback
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
