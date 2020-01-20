#!/usr/bin/env bash

DIR=$( cd $( dirname "$0" ) && pwd )
VENV_DIR="$DIR/env"

if [ -d "$VENV_DIR/bin" ]; then
  # Unix
  source "$VENV_DIR/bin/activate"
elif [ -d "$VENV_DIR/Scripts" ]; then
  # Windows
  source "$VENV_DIR/Scripts/activate"
else
  echo "Could not activate virtual environment. Did you run setup-venv.sh?"
  exit 1
fi

"$DIR/google-stt.py" $@
