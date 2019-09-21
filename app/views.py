from app import app
import app.models as db
import app.md as md

from flask import render_template, redirect, abort, request
from flask_limiter import Limiter
from flask_limiter.util import get_remote_address


limiter = Limiter(app=app, key_func=get_remote_address)


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/docs')
def docs():
    return render_template('docs.html')


@app.route('/docs/guide')
def guide_home():
    return redirect('/docs/guide/chapter1', 302)


@app.route('/docs/guide/<chapter>')
def guide(chapter):
    html = md.load_markdown(f'/guide/{chapter}.md')

    return render_template('docs-guide.html', content=html, guide_names=md.guide_names)


@app.route('/docs/spec')
def lang_spec():
    f = open('static/html/spec.html', 'r')
    html = f.read()

    f.close()

    return render_template('docs-spec.html', doc=html)


@app.route('/suggestions')
def suggestions():
    return render_template('suggestions.html')


@app.route('/suggestions/make')
def make_suggestion():
    return render_template('suggestions-make.html')


@app.route('/suggestions/view')
def view_suggestions():
    return render_template('suggestions-view.html', orderby=request.args.get('orderby', 'recent'), page=request.args.get('page', '1'))


@app.route('/api/suggestions/view', methods=['GET'])
@limiter.limit('200/day;10/minute')
def api_view_suggestions():
    check_result = db.check_suggestion_args(request.args)

    if check_result:
        suggestions = db.get_suggestions(**check_result)
        return '\n'.join((map(lambda x: x.to_html(), suggestions)))
    else:
        abort(422)


@app.route('/api/suggestions/make', methods=['POST'])
@limiter.limit('10/day;1/minute')
def api_make_suggestion():
    if request.is_json:
        obj = request.get_json(silent=True)

        try:
            db.add_suggestion(obj['title'], obj['name'], obj['email'], obj['body'])
        except IndexError:
            abort(422)
        except Exception:
            abort(500)

        return "Suggestion was made successfully"
    else:
        abort(415)
