from app import app
from flask import render_template, redirect
import markdown
import json
import os


guide_names = []

for file in os.listdir('static/markdown/guide'):
    with open('static/markdown/guide/' + file) as fp:
        guide_names.append(fp.readline()[1:].strip())


def load_markdown(path):
    with open('static/markdown' + path) as file:
        md_file = file.read()

    html = markdown.markdown(md_file).replace('<code>', '<code class=\'language-whirlwind\'>')

    return html


@app.route('/')
def index():
    return render_template('index.html', prefix='')


@app.route('/docs')
def docs():
    return render_template('docs.html', prefix='../')


@app.route('/docs/guide')
def guide_home():
    return redirect('docs/guide/chapter1', 302)


@app.route('/docs/guide/<chapter>')
def guide(chapter):
    html = load_markdown(f'/guide/{chapter}.md')

    return render_template('guide.html', prefix='../../../', content=html, guide_names=guide_names)
