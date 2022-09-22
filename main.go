package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path"
	"strconv"
	"strings"
)

func getAllFile(dirPath string, parentPath string) (allFiles []string) {
	files, _ := os.ReadDir(dirPath)
	for _, f := range files {
		if f.IsDir() {
			subFlies := getAllFile(dirPath+string(os.PathSeparator)+f.Name(), parentPath+string(os.PathSeparator)+f.Name())
			allFiles = append(allFiles, subFlies...)
		} else {
			allFiles = append(allFiles, parentPath+string(os.PathSeparator)+f.Name())
		}
	}
	return allFiles
}
func main() {
	endpoint := os.Getenv("PLUGIN_ENDPOINT")
	accessKeyID := os.Getenv("PLUGIN_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("PLUGIN_ACCESS_KEY_SECRET")
	bucketName := os.Getenv("PLUGIN_BUCKET_NAME")
	target := os.Getenv("PLUGIN_TARGET")
	source := os.Getenv("PLUGIN_SOURCE")
	if strings.HasPrefix(target, "/") {
		target = ""
	}

	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	files := getAllFile(source, "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	filesLen := float64(len(files))
	for index, file := range files {
		idx := float64(index+1) * 100
		target = path.Join(target, file)
		if strings.HasPrefix(target, "/") {
			target = target[1:]
		}
		err = bucket.PutObjectFromFile(target, path.Join(source, file))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		float, err := strconv.ParseFloat(fmt.Sprintf("%.2f", idx/filesLen), 64)
		fmt.Println(float, "%", target)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
