package notionpkg

import (
	"context"
	"fmt"
	"os"

	utils "github.com/KuriharaYuya/yuya-kanshi-serverless/util"
	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

var client *notionapi.Client

func init() {
	if utils.ENVIRONMENT == "development" {
		// ローカル開発環境用の処理
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("読み込み出来ませんでした_setUp.go: %v", err)
		}
	}
	client = CreateClient()
}

func CreateClient() *notionapi.Client {
	integration_token := os.Getenv("NOTION_API_KEY")
	return notionapi.NewClient(notionapi.Token(integration_token))
}

const (
	debugDbID = "b2c4752a33904be3a434f2c6542a4b75"
	prodDbID  = "8af74dfac9a0482bab353741bb355971"
)

func setLogDB() (db *notionapi.Database, err error) {
	dbId := prodDbID
	db, err = client.Database.Get(context.Background(), notionapi.DatabaseID(dbId))
	if err != nil {
		return nil, err
	}
	return db, err
}
