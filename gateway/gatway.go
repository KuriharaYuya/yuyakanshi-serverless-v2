package gateway

import (
	"bytes"
	"strings"
	"sync"

	utils "github.com/KuriharaYuya/yuya-kanshi-serverless/util"
)

var buf bytes.Buffer

const userAgent = "user-agent"
const lineBotWebhook = "LineBotWebhook"

func Gateway(req utils.Request) *utils.Response {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ua := req.Headers[userAgent]

	go func() {
		defer wg.Done()
		// line-bot-request
		if strings.Contains(ua, lineBotWebhook) {
			LineGateway(req)
		}
	}()

	resp := utils.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}
	wg.Wait()
	return &resp
}
