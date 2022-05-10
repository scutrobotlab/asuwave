package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
)

func PrintDownloadPercent(done chan int64, path string, total int64) {
	var stop bool = false
	for {
		select {
		case <-done:
			stop = true
		default:
			file, err := os.Open(path)
			if err != nil {
				glog.Fatalln(err.Error())
			}
			fi, err := file.Stat()
			if err != nil {
				glog.Fatalln(err.Error())
			}
			size := fi.Size()
			if size == 0 {
				size = 1
			}
			var percent float64 = float64(size) / float64(total) * 100
			fmt.Printf("%.0f", percent)
			fmt.Println("%")
		}
		if stop {
			break
		}
		time.Sleep(time.Second)
	}
}

func DownloadFile(url, filename string) error {
	fmt.Println("downloading...")

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	headResp, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		panic(err)
	}

	done := make(chan int64)
	go PrintDownloadPercent(done, filename, int64(size))

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("network error: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	n, err := io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	done <- n

	fmt.Println("download complete: " + filename)
	return nil
}
