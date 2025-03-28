# 阿里云oss上传插件

```yaml
steps:
  - name: upload-to-oss
    image: nmtz/drone-aliyun-oss
    settings:
      ENDPOINT:
        from_secret: endpoint
      ACCESS_KEY_ID:
        from_secret: accessKeyID
      ACCESS_KEY_SECRET:
        from_secret: accessKeySecret
      BUCKET_NAME:
        from_secret: bucketName
      target: /destination/path/
      source: /source/path/*
```
### dev
```bash
go get -u && go mod tidy
#go get -u 会更新所有的包
#tidy 会重新整理 go.mod 文件。
```
