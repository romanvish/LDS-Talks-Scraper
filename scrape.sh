#!/bin/bash

mkdir -p Talks

for ((i=1; i<=10; i++)); do
  url="https://scriptures.byu.edu/content/talks_ajax/$i/"
  output="Talks/$i.html"
  
  echo "Downloading $url to $output"
  curl -fsSL -o "$output" "$url" || echo "Failed to download $url"
done