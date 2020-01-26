#!/usr/bin/env python3

from argparse import ArgumentParser
from sys import stdin

from google.cloud import speech
from google.cloud.speech import enums

arg_parser = ArgumentParser()
arg_parser.add_argument("-l", "--language", help="set the language of the incoming voice data (e.g. 'en-US')", default="en-US")
args = arg_parser.parse_args()

client = speech.SpeechClient()

content = stdin.buffer.read()
audio = { "content": content }
encoding = enums.RecognitionConfig.AudioEncoding.LINEAR16
config = {
  "language_code": args.language,
  "sample_rate_hertz": 16000,
  "encoding": encoding,
}

response = client.recognize(config, audio)

print(response.results[0].alternatives[0].transcript)
