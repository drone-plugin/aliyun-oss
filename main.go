package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path"
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
		fmt.Println("target不能以/开头")
		os.Exit(1)
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
	for _, file := range files {
		err = bucket.PutObjectFromFile(path.Join(target, file), path.Join(source, file))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	}
}
