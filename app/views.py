from app import app
from flask import render_template


@app.route('/')
def index():
    return app.send_static_file('html/index.html')


@app.route('/docs')
def docs():
    return render_template('docs.html', prefix='../')
