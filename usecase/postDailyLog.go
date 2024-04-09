package usecase

import (
	"encoding/json"
	"fmt"
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
	"github.com/KuriharaYuya/yuya-kanshi-serverless/repository/storage"
	"github.com/KuriharaYuya/yuya-kanshi-serverless/repository/tweet"
)

func PostDailyLog(date string) {
	s := storage.SetUp()

	wg := sync.WaitGroup{}
	wg.Add(1)
	var msg string
	linepkg.ReplyToUser(msg)
	go func() {
		defer wg.Done()

		l, valid := notionpkg.ValidateLog(date)
		marshaled, _ := json.Marshal(l)
		msg = string(marshaled)
		if valid {
			linepkg.ReplyToUser("バリデーションを通過しました")
		} else {
			linepkg.ReplyToUser("バリデーションに失敗しました")
			linepkg.ReplyToUser(msg)
			return
		}

		storage.UploadImages(&l, s)
		fmt.Println(msg)
		linepkg.ReplyToUser("画像のアップロードが完了しました")

		s := notionpkg.DiaryHeaderTemplate(&l)

		m := notionpkg.MorningTemplate(&l)
		d := notionpkg.DeviceTemplate(&l)
		h := notionpkg.HealthTemplate(&l)
		notionpkg.AppendContentToPage(l.UUID, &s, &m, &d, &h)
		linepkg.ReplyToUser("Notionへの書き込みが完了しました")

		// twitter
		tweetID := tweet.CallVercelTwitterAPI(&l)
		if tweetID == "" {
			linepkg.ReplyToUser("Twitterへの投稿に失敗しました")
			return
		}
		linepkg.Announce("Twitterへの投稿が完了しました" + "\n" + "https://twitter.com/kurihara_poni3/status/" + tweetID)
		linepkg.ReplyToUser("Twitterへの投稿が完了しました" + "\n" + tweetID)

	}()

	wg.Wait() // この行でgoroutineが完了するのを待ちます。
}
