package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/KuriharaYuya/yuya-kanshi-serverless/gateway"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
	"github.com/KuriharaYuya/yuya-kanshi-serverless/repository/tweet"
	utils "github.com/KuriharaYuya/yuya-kanshi-serverless/util"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req utils.Request) (utils.Response, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	// respをポインタ変数として定義
	var resp *utils.Response
	go func() {
		resp = gateway.Gateway(req)
		wg.Done()
	}()

	wg.Wait()
	return *resp, nil
}

func init() {
	if utils.ENVIRONMENT == "development" {
		// ローカル開発環境用の処理
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("読み込み出来ませんでした_main.go: %v", err)
		}
	} else {
		// 本番環境用の処理
		// 例: AWS Secrets Managerから秘密情報を取得する
	}

}

func main() {
	debugMode := os.Getenv("DEBUG_MODE")
	if debugMode == "true" {
		// run some tgt funcs below
		fmt.Println("debug mode")
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			l, _ := notionpkg.ValidateLog("2023-11-03")
			tweet.CallVercelTwitterAPI(&l)
			defer wg.Done()
			// リクエストをを全てunmarshalして表示する

		}()
		wg.Wait()

	} else {
		lambda.Start(Handler)
	}

}
