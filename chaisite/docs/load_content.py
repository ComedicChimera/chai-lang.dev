import markdown
import os

CONTENT_DIR = "chaisite/docs/content/"
CONTENT_CACHE_DIR = CONTENT_DIR + "__cache__/"

# `file_path` should not have a trailing `.md`
def load_markdown(file_path):
    # if we have an error here, we just catch it in the `view`

    md_path = os.path.join(CONTENT_DIR, file_path + ".md")
    cached_path = os.path.join(CONTENT_CACHE_DIR, file_path + ".html")
    if os.path.exists(cached_path):
        if os.stat(cached_path).st_mtime > os.stat(md_path).st_mtime:
            with open(cached_path, 'r') as f:
                return f.read()

    with open(md_path, 'rb') as md_f:
        create_path(cached_path)

        with open(cached_path, 'wb+') as cached_f:
            markdown.markdownFromFile(input=md_f, output=cached_f, extensions=["tables"])

        with open(cached_path, 'r') as cached_f:
            return cached_f.read()

def create_path(file_path):
    dir_path = os.path.dirname(file_path)
    dirs_to_make = []
    while not os.path.exists(dir_path):
        dirs_to_make.append(dir_path)
        dir_path = os.path.dirname(dir_path)

    for path in reversed(dirs_to_make):
        os.mkdir(path)

    