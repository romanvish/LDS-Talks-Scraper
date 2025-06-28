from bs4 import BeautifulSoup
import re
import os

# Load your scraped HTML (replace this with actual file load if needed)
with open("Talks/9.html", "r", encoding="utf-8") as f:
    html = f.read()

soup = BeautifulSoup(html, 'html.parser')

# Find all <p> tags
paragraphs = soup.find_all('p')

# get title/speaker to save the text file with this
title = re.sub(r'[^A-Za-z0-9]', '_', paragraphs[0].get_text(strip=True)).strip('_')
speaker = re.sub(r'[^A-Za-z0-9]', '_', paragraphs[1].get_text(strip=True)).replace('__', '_').strip('_')

paragraph_texts = []
for p in paragraphs:
    text = p.get_text(strip=True)
    if text:  # Only add non-empty paragraphs
        paragraph_texts.append(text)

filename = os.path.join('talksProcessed', re.sub(' ', '', title + '_' + speaker))
with open(filename, 'w', encoding='utf-8') as f:
    for text in paragraph_texts:
        f.write(text + '\n\n')  # Double newline between paragraphs

print(f"Saved to: {filename}.txt")