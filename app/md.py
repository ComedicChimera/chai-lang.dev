import markdown
import os

guide_names = []
unordered_guide = {}

for file in os.listdir('static/markdown/guide'):
    with open('static/markdown/guide/' + file) as fp:
        unordered_guide[int(file[7:-3])] = fp.readline()[1:].strip()

for i in range(len(unordered_guide)):
    guide_names.append(unordered_guide[i + 1])

unordered_guide = {}


def load_markdown(path):
    with open('static/markdown' + path) as file:
        md_file = file.read()

    html = markdown.markdown(md_file).replace('<code>', '<code class=\'language-whirlwind\'>')

    return html.replace('<br />', '</p><p>')
