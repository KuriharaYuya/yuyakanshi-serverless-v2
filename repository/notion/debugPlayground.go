package notionpkg

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

// func DebugStructExchange
func GetDebugData(publish bool) (DebugLogProperty, error) {
	db, err := setLogDB()
	if err != nil {
		fmt.Println("setLogDBでエラー", err)
		return DebugLogProperty{}, err
	}
	query := setQuery(publish)

	// クエリを使用してデータを取得

	results, err := client.Database.Query(context.Background(), notionapi.DatabaseID(db.ID), query)
	if err != nil {
		fmt.Println("client.Database.Queryでエラー", err)
	}

	resultProps := results.Results[0].Properties
	props := map[string]interface{}{
		"Name":         &notionapi.TitleProperty{},
		"allowPublish": &notionapi.CheckboxProperty{},
	}

	debugLog := DebugLogProperty{}
	for key := range props {
		prop, ok := resultProps[key]
		if !ok {
			fmt.Println("Error: ", key, " property not found.")
			return DebugLogProperty{}, err
		}

		switch key {
		case "Name":
			if nameProp, ok := prop.(*notionapi.TitleProperty); ok {
				debugLog.Name = nameProp
			} else {
				fmt.Println("Error extracting properties for ", key)
				return DebugLogProperty{}, err
			}
		case "allowPublish":
			if allowPublishProp, ok := prop.(*notionapi.CheckboxProperty); ok {
				debugLog.AllowPublish = allowPublishProp
			} else {
				fmt.Println("Error extracting properties for ", key)
				return DebugLogProperty{}, err
			}
		}
	}
	return debugLog, nil
}

func setQuery(publish bool) *notionapi.DatabaseQueryRequest {
	checkboxCondition := &notionapi.CheckboxFilterCondition{}
	if publish {
		checkboxCondition.Equals = true
	} else {
		checkboxCondition.DoesNotEqual = true
	}

	query := &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: "allowPublish",
			Checkbox: checkboxCondition,
		},
	}
	return query
}
