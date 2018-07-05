import flask

app = flask.Flask(__name__)


@app.route('/')
def index():
    return flask.render_template('index.html')


def run_server(**kwargs):
    app.run(**kwargs)
