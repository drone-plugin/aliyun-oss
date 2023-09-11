package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func BytesToSize(bytes int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	if bytes == 0 {
		return "0 B"
	}
	i := int(math.Floor(math.Log(float64(bytes)) / math.Log(1024)))
	return strconv.FormatFloat(float64(bytes)/math.Pow(1024, float64(i)), 'f', 2, 64) + " " + sizes[i]
}

func getAllFile(dirPath string, parentPath string) ([]string, error) {
	var allFiles []string
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			subFlies, err := getAllFile(filepath.Join(dirPath, f.Name()), filepath.Join(parentPath, f.Name()))
			if err != nil {
				return nil, err
			}
			allFiles = append(allFiles, subFlies...)
		} else {
			allFiles = append(allFiles, filepath.Join(parentPath, f.Name()))
		}
	}
	return allFiles, nil
}
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func main() {
	endpoint := os.Getenv("PLUGIN_ENDPOINT")
	accessKeyID := os.Getenv("PLUGIN_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("PLUGIN_ACCESS_KEY_SECRET")
	bucketName := os.Getenv("PLUGIN_BUCKET_NAME")
	target := os.Getenv("PLUGIN_TARGET")
	source := os.Getenv("PLUGIN_SOURCE")

	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	exitOnError(err)
	bucket, err := client.Bucket(bucketName)
	exitOnError(err)
	files, err := getAllFile(source, "")
	exitOnError(err)
	filesLen := float64(len(files))
	for index, file := range files {
		percentage := (float64(index+1) / filesLen) * 100
		fileTarget := filepath.Join(target, file)
		fileSource := filepath.Join(source, file)
		// 获取文件大小并转换为直观的大小表示
		fileInfo, err := os.Stat(fileSource)
		exitOnError(err)
		humanReadableSize := BytesToSize(fileInfo.Size())
		fmt.Println("uploading " + fileTarget + " " + strconv.FormatFloat(percentage, 'f', 2, 64) + "%" + " (" + humanReadableSize + ")")
		err = bucket.PutObjectFromFile(fileTarget, fileSource)
		exitOnError(err)
	}
	fmt.Println("uploading complete")
}
