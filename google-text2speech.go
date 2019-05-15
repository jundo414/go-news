package main

// http://codegists.com/code/google-cloud-speech-example/

import (
	"context"
	"io/ioutil"
	"log"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/jessevdk/go-flags"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type Options struct {
	// https://cloud.google.com/text-to-speech/docs/voices
	// "ja-JP-Wavenet-A" or "ja-JP-Standard-A" in Japan
	Voice string  `short:"v" long:"voice" description:"voice name" default:"ja-JP-Standard-A"`
	Lang  string  `short:"l" long:"lang" description:"language code" default:"ja-JP"`
	Rate  float64 `short:"r" long:"rate" description:"speaking rate" default:"1.0"`
	Out   string  `short:"o" long:"out" description:"output filename" default:"speech.mp3"`
}

func ReadMessage(text string) {
	var opts Options

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	bytes := text

	ctx := context.Background()
	c, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: string(bytes),
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: opts.Lang,
			Name:         opts.Voice,
			// SsmlGender: 1,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			SpeakingRate:  opts.Rate,
		},
	}
	resp, err := c.SynthesizeSpeech(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(opts.Out, resp.GetAudioContent(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	SayText(opts.Out)
}
