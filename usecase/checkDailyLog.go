package usecase

import (
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
)

func CheckDailyLog(date string) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		logData, valid := notionpkg.ValidateLog(date)
		if !valid {
			return
		}

		title := logData.Title
		tweetURL := logData.TweetURL
		textData := "タイトル: " + title + "\n" + "ツイートURL: " + tweetURL
		linepkg.ReplyToUser(textData)
		wg.Done()
	}()
	wg.Wait()
	return

}
