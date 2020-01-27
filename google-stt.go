package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	languageCodePtr := flag.String("language", "en-US", "set the language of the incoming voice data (e.g. 'en-US')")
	flag.Parse()

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
			LanguageCode:    *languageCodePtr,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: content},
		},
	})

	if err != nil {
		log.Fatalf("Failed to recognize audio: %v", err)
	}

	if len(response.Results) == 0 || len(response.Results[0].Alternatives) == 0 {
		log.Fatal("Failed to transcode audio.")
	}

	fmt.Print(response.Results[0].Alternatives[0].Transcript)
}
