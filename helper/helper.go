package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
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

func CheckUpdate() {
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
			fmt.Scanln(&a)
			if a == "y" || a == "Y" || a == "yes" {
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

func DownloadFile(url, filename string) error {
	fmt.Println("downloading...")
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("network error: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("download complete: " + filename)
	return nil
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
			time.Sleep(1 * time.Second)
			fmt.Println("Your browser will start in 3 seconds")
			time.Sleep(3 * time.Second)
			exec.Command(run, url).Start()
		}()
	}
}
