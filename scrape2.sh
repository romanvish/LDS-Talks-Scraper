#!/bin/bash

mkdir -p raw_text/talks

RANGE_START=$1
RANGE_END=$2

for ((i=RANGE_START; i<=RANGE_END; i++)); do
  url="https://scriptures.byu.edu/content/talks_ajax/$i/"  # max is 8800 (Confidence in the Presence of God by Russell M. Nelson, April 2025)
  output="raw_text/talks/$i.html"
  
  echo "Attempting to download $url to $output"
  curl -sSL -o "$output" "$url" || echo "Failed to download $url"
done