package notionpkg

import (
	"context"
	"fmt"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	"github.com/jomei/notionapi"
)

func ValidateLog(date string) (log LifeLog, valid bool) {
	result, err := getData(date)
	if len(result.Results) == 0 {
		return LifeLog{}, false
	}

	resultProps := result.Results[0].Properties

	filledAtrProp, ok := resultProps[NotionFilledAtr].(*notionapi.FormulaProperty)
	if !ok || filledAtrProp == nil {
		fmt.Println(filledAtrProp, "と", resultProps)
		return LifeLog{}, false
	}
	filledAtr := filledAtrProp.Formula.Boolean

	allowPublishProp, ok := resultProps[NotionAllowPublish].(*notionapi.CheckboxProperty)
	if !ok || allowPublishProp == nil {
		fmt.Println(allowPublishProp, "と", resultProps)
		return LifeLog{}, false
	}
	allowPublish := allowPublishProp.Checkbox

	if err != nil {
		fmt.Println("repository/notion/getLog.goのdatabase.Queryでエラーが発生しました", err)
		return LifeLog{}, false
	}
	if filledAtr && allowPublish {
		log, err = serializeToLogProp(result)
		if err != nil {
			fmt.Println("repository/notion/getLog.goのserializeToLogPropでエラーが発生しました", err)
			return LifeLog{}, false
		}
		return log, true
	} else {
		linepkg.ReplyToUser("バリデーションに失敗しました")
		return LifeLog{}, false
	}
}

func getData(date string) (*notionapi.DatabaseQueryResponse, error) {
	db, err := setLogDB()
	if err != nil {
		return nil, err
	}

	parsedTime := parseTime(date)

	// time.Time 型を notionapiのDate 型へキャスト
	targetDate := notionapi.Date(parsedTime)

	// クエリを作成
	query := createDateQuery(targetDate)

	// クエリを使用してデータを取得
	results, err := client.Database.Query(context.Background(), notionapi.DatabaseID(db.ID), query)
	if err != nil {
		panic(err)
	}
	return results, nil
}

func GetLogData(date string) (log *LifeLog) {
	result, err := getData(date)
	if err != nil {
		fmt.Println("repository/notion/getLog.goのgetDataでエラーが発生しました", err)
		return
	}
	l, err := serializeToLogProp(result)
	if err != nil {
		fmt.Println("repository/notion/getLog.goのserializeToLogPropでエラーが発生しました", err)
		return
	}
	return &l
}
