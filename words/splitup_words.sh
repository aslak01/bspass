#!/bin/bash

input_file="words"

process_word() {
  word="$1"
  length="${#word}"
  output_file="words_${length}"
  echo "$word" >>"$output_file"
}

while IFS= read -r line; do
  process_word "$line"
done <"$input_file"

echo "Word list split up by length."
