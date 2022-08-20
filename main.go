package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path"
)

func getAllFile(dirPath string) (allFiles []string) {
	var dirs []string
	files, _ := os.ReadDir(dirPath)
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, dirPath+string(os.PathSeparator)+f.Name())
			getAllFile(dirPath + string(os.PathSeparator) + f.Name())
		} else {
			allFiles = append(allFiles, dirPath+string(os.PathSeparator)+f.Name())
		}
	}
	fmt.Println(dirs)
	for _, table := range dirs {
		temp := getAllFile(table)
		for _, temp1 := range temp {
			allFiles = append(allFiles, temp1)
		}
	}
	return allFiles
}
func main() {
	endpoint := "oss-cn-shanghai.aliyuncs.com"
	accessKeyID := "LTAI4FdxJLj9gSDFcAM9SXu2"
	accessKeySecret := "9KIAbO6ma599LDRVoTvStPCVRvJqug"
	bucketName := "zzfzzf"
	target := "example"
	local := "example"
	//endpoint := os.Getenv("PLUGIN_ENDPOINT")
	//accessKeyID := os.Getenv("PLUGIN_ACCESS_KEY_ID")
	//accessKeySecret := os.Getenv("PLUGIN_ACCESS_KEY_SECRET")
	//bucketName := os.Getenv("PLUGIN_BUCKET_NAME")
	//target := os.Getenv("PLUGIN_TARGET")
	//local := os.Getenv("PLUGIN_LOCAL")
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	files := getAllFile(local)
	for _, file := range files {
		fmt.Println("res", path.Join(target, file), file)
		err = bucket.PutObjectFromFile(path.Join(target, file), file)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
