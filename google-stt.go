package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	content, _ := ioutil.ReadAll(os.Stdin)
	ctx := context.Background()
	client, err := speech.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create speech-to-text client: %v", err)
	}

	response, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "de-DE",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: content},
		},
	})

	if err != nil {
		log.Fatalf("Failed to recognize audio: %v", err)
	}

	fmt.Printf("Result count: %v\n", len(response.Results))

	for _, result := range response.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("'%v', (confidence = %3f)", alt.Transcript, alt.Confidence)
		}
	}
}
