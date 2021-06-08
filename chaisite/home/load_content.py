import markdown
import os

CONTENT_DIR = "content/"
CONTENT_CACHE_DIR = "content/__cache__/"

# `file_path` should not have a trailing `.md`
def load_markdown(file_path):
    cached_path = os.path.join(CONTENT_CACHE_DIR, file_path, ".html")
    if os.path.exists(cached_path):
        with open(cached_path, 'r') as f:
            return f.read()

    # if we have an error here, we just catch it in the `view`
    md_path = os.path.join(CONTENT_DIR, file_path, ".md")
    with open(md_path, 'r') as md_f:
        with open(cached_path, 'w+') as cached_f:
            markdown.markdownFromFile(md_f, cached_f)

            return cached_f.read()