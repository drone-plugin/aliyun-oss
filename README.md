# 阿里云oss上传插件

```yaml
  - name: oss
    image: registry.cn-shanghai.aliyuncs.com/zzf2001/drone-aliyun-oss
    settings:
      endpoint:
        from_secret: endpoint
      ACCESS_KEY_ID:
        from_secret: accessKeyID
      ACCESS_KEY_SECRET:
        from_secret: accessKeySecret
      BUCKET_NAME:
        from_secret: bucketName
      target: dist
      source: dist
```
### dev
```bash
go get -u && go mod tidy
#go get -u 会更新所有的包
#tidy 会重新整理 go.mod 文件。
```