package services

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

var client *tablestore.TableStoreClient

func Init() {
	endPoint := "http://imp-slipped-away.cn-beijing.ots.aliyuncs.com"
	instanceName := "imp-slipped-away"
	accessKeyId := "LTAIWcpBC4Ce3mz4"
	accessKeySecret := "6o2F4cYfLAD4mS03Lwsw0sno2dhivz"

	client = tablestore.NewClient(endPoint, instanceName, accessKeyId, accessKeySecret)
}