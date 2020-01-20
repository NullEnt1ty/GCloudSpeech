# GCloudSpeech

Transcribe voice data from stdin to text using [Google Cloud Speech-to-Text](https://cloud.google.com/speech-to-text/).

## Introduction

GCloudSpeech is a small application that allows you to parse voice data quick
and easy using the command line.

This project can be used for [Rhasspy](https://github.com/synesthesiam/rhasspy),
an offline, multilingual voice assistant toolkit. It's not an offline voice
assistant anymore when you're using Google STT but it's worth a shot if you need
good automatic speech recognition on a low-end device.

## Installation

1. Clone this repository

   ```
   $ git clone https://github.com/NullEnt1ty/GCloudSpeech
   ```

1. Setup the virtual environment

   ```
   $ ./setup-venv.sh
   ```

## Usage

You can transcribe voice data by piping it to `run.sh`. The transcription will
be printed on standard output.

For example:

```
$ cat podcast.wav | ./run.sh
```

**Attention**: Currently the only accepted format for voice data are
uncompressed 16-bit signed little-endian samples (Linear PCM) with a sample rate
of 16 kHz. This might be configurable in the future.

## Integration into Rhasspy

Use the following configuration for your profile to integrate GCloudSpeech into
Rhasspy:

```json
"speech_to_text": {
    "command": {
        "program": "<path to run.sh>"
    },
    "system": "command"
},
```
