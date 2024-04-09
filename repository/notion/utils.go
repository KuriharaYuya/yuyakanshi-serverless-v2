package notionpkg

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	repoutils "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/utils"
	"github.com/jomei/notionapi"
)

func serializeToLogProp(result *notionapi.DatabaseQueryResponse) (LifeLog, error) {
	if len(result.Results) == 0 {
		return LifeLog{}, fmt.Errorf("no results found")
	}

	resultProps := result.Results[0].Properties

	uuidProp, ok := resultProps[NotionUUID].(*notionapi.FormulaProperty)
	if !ok || uuidProp == nil {
		fmt.Println(uuidProp, "と", resultProps)
		return LifeLog{}, fmt.Errorf("failed to cast UUID property")
	}

	filledAtrProp, ok := resultProps[NotionFilledAtr].(*notionapi.FormulaProperty)
	if !ok || filledAtrProp == nil {
		fmt.Println(filledAtrProp, "と", resultProps)
		return LifeLog{}, fmt.Errorf("failed to cast FilledAtr property")
	}

	titleProp, ok := resultProps[NotionTitle].(*notionapi.TitleProperty)
	if !ok || titleProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast Title property")
	}

	dateProp, ok := resultProps[NotionDate].(*notionapi.DateProperty)
	if !ok || dateProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast Date property")
	}

	morningImageProp, ok := resultProps[NotionMorningImage].(*notionapi.FilesProperty)
	if !ok || morningImageProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast MorningImage property")
	}

	myFitnessPalProp, ok := resultProps[NotionMyFitnessPal].(*notionapi.FilesProperty)
	if !ok || myFitnessPalProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast MyFitnessPal property")
	}

	todayCalorieProp, ok := resultProps[NotionTodayCalorie].(*notionapi.NumberProperty)
	if !ok || todayCalorieProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TodayCalorie property")
	}

	screenTimeProp, ok := resultProps[NotionScreenTime].(*notionapi.FilesProperty)
	if !ok || screenTimeProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast ScreenTime property")
	}

	todayScreenTimeProp, ok := resultProps[NotionTodayScreenTime].(*notionapi.NumberProperty)
	if !ok || todayScreenTimeProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TodayScreenTime property")
	}

	morningActivityTimeProp, ok := resultProps[NotionMorningActivityTime].(*notionapi.DateProperty)
	if !ok || morningActivityTimeProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast MorningActivityTime property")
	}

	publishedProp, ok := resultProps[NotionPublished].(*notionapi.FormulaProperty)
	if !ok || publishedProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast Published property")
	}

	tweetURLProp, ok := resultProps[NotionTweetURL].(*notionapi.URLProperty)
	if !ok || tweetURLProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TweetURL property")
	}

	isDiaryDoneProp, ok := resultProps[NotionIsDiaryDone].(*notionapi.CheckboxProperty)
	if !ok || isDiaryDoneProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast IsDiaryDone property")
	}

	isChatLogDoneProp, ok := resultProps[NotionIsChatLogDone].(*notionapi.CheckboxProperty)
	if !ok || isChatLogDoneProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast IsChatLogDone property")
	}

	todayHostsImageProp, ok := resultProps[NotionTodayHostsImage].(*notionapi.FilesProperty)
	if !ok || todayHostsImageProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TodayHostsImage property")
	}

	morningActivityEstimatedTimeProp, ok := resultProps[NotionMorningActivityEstimatedTime].(*notionapi.RollupProperty)
	if !ok || morningActivityEstimatedTimeProp == nil {
		panic("failed to cast MorningActivityEstimatedTime property")
	}

	morningActivityLastEditedProp, ok := resultProps[NotionMorningActivityLastEdited].(*notionapi.RollupProperty)
	if !ok || morningActivityLastEditedProp == nil {
		panic("failed to cast MorningActivityLastEdited property")
	}

	morningActPlaceProp, ok := resultProps[NotionMorningActPlace].(*notionapi.RollupProperty)

	if !ok || morningActPlaceProp == nil {
		panic("failed to cast MorningActPlace property")
	}

	monthlyScreenTimeProp, ok := resultProps[NotionMonthlyScreenTime].(*notionapi.RollupProperty)
	if !ok || monthlyScreenTimeProp == nil {
		panic("failed to cast MonthlyScreenTime property")
	}

	trainingPageRelationProp, ok := resultProps[NotionTrainingRelationPage].(*notionapi.RollupProperty)
	if !ok || trainingPageRelationProp == nil {
		panic("failed to cast TrainingPageRelation property")
	}

	allowPublishProp, ok := resultProps[NotionAllowPublish].(*notionapi.CheckboxProperty)
	if !ok || allowPublishProp == nil {
		panic("failed to cast AllowPublish property")
	}

	calenderPictureProp, ok := resultProps[NotionCalenderPicture].(*notionapi.FilesProperty)
	if !ok || calenderPictureProp == nil {
		panic("failed to cast CalenderPicture property")
	}

	// copy
	memoProp, ok := resultProps[NotionMemo].(*notionapi.RichTextProperty)
	if !ok || memoProp == nil {
		panic("failed to cast Memo property")
	}

	log := LifeLog{
		UUID:                         uuidProp.Formula.String,
		FilledAtr:                    filledAtrProp.Formula.Boolean,
		Title:                        titleProp.Title[0].PlainText,
		Date:                         convertDateFormat(dateProp.Date.Start.String()),
		MorningImage:                 morningImageProp.Files[0].File.URL,
		MyFitnessPal:                 myFitnessPalProp.Files[0].File.URL,
		TodayCalorie:                 int(todayCalorieProp.Number),
		ScreenTime:                   screenTimeProp.Files[0].File.URL,
		TodayScreenTime:              int(todayScreenTimeProp.Number),
		MorningActivityTime:          morningActivityTimeProp.Date.Start.String(),
		Published:                    publishedProp.Formula.Boolean,
		TweetURL:                     tweetURLProp.URL,
		IsDiaryDone:                  isDiaryDoneProp.Checkbox,
		IsChatLogDone:                isChatLogDoneProp.Checkbox,
		TodayHostsImage:              todayHostsImageProp.Files[0].File.URL,
		MorningActivityEstimatedTime: digRollupDate(morningActivityEstimatedTimeProp, rollupTypeDate),
		MorningActivityLastEdited:    digRollupDate(morningActivityLastEditedProp, rollupTypeLastEdt),
		MorningActPlace:              digRollupText(morningActPlaceProp),
		MonthlyScreenTime:            digRollupNumber(monthlyScreenTimeProp),
		TrainingPageId:               digRollupFormulaText(trainingPageRelationProp),
		AllowPublish:                 allowPublishProp.Checkbox,
		CalenderPicture:              calenderPictureProp.Files[0].File.URL,
		Memo:                         memoProp.RichText[0].PlainText,
	}
	return log, nil
}

