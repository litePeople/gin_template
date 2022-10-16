package alioss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func Bucket() (bucket *oss.Bucket, err error) {
	return alioss.cli.Bucket(alioss.Bucket)
}
