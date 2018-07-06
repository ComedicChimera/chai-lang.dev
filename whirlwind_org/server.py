import flask
import whirlwind_org.guide as guide_loader

app = flask.Flask(__name__)


@app.route('/')
def index():
    return flask.render_template('index.html')


@app.route('/guide/chapter<int:chap_num>/<name>')
def guide(chap_num, name):
    if '#' in name:
        flask.abort(404)
    content = guide_loader.load_guide(chap_num, name)
    if not content:
        flask.abort(404)
    return flask.render_template('guide.html', content=content, chapters=guide_loader.load_chapter_bar(chap_num - 1, name))


@app.route('/guide/chapter<int:chap_num>')
def guide_chapter_title(chap_num):
    content, name = guide_loader.load_chapter_title(chap_num)
    if not content:
        flask.abort(404)
    return flask.render_template('guide.html', content=content, chapters=guide_loader.load_chapter_bar(chap_num - 1, name))


def run_server(**kwargs):
    app.run(**kwargs)
