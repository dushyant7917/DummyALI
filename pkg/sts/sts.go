// +build ali

package sts

type STSClient struct {
}

func NewSTSClient() *STSClient {
	return &STSClient{}
}

func (c *STSClient) ConstructRequest() string {
	return "foo"
}