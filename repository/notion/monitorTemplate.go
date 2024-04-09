package notionpkg

import (
	"strconv"

	"github.com/jomei/notionapi"
)

func MorningTemplate(log *LifeLog) notionapi.AppendBlockChildrenRequest {
	// 前日時点での目標
	tgtDateTimeTitle := head3Template("朝活開始時刻")
	tgtDateTime := quoteTemplate(convertDateTimeFormat(log.MorningActivityTime))

	// 目標時刻
	tgtDateTimeEstTitle := head3Template("前日時点での目標")
	tgtDateTimeEst := quoteTemplate(convertDateTimeFormat(log.MorningActivityEstimatedTime))

	// 場所
	placeTitle := head3Template("場所")
	place := quoteTemplate(log.MorningActPlace)

	// 目標設定のタイミング
	tgtDateLastEdTitle := head3Template("目標設定のタイミング")
	tgtDateLastEd := quoteTemplate(convertDateTimeFormat(log.MorningActivityLastEdited))

	// 証明写真
	photoTitle := head3Template("証明写真")
	photoMsg := quoteTemplate("レシートや時計の時刻が、朝活開始時刻と一致していることを確認してください。")
	photo := imageTemplate(log, NotionMorningImage)

	newBlocks := notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			tgtDateTimeTitle,
			tgtDateTime,
			tgtDateTimeEstTitle,
			tgtDateTimeEst,
			tgtDateLastEdTitle,
			tgtDateLastEd,
			placeTitle,
			place,
			photoTitle,
			photoMsg,
			photo,
		},
	}

	newToggleBlock := toggleTemplate("朝活", &newBlocks)

	return newToggleBlock
}

func DeviceTemplate(log *LifeLog) notionapi.AppendBlockChildrenRequest {
	// 	###スマホ
	spTitle := head3Template("スマホ")

	// 一日の目標時間 h3
	spTgtTimeTitle := head3Template("一日の目標時間")
	// 一日の目標時間 quote
	spTgtTime := quoteTemplate(strconv.Itoa(log.MonthlyScreenTime))

	// スクリーンタイムの画像 img
	spImg := imageTemplate(log, NotionScreenTime)

	// ### pc
	pcTitle := head3Template("pc")

	// 説明 toggle
	pcExpQuote := quoteTemplate(" pcのhostsファイルを用いて、pcでsnsをブロックしています。 \n 最後にhostsファイルが編集された時刻と現在の時刻を同時に表示することで、ブロックが有効になっていることを証明しています。")

	// hostsファイルの画像 h3
	pcHostsTitle := head3Template("hostsファイルの画像")
	pcHostsImg := imageTemplate(log, NotionTodayHostsImage)
	//  quote

	children := notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			spTitle,
			spTgtTimeTitle,
			spTgtTime,
			spImg,
			pcTitle,
			pcExpQuote,
			pcHostsTitle,
			pcHostsImg,
		},
	}

	newToggleBlock := toggleTemplate("デバイス", &children)
	return newToggleBlock
}

func DiaryHeaderTemplate(log *LifeLog) notionapi.AppendBlockChildrenRequest {
	// 	## 日記

	// これまでの記録を全て見る
	listPageLinkQuote := quoteTemplate("これまでの記録を全て見る")
	listPageLink := pageTemplate(YuyaKanshiPageId)
	diaryTitle := head3Template("本日の予定")
	diaryCalenderImg := imageTemplate(log, NotionCalenderPicture)
	newBlocks := notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			listPageLinkQuote,
			listPageLink,
			diaryTitle,
			diaryCalenderImg,
		},
	}
	newToggleBlock := toggleTemplate("サマリー", &newBlocks)
	return newToggleBlock
}

func HealthTemplate(log *LifeLog) notionapi.AppendBlockChildrenRequest {
	// 	## 食事
	healthTitle := head3Template("食事")

	// 目標カロリーが上限か下限か

	// - 摂取カロリー
	healthCalTitle := head3Template("摂取カロリー")
	healthCal := quoteTemplate(strconv.Itoa(log.TodayCalorie))

	// - ズレ

	// ## 筋トレ
	muscleTitle := head3Template("筋トレ")

	// トレーニング記録
	musclePageTitle := head3Template("本日のトレーニング記録")
	musclePage := pageTemplate(log.TrainingPageId)

	if log.TrainingPageId != "" {
		// - ジムに行った証明
		muscleImageTitle := head3Template("ジムに行った証明")
		muscleImage := imageTemplate(log, NotionMyFitnessPal)

		newBlocks := notionapi.AppendBlockChildrenRequest{
			Children: []notionapi.Block{
				healthTitle,
				healthCalTitle,
				healthCal,
				muscleTitle,
				musclePageTitle,
				musclePage,
				muscleImageTitle,
				muscleImage,
			},
		}
		newToggleBlock := toggleTemplate("食事・筋トレ", &newBlocks)
		return newToggleBlock
	} else {
		newBlocks := notionapi.AppendBlockChildrenRequest{
			Children: []notionapi.Block{
				healthTitle,
				healthCalTitle,
				healthCal,
			}}
		newToggleBlock := toggleTemplate("食事", &newBlocks)
		return newToggleBlock
	}

}
func toggleBlockTemplate(text string, children *notionapi.AppendBlockChildrenRequest) *notionapi.ToggleBlock {
	newToggleBlock := &notionapi.ToggleBlock{
		BasicBlock: notionapi.BasicBlock{
			Object: "block",
			Type:   notionapi.BlockTypeToggle,
		},
		Toggle: notionapi.Toggle{
			RichText: []notionapi.RichText{
				{
					Type: "text", // ここを修正しました
					Text: &notionapi.Text{
						Content: text,
					},
				},
			},
			Children: children.Children,
		},
	}
	return newToggleBlock
}
func toggleTemplate(text string, children *notionapi.AppendBlockChildrenRequest) notionapi.AppendBlockChildrenRequest {
	newToggleBlock := notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			&notionapi.ToggleBlock{
				BasicBlock: notionapi.BasicBlock{
					Object: "block",
					Type:   notionapi.BlockTypeToggle,
				},
				Toggle: notionapi.Toggle{
					RichText: []notionapi.RichText{
						{
							Type: "text", // ここを修正しました
							Text: &notionapi.Text{
								Content: text,
							},
						},
					},
					Children: children.Children,
				},
			},
		},
	}
	return newToggleBlock
}
