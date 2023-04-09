// +build ali

package oss

type OSSClient struct {
}

func NewOSSClient() *OSSClient {
	return &OSSClient{}
}

func (c *OSSClient) GetBucketName() string {
	return "foo"
}

func (c *OSSClient) GetLatestFile(key string) string {
	return key + "/" + "foo"
}