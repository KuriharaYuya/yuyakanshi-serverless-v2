package linepkg

import (
	"fmt"
	"log"
	"os"

	utils "github.com/KuriharaYuya/yuya-kanshi-serverless/util"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func ReplyToUser(cnt string) {
	if utils.ENVIRONMENT == "development" {
		// ローカル開発環境用の処理
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("読み込み出来ませんでした_line.go: %v", err)
		}
	} else {
		// 本番環境用の処理
		// 例: AWS Secrets Managerから秘密情報を取得する
	}

	secret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	bot, err := linebot.New(secret, channelToken)
	if err != nil {
		fmt.Println(err)
		return

	}
	log.Println("Reply")
	bot.BroadcastMessage(linebot.NewTextMessage(cnt)).Do()
	return
}

func Announce(cnt string) {
	if utils.ENVIRONMENT == "development" {
		// ローカル開発環境用の処理
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("読み込み出来ませんでした_line.go: %v", err)
		}
	} else {
		// 本番環境用の処理
		// 例: AWS Secrets Managerから秘密情報を取得する
	}

	secret := os.Getenv("LINE_NOTIFY_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_NOTIFY_CHANNEL_ACCESS_TOKEN")

	bot, err := linebot.New(secret, channelToken)
	if err != nil {
		fmt.Println(err)
		return

	}
	log.Println("Announce")
	lineGroupID_RION := os.Getenv("LINE_GROUP_ID_RION")
	lineGroupID_NIKI := os.Getenv("LINE_GROUP_ID_NIKI")
	lineGroupID_AMANE := os.Getenv("LINE_GROUP_ID_AMANE")
	bot.PushMessage(lineGroupID_RION, linebot.NewTextMessage(cnt)).Do()
	bot.PushMessage(lineGroupID_NIKI, linebot.NewTextMessage(cnt)).Do()
	bot.PushMessage(lineGroupID_AMANE, linebot.NewTextMessage(cnt)).Do()
	return
}
