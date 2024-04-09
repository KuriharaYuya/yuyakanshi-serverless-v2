package repoutils

import (
	"errors"
	"fmt"
	"os"
)

func S3FileName(logDate string, imageType string) string {
	return fmt.Sprintf("/tmp/%s_%s.png", logDate, imageType)
}
func S3FileUrl(logDate string, imageType string) string {
	s3Fname := S3FileName(logDate, imageType)
	return DefineImageURLs(s3Fname)
}

func GetImageExternalURl(logDate string, imageType string) string {
	if logDate == "" {
		// エラーを出して終了
		err := errors.New("logDate is empty")
		fmt.Println(err)
		return ""
	}
	path := S3FileName(logDate, imageType)
	url := "https://%s.s3.amazonaws.com/images%s"
	url = fmt.Sprintf(url, os.Getenv("BUCKET_NAME"), path)
	return url
}

func DefineImageURLs(path string) string {
	url := "https://%s.s3-%s.amazonaws.com/images/%s"
	url = fmt.Sprintf(url, os.Getenv("BUCKET_NAME"), os.Getenv("REGION"), path)
	return url
}
