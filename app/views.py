from app import app
from flask import render_template


@app.route('/')
def index():
    return render_template('index.html', prefix='')


@app.route('/docs')
def docs():
    return render_template('docs.html', prefix='../')
