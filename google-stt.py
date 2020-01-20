#!/usr/bin/env python3

from sys import stdin

from google.cloud import speech
from google.cloud.speech import enums

client = speech.SpeechClient()

content = stdin.buffer.read()
audio = { "content": content }
encoding = enums.RecognitionConfig.AudioEncoding.LINEAR16
config = {
  "language_code": "de-DE",
  "sample_rate_hertz": 16000,
  "encoding": encoding,
}

response = client.recognize(config, audio)

print(response.results[0].alternatives[0].transcript)
