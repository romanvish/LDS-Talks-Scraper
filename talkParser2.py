from bs4 import BeautifulSoup
import re
import os
import sys

talk_name_start = int(sys.argv[1])
talk_name_end = int(sys.argv[2]) + 1
# print(f'python num is {talk_name_start}, {talk_name_end}')

for talk_num in range(talk_name_start, talk_name_end):
    # Load your scraped HTML (replace this with actual file load if needed)
    raw_file = f"raw_text/talks/{talk_num}.html"
    with open(raw_file, "r", encoding="utf-8") as f:
        html = f.read()

    if 'Error (500)' in str(html):  # error handling for bad status codes (for some reason it returns lots of 500s)
        os.remove(raw_file)  # delete the bad files
        continue

    soup = BeautifulSoup(html, 'html.parser')

    # Find all <p> tags
    paragraphs = soup.find_all('p')

    paragraph_texts = []
    for p in paragraphs:
        text = p.get_text(strip=True)
        if text:  # Only add non-empty paragraphs
            paragraph_texts.append(text)

    filename = os.path.join('clean_text', 'talks_processed', f'{talk_num}.txt')
    with open(filename, 'w', encoding='utf-8') as f:
        for text in paragraph_texts:
            f.write(text + '\n\n')  # Double newline between paragraphs

    # print(f"Saved to: {filename}.txt")
