package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

func main() {
	// fileUrl := "https://assets.st-note.com/production/uploads/images/72599744/rectangle_large_type_2_c2bf9e83a84a4cc7287f1ad827d94847.jpg"

	// if err := DonwnloadFile("primazi.jpg", fileUrl); err != nil {
	// 	panic(err)
	// }
	url := "https://assets.st-note.com/production/uploads/images/72599744/rectangle_large_type_2_c2bf9e83a84a4cc7287f1ad827d94847.jpg"
	// resp, _ := http.Head("https://assets.st-note.com/production/uploads/images/72599744/rectangle_large_type_2_c2bf9e83a84a4cc7287f1ad827d94847.jpg")

	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set("Range", "bytes=0-499")
	// client := new(http.Client)
	// resp, _ := client.Do((req))

	totalLength, err := GetContentLength(url)
	if err != nil {

	}

	size := 10000

	n := int(totalLength) / size
	// var wg sync.WaitGroup
	eg := errgroup.Group{}
	for i := 0; i <= n; i++ {
		i := i
		// eg.Go(MultidDownloadFile(, url, offset, latest))
		eg.Go(func() error {

			offset := 0
			latest := 0

			offset = size * i
			if (offset + size) > int(totalLength) {
				latest = int(totalLength)
			} else {
				latest = offset + size - 1
			}

			filepath := "image/" + strconv.Itoa(i) + ".tmp"

			fmt.Println(i, filepath, offset, latest)
			return DonwnloadFile(filepath, url, offset, latest)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dondload Done.")

}

func ListFiles(dir string) ([]string, error) {

	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && IsTmpFile(info.Name()) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func IsTmpFile(filename string) bool {
	pos := strings.LastIndex(filename, ".")
	if ".tmp" == filename[pos:] {
		return true
	}
	return false
}

func OutputImage(filepath string) {

	//

	// files := make([]io.Reader, n+1)
	// for i := 0; i <= n; i++ {
	// 	file, err := os.Open("temp" + strconv.Itoa(i))

	// 	if os.IsNotExist(err) {
	// 		fmt.Println("file does not exists")
	// 	}

	// 	files[i] = file
	// }

	// reader := io.MultiReader(files...)
	// out, err := os.Create(filepath)
	// if err != nil {

	// }
	// defer out.Close()
	// _, err = io.Copy(out, reader)

}

func MultidDownloadFile(filepath string, url string, offset int, latest int) error {
	// defer wg.Done()
	err := DonwnloadFile(filepath, url, offset, latest)
	return err
}

func DonwnloadFile(filepath string, url string, offset int, latest int) error {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", "bytes="+strconv.Itoa(offset)+"-"+strconv.Itoa(latest))
	client := new(http.Client)
	resp, err := client.Do((req))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func GetContentLength(url string) (int64, error) {
	req, _ := http.NewRequest("HEAD", url, nil)
	client := new(http.Client)
	resp, err := client.Do((req))

	return resp.ContentLength, err
}
