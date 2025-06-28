from bs4 import BeautifulSoup

# Load your scraped HTML (replace this with actual file load if needed)
with open("Talks/1.html", "r", encoding="utf-8") as f:
    html = f.read()

soup = BeautifulSoup(html, 'html.parser')

# Find all <p> tags
paragraphs = soup.find_all('p')

# Version 1: Text with reference links included
print("\n=== Paragraphs with HTML (including links) ===")
for p in paragraphs:
    print(str(p))  # or p.prettify() for indented HTML

# Version 2: Plain text only (strip tags)
print("\n=== Cleaned Paragraph Text ===")
for p in paragraphs:
    print(p.get_text(strip=True))
