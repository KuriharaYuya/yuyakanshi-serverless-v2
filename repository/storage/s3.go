package storage

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
	repoutils "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Session struct {
	S3Session *session.Session
}
type notionImageUrls map[string]string

const (
	MorningImage    = "morningImage"
	MyFitnessPal    = "myFitnessPal"
	ScreenTime      = "screenTime"
	TodayHostsImage = "todayHostsImage"
	CalenderPic     = "calenderPicture"
)

func DefineImageUrlAndPath(logDate string, imageType string) (s3URL string, localPath string) {
	localPath = repoutils.S3FileName(logDate, imageType)
	s3URL = repoutils.DefineImageURLs(localPath)
	return s3URL, localPath
}

type s3ImageUrls notionImageUrls

func SetUp() *Session {
	// paths := []string{"", ""}
	credential := credentials.NewStaticCredentials(
		os.Getenv("SECRET_ID"),
		os.Getenv("SECRET_KEY"),
		"",
	)
	awsConfig := aws.Config{
		Region:      aws.String(os.Getenv("REGION")),
		Credentials: credential,
	}
	s, err := session.NewSession(&awsConfig)
	if err != nil {
		log.Println("failed to create S3 session:", err.Error())
		return nil
	}
	return &Session{
		S3Session: s,
	}
}
func (s Session) uploadToS3(path string) error {

	upFile, err := os.Open(path)
	if err != nil {
		log.Printf("failed %s, error: %v", path, err.Error())
	}
	defer upFile.Close()
	upFileInfo, err := upFile.Stat()
	if err != nil {
		log.Printf("failed to get stat %s, error: %v", path, err.Error())
	}
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	refPath := "images/" + path
	// uploading
	_, err = s3.New(s.S3Session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("BUCKET_NAME")),
		Key:                  aws.String(refPath),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(int64(fileSize)),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		log.Printf("failed to upload %s, error: %v", path, err.Error())
	}
	url := repoutils.DefineImageURLs(refPath)
	fmt.Printf("Uploaded File Url %s\n", url)

	return nil
}

func downloadFromUrl(logDate string, imageType string, imageUrl string, localImagePath string) string {
	if imageUrl == "" {
		panic("imageUrl is empty: repository/storage/s3.go")
	}
	f, err := os.Create(localImagePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	res, err := http.Get(imageUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	_, err = io.Copy(f, res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("image downloaded")
	return localImagePath
}

func deleteFile(fname string) {
	err := os.Remove(fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("image deleted")
}

type uploadImages struct {
	images map[string]string
}

func UploadImages(log *notionpkg.LifeLog, s *Session) uploadImages {
	// まず、日付とtypeとpathのmapを受け取る
	imageProps := getImagesProps(log)
	logDate := log.Date

	// あとはforで回して、それぞれのpathをダウンロードして、S3にアップロードする
	// その後、ダウンロードした画像を削除する

	for imageType, imageUrl := range imageProps {
		s3URL, localPath := DefineImageUrlAndPath(logDate, imageType)
		downloadFromUrl(logDate, imageType, imageUrl, localPath)
		// uploadする
		s.uploadToS3(localPath)
		// その後、削除する
		deleteFile(localPath)
		fmt.Println(s3URL)

	}
	return uploadImages{}
}

func getImagesProps(log *notionpkg.LifeLog) notionImageUrls {
	data := notionImageUrls{
		MorningImage:    log.MorningImage,
		MyFitnessPal:    log.MyFitnessPal,
		ScreenTime:      log.ScreenTime,
		TodayHostsImage: log.TodayHostsImage,
		CalenderPic:     log.CalenderPicture,
	}

	return data
}

// func replaceImageURLToS3(lifeLog *notionpkg.LifeLog) *notionpkg.LifeLog {
// 	// images/tmp/2021-08-14_morningImage.png
// 	path :=
// 	newLifeLog := notionpkg.LifeLog{
// 		MorningImage:   repoutils.DefineImageURLs(S3FileName(lifeLog.Date, MorningImage)),
// 	}
// }
