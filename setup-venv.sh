#!/usr/bin/env bash

DIR=$( cd $( dirname "$0" ) && pwd )
VENV_DIR="$DIR/env"

if [ -d "$VENV_DIR" ]; then
  echo "Virtual environment '$VENV_DIR' already exists."
  exit 1
fi

echo "Creating virtual environment"
python3 -m venv "$VENV_DIR"

if [ -d "$VENV_DIR/bin" ]; then
  # Unix
  source "$VENV_DIR/bin/activate"
elif [ -d "$VENV_DIR/Scripts" ]; then
  # Windows
  source "$VENV_DIR/Scripts/activate"
else
  echo "Something went wrong while creating the virtual environment."
  exit 1
fi

echo "Installing dependencies"
pip3 install -r "$DIR/requirements.txt"

echo "Done!"
