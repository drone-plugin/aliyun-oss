package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

func BytesToSize(bytes int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	if bytes == 0 {
		return "0 B"
	}
	i := int(math.Floor(math.Log(float64(bytes)) / math.Log(1024)))
	return strconv.FormatFloat(float64(bytes)/math.Pow(1024, float64(i)), 'f', 2, 64) + " " + sizes[i]
}

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
		fileTarget := path.Join(target, file)
		fileSource := path.Join(source, file)
		if strings.HasPrefix(fileTarget, "/") {
			fileTarget = fileTarget[1:]
		}
		// 获取文件大小并转换为直观的大小表示
		fileInfo, err := os.Stat(fileSource)
		if err != nil {
			fmt.Println("Error getting file size:", err)
			os.Exit(1)
		}
		humanReadableSize := BytesToSize(fileInfo.Size())
		fmt.Println("uploading " + fileTarget + " " + strconv.FormatFloat(idx/filesLen, 'f', 2, 64) + "%" + " (" + humanReadableSize + ")")
		err = bucket.PutObjectFromFile(fileTarget, fileSource)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("uploading complete")
}
