#!/bin/bash

mkdir -p Talks

RANGE_START=$1
RANGE_END=$2

for ((i=RANGE_START; i<=RANGE_END; i++)); do
# TODO -- update this so that the range can be specified as anything (not starting at 1)
  url="https://scriptures.byu.edu/content/talks_ajax/$i/"
  output="Talks/$i.html"
  
  echo "Downloading $url to $output"
  curl -fsSL -o "$output" "$url" || echo "Failed to download $url"
done