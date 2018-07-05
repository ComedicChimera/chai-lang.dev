import flask
from argparse import ArgumentParser


app = flask.Flask(__name__)


@app.route('/')
def home():
    return "<h1>Home</h1>"


@app.route('/guide/<name>')
def guide(name):
    return flask.render_template('guide.html', page_content=name)


if __name__ == '__main__':
    parser = ArgumentParser()
    parser.add_argument('--debug', '-d', action='store_true')
    parser.add_argument('--port', '-p', type=int, default=5000)
    args = parser.parse_args()
    app.run(debug=args.debug, port=args.port)