func parseTime(date string) time.Time {
	// dateを正規表現で検証する。
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !re.MatchString(date) {
		panic("日付の形式が正しくありません" + date)
	}

	// parsedTime, err := time.Parse("2006-01-02", date)
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

func createDateQuery(targetDate notionapi.Date) *notionapi.DatabaseQueryRequest {
	q := &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: NotionDate,
			Date: &notionapi.DateFilterCondition{
				Equals: &targetDate,
			},
		},
	}
	return q
}

func convertDateFormat(dateStr string) string {
	parts := strings.Split(dateStr, "T")
	datePart := parts[0]
	return datePart
}
func convertDateTimeFormat(dateStr string) string {
	// convert "2023-08-14T09:00:00+09:00" to "2023-08-14 09:00"
	parts := strings.Split(dateStr, "T")
	datePart := parts[0]
	timePart := parts[1]
	timePart = strings.Split(timePart, "+")[0]
	// 最後の秒を削除
	timePart = strings.Split(timePart, ":")[0] + ":" + strings.Split(timePart, ":")[1]
	return datePart + " " + timePart
}

// append blocks

// 複数の*notionapi.AppendBlockChildrenRequestを引数として受ける
func AppendContentToPage(pageID string, blocks ...*notionapi.AppendBlockChildrenRequest) {
	client := CreateClient()

	for _, block := range blocks {

		// 指定されたページの子要素として新しいブロックを追加

		res, err := client.Block.AppendChildren(context.Background(), notionapi.BlockID(pageID), block)
		if err != nil {
			fmt.Println(res)
			panic(err)
		}
	}
}

func head3Template(text string) *notionapi.Heading3Block {
	data := notionapi.Heading3Block{
		BasicBlock: notionapi.BasicBlock{
			Object: "block",
			Type:   notionapi.BlockTypeHeading3,
		},
		Heading3: notionapi.Heading{
			RichText: []notionapi.RichText{
				{
					Type: "text", // ここを修正しました
					Text: &notionapi.Text{
						Content: text,
					},
				},
			},
		},
	}

	return &data

}

func pageTemplate(pageID string) *notionapi.LinkToPageBlock {
	return &notionapi.LinkToPageBlock{
		BasicBlock: notionapi.BasicBlock{
			Object: "block",
			Type:   notionapi.BlockTypeLinkToPage,
		},
		LinkToPage: notionapi.LinkToPage{
			Type:       notionapi.BlockType(notionapi.ParentTypePageID),
			DatabaseID: "",
			PageID:     notionapi.PageID(pageID),
		},
	}
}

