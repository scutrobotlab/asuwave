package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var (
	Port      int
	GitTag    string
	GitHash   string
	BuildTime string
	GoVersion string
)

type githubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func init() {
	if _, err := os.Stat(AppConfigDir()); os.IsNotExist(err) {
		err := os.MkdirAll(AppConfigDir(), 0755)
		if err != nil {
			panic(err)
		}
	}
}

func AppConfigDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "./"
	}

	return path.Join(dir, "asuwave")
}

func GetVersion() string {
	return fmt.Sprintf("asuwave %s\nbuild time %s\n%s", GitHash, BuildTime, GoVersion)
}

func CheckUpdate(auto bool) {
	resp, err := http.Get("https://api.github.com/repos/scutrobotlab/asuwave/releases/latest")
	if err != nil {
		fmt.Println("network error: " + err.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var gr githubRelease
	if err := json.Unmarshal([]byte(body), &gr); err != nil {
		return
	}
	if GitTag == gr.TagName {
		fmt.Println("already the latest version: " + GitTag)
		return
	}
	for _, asset := range gr.Assets {
		if strings.Contains(asset.Name, runtime.GOOS+"_"+runtime.GOARCH) {
			fmt.Println("current version is " + GitTag)
			fmt.Println("new version available: " + gr.TagName)
			fmt.Print("download now? (y/n) ")
			var a string
			if !auto {
				fmt.Scanln(&a)
			}
			if auto || a == "y" || a == "Y" || a == "yes" {
				if err := DownloadFile(asset.BrowserDownloadURL, asset.Name); err != nil {
					fmt.Println("download error: " + err.Error())
					fmt.Println("trying hub.fastgit.org...")
					asset.BrowserDownloadURL = strings.Replace(asset.BrowserDownloadURL, "https://github.com", "https://hub.fastgit.org", 1)
					DownloadFile(asset.BrowserDownloadURL, asset.Name)
				}
			}
			return
		}
	}
	fmt.Printf("don't know your platform: %s, %s", runtime.GOOS, runtime.GOARCH)
}

func StartBrowser(url string) {
	var commands = map[string]string{
		"windows": "explorer.exe",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	run, ok := commands[runtime.GOOS]
	if !ok {
		fmt.Printf("don't know how to open things on %s platform", runtime.GOOS)
	} else {
		go func() {
			fmt.Println("Your browser will start")
			exec.Command(run, url).Start()
		}()
	}
}
