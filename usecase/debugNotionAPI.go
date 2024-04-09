package usecase

import (
	"fmt"
	"strings"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
)

func DebugNotionAPI(msg string) {
	// msgに"publish=true"が含まれているか確認する
	var publish bool
	publish = true
	publish = strings.Contains(msg, "publish=true") && !strings.Contains(msg, "publish=false")

	// データを取得
	data, err := notionpkg.GetDebugData(publish)
	if err != nil {
		linepkg.ReplyToUser("データの取得に失敗しました")
		fmt.Println("データの取得に失敗しました", err)
		return
	}

	// データを整形し、文字列に格納する

	nameStr := "Name: " + data.Name.Title[0].Text.Content + "\n"
	// convert bool(data.AllowPublish.Checkbox) to str
	allow := data.AllowPublish.Checkbox

	allowPublishStr := "\n" + "allowPublish: " + fmt.Sprintf("%t", allow) + "\n"

	replyMsg := nameStr + allowPublishStr

	// 返信する
	linepkg.ReplyToUser(replyMsg)
}
