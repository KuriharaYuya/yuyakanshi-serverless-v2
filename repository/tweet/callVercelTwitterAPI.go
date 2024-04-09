package tweet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
	"github.com/KuriharaYuya/yuya-kanshi-serverless/repository/storage"
	repoutils "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/utils"
)

func CallVercelTwitterAPI(log *notionpkg.LifeLog) string {
	// タイトルを取得
	title := log.Title
	// 日付を取得
	date := log.Date

	uuid := log.UUID
	// ツイート文を生成
	tweetText := generateTweetText(title, uuid)

	// スクリーンタイムの画像を取得
	ScreenTimeS3URL := repoutils.GetImageExternalURl(date, storage.ScreenTime)
	//カレンダーの画像を取得
	CalenderPicS3URL := repoutils.GetImageExternalURl(date, storage.CalenderPic)
	// vercelをcallしてツイートを投稿
	fmt.Println("screenTimeS3URL:", ScreenTimeS3URL)
	tweetId := callVercelAPI(tweetText, ScreenTimeS3URL, log.Memo, CalenderPicS3URL)
	return tweetId

}

func generateTweetText(t string, u string) string {
	baseURL := "https://yuyanki.notion.site/"
	header := "本日の活動報告！🤓"
	url := baseURL + u
	msg := "この仕組みでは、自分をスマホ・ジム・食事・時間の使い方の4点で監視をしています！詳細はこちらから↓"
	msg2 := "下記は、一番サボりやすいスマホの使用時間です！"
	nl := "\n"
	txt := header + nl + t + nl + nl + msg + nl + url + nl + msg2
	return txt
}

const (
	prodApiURL = "https://yuya-kanshi.vercel.app/api/tweet/postTweetFromLambda"
	devApiURL  = "https://fc76-2001-268-c209-8385-b5ef-ea5e-1604-30bc.ngrok-free.app/api/tweet/postTweetFromLambda"
)

func callVercelAPI(t string, s3URL string, memo string, calenderPicURl string) string {
	data := map[string]string{
		"text":           t,
		"screenTimeURl":  s3URL,
		"memo":           memo,
		"calenderPicURl": calenderPicURl,
	}
	// media
	fmt.Println("mediaUrl:", s3URL)

	bodyData, err := json.Marshal(data)
	resForVercel, err := http.Post(prodApiURL, "application/json", bytes.NewBuffer(bodyData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resForVercel.Body.Close()

	body, err := ioutil.ReadAll(resForVercel.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}
	// ex body) {lastTweetId: 1426210000000000000}
	// bodyをjsonに変換
	var resp map[string]string
	json.Unmarshal(body, &resp)
	lastTweetId := resp["lastTweetId"]
	fmt.Println("Response from Vercel:", string(body))
	return lastTweetId
}
