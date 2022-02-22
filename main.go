package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// 機能
// URLを指定
// 5つのgoroutineでダウンロード
// ダウンロード後マージ
// ダウンロード中に一つでもgoroutineでエラーが起きたら終了
// エラーログ

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

	fmt.Println(totalLength)
	fmt.Println()

	size := 10000
	offset := 0
	latest := 0
	n := int(totalLength) / size
	for i := 0; i <= n; i++ {
		offset = size * i
		if (offset + size) > int(totalLength) {
			latest = int(totalLength)
		} else {
			latest = offset + size - 1
		}
		fmt.Println(offset, "-", latest)
		DonwnloadFile("temp"+strconv.Itoa(i), url, offset, latest)
	}

	files := make([]io.Reader, n+1)
	for i := 0; i <= n; i++ {
		file, err := os.Open("temp" + strconv.Itoa(i))

		if os.IsNotExist(err) {
			fmt.Println("file does not exists")
		}

		files[i] = file
	}

	reader := io.MultiReader(files...)
	out, err := os.Create("primazi.jpg")
	if err != nil {

	}
	defer out.Close()
	_, err = io.Copy(out, reader)
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
