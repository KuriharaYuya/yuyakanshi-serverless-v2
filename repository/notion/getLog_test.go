package notionpkg_test

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
// )

// func init() {

// 	os.Setenv("NOTION_API_KEY", "secret_htiPMbGDUwgvCuzHYQJnS64KAPMYaeyW8yQvNkKhXrC")
// }
// func Test_GetLog(t *testing.T) {

// 	t.Parallel()
// 	// Databaseの型が帰ってくるか
// 	// キャストする

// 	value, err := notionpkg.GetLog()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	value, ok := interface{}(value).(notionpkg.Database)
// 	fmt.Println(value)
// 	if !ok {
// 		fmt.Println(value)
// 	}
// }
