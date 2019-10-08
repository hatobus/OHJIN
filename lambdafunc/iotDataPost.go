package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hatobus/OHJIN/log"
	model "github.com/hatobus/OHJIN/models"
	"go.uber.org/zap"
)

func StoreIotData(ctx context.Context, data model.Iotdata) error {
	// 識別子を設定しておく (MAC addressとかがよさそう)
	// uniq := hogehoge
	log.Log().Info( /*hoghoge*/ "data posted", zap.Reflect("iotdata", data))
	return nil
}

func main() {
	lambda.Start(StoreIotData)
}
