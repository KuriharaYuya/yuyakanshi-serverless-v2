package notionpkg

import "github.com/jomei/notionapi"

type DebugLogProperty struct {
	Name         *notionapi.TitleProperty
	AllowPublish *notionapi.CheckboxProperty
}

type LifeLogProperty struct {
	UUID                         *notionapi.FormulaProperty
	FilledAtr                    *notionapi.FormulaProperty
	Title                        *notionapi.TitleProperty
	Date                         *notionapi.DateProperty
	MorningImage                 *notionapi.FilesProperty
	MyFitnessPal                 *notionapi.FilesProperty
	TodayCalorie                 *notionapi.NumberProperty
	ScreenTime                   *notionapi.FilesProperty
	TodayScreenTime              *notionapi.NumberProperty
	MorningActivityTime          *notionapi.DateProperty
	Published                    *notionapi.FormulaProperty
	TweetURL                     *notionapi.URLProperty
	IsDiaryDone                  *notionapi.CheckboxProperty
	IsChatLogDone                *notionapi.CheckboxProperty
	TodayHostsImage              *notionapi.FilesProperty
	MorningActivityEstimatedTime *notionapi.RollupProperty
	MorningActivityLastEdited    *notionapi.RollupProperty
	MorningActPlace              *notionapi.RollupProperty
	MonthlyScreenTime            *notionapi.RollupProperty
	TrainingPageRelation         *notionapi.RelationProperty
	AllowPublish                 *notionapi.CheckboxProperty
	CalenderPicture              *notionapi.FilesProperty
	Memo                         *notionapi.RichTextProperty
}

type LifeLog struct {
	UUID                         string
	FilledAtr                    bool
	Title                        string
	Date                         string
	MorningImage                 string
	MyFitnessPal                 string
	TodayCalorie                 int
	ScreenTime                   string
	TodayScreenTime              int
	MorningActivityTime          string
	Published                    bool
	TweetURL                     string
	IsDiaryDone                  bool
	IsChatLogDone                bool
	TodayHostsImage              string
	MorningActivityEstimatedTime string
	MorningActivityLastEdited    string
	MorningActPlace              string
	MonthlyScreenTime            int
	TrainingPageId               string
	AllowPublish                 bool
	CalenderPicture              string
	Memo                         string
}

const (
	NotionUUID                         = "uuid"
	NotionFilledAtr                    = "filledAtr"
	NotionTitle                        = "title"
	NotionDate                         = "date"
	NotionMorningImage                 = "morningImage"
	NotionMyFitnessPal                 = "myFitnessPal"
	NotionTodayCalorie                 = "todayCalorie"
	NotionScreenTime                   = "screenTime"
	NotionTodayScreenTime              = "todayScreenTime"
	NotionMorningActivityTime          = "morningActivityTime"
	NotionPublished                    = "published"
	NotionTweetURL                     = "tweetUrl"
	NotionIsDiaryDone                  = "isDiaryDone"
	NotionIsChatLogDone                = "isChatLogDone"
	NotionTodayHostsImage              = "todayHostsImage"
	NotionMorningActivityEstimatedTime = "_morningActivityEstimatedTime"
	NotionMorningActivityLastEdited    = "_morningActivityLastEdited"
	NotionMorningActPlace              = "_morningActPlace"
	NotionMonthlyScreenTime            = "_monthlyScreenTime"
	NotionTrainingRelationPage         = "_trainingpageUUID"
	NotionAllowPublish                 = "allowPublish"
	NotionCalenderPicture              = "calenderPicture"
	NotionMemo                         = "memo"
)

const YuyaKanshiPageId = "92a5712217cb4933968c5bfd19ee2c0d"
