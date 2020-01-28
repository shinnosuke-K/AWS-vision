package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func main() {
	url := "/Users/shinnosuke/Desktop/ocr/IMG_2561.JPG"
	image, err := os.Open(url)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(image)
	if err != nil {
		log.Fatal(err)
	}

	sess := session.Must(session.NewSession())
	svc := rekognition.New(sess, aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))

	params := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			Bytes: bytes,
		},
	}

	resp, err := svc.DetectText(params)
	if err != nil {
		log.Fatal(err)
	}

	for _, text := range resp.TextDetections {
		fmt.Println(*text.DetectedText)
	}
}