func quoteTemplate(text string) *notionapi.QuoteBlock {
	return &notionapi.QuoteBlock{
		BasicBlock: notionapi.BasicBlock{
			Object: "block",
			Type:   notionapi.BlockQuote,
		},
		Quote: notionapi.Quote{
			RichText: []notionapi.RichText{
				{
					Type: "text", // ここを修正しました
					Text: &notionapi.Text{
						Content: text,
					},
				},
			},
		},
	}
}

func imageTemplate(lifeLog *LifeLog, imageType string) *notionapi.ImageBlock {
	// constにない名前なら早期リターンする
	if imageType != NotionMorningImage && imageType != NotionScreenTime && imageType != NotionTodayHostsImage && imageType != NotionMyFitnessPal && imageType != NotionCalenderPicture {
		// imageTypeが不正なら、valueをpanicで表示
		panic("imageTypeが不正です" + imageType)
	}

	externalUrl := repoutils.GetImageExternalURl(lifeLog.Date, imageType)
	fmt.Println("externalUrl", externalUrl)
	return &notionapi.ImageBlock{
		BasicBlock: notionapi.BasicBlock{
			Object: "block",
			Type:   notionapi.BlockTypeImage,
		},
		Image: notionapi.Image{
			Type: notionapi.FileTypeExternal,
			External: &notionapi.FileObject{
				URL: externalUrl,
			},
		},
	}
}

const (
	rollupTypeDate    = "date"
	rollupTypeLastEdt = "last_edited_time"
)

func digRollupDate(rollup *notionapi.RollupProperty, rollupType string) string {
	json, _ := json.Marshal(rollup.Rollup.Array[0])
	// date = {"type":"date","date":{"start":"2023-08-14T09:00:00+09:00","end":null}}
	// last_edited_time = {"type":"last_edited_time","last_edited_time":"2023-08-13T18:01:00Z"}
	var dateTime []byte
	if rollupType == rollupTypeDate {

		dateTime = json[32:57]
		regex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)
		if !regex.MatchString(string(dateTime)) {
			panic("日付の形式が正しくありません" + string(dateTime))
		}

	}

	if rollupType == rollupTypeLastEdt {
		dateTime = json[47:66]

		dateTime = append(dateTime, []byte("+09:00")...)
		regex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)
		if !regex.MatchString(string(dateTime)) {
			panic("日付の形式が正しくありません" + string(dateTime))
		}
	}
	date := string(dateTime)
	notionDateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+\d{2}:\d{2}$`)
	if !notionDateRegex.MatchString(date) {
		panic("最終的な日付の形式が正しくありません" + date)
	}

	return date
}

func parseJT(input string) string {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// JST(日本標準時)のロケーションを取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// パースした時間をJSTに変換
	t = t.In(jst)

	// フォーマット変更
	formatted := t.Format("2006/01/02 15:04")
	fmt.Println(formatted) // 2023/08/14 9:00
	return formatted
}

func digRollupText(rollup *notionapi.RollupProperty) string {
	j, err := json.Marshal(rollup.Rollup.Array[0])
	if err != nil {
		panic(err)
	}

	regex := regexp.MustCompile(`"plain_text":"([^"]+)"`)
	match := regex.FindStringSubmatch(string(j))
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func digRollupFormulaText(rollup *notionapi.RollupProperty) string {
	if len(rollup.Rollup.Array) == 0 {
		return ""
	}
	j, err := json.Marshal(rollup.Rollup.Array[0])
	// _trainingpageUUID":{"id":"Mlh%40","type":"rollup","rollup":{"type":"array","array":[{"type":"formula","formula":{"type":"string","string":"82344e07205346949c3881f4126bced7"}}]}},

	if err != nil {
		// error
		fmt.Println("error", err)
		return ""
	}

	regex := regexp.MustCompile(`"string":"([^"]+)"`)
	match := regex.FindStringSubmatch(string(j))
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func digRollupNumber(rollup *notionapi.RollupProperty) int {
	// {"id":"BPAo","type":"rollup","rollup":{"type":"array","array":[{"type":"number","number":30}]}}
	// numberを取り出す
	j, err := json.Marshal(rollup.Rollup.Array[0])
	if err != nil {
		panic(err)
	}

	regex := regexp.MustCompile(`"number":(\d+)`)
	match := regex.FindStringSubmatch(string(j))
	if len(match) > 1 {
		// stringで取れるので、数字に変換して返す
		i, _ := strconv.Atoi(match[1])
		return i
	}

	return 0
}
