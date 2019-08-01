from app import app
import app.models as db
import app.md as md

import string
from urllib.parse import unquote

from flask import render_template, redirect, abort, request


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


@app.route('/suggestions')
def suggestions():
    return render_template('suggestions.html')


@app.route('/suggestions/make', methods=['GET', 'POST'])
def make_suggestion():
    if request.method == 'POST':
        if request.is_json:
            obj = request.get_json(silent=True)

            try:
                db.add_suggestion(obj['title'], obj['name'], obj['email'], obj['body'])
            except IndexError:
                abort(422)
            except Exception as e:
                return str(e)

            return "Suggestion was made successfully"
        else:
            abort(415)
    else:
        return render_template('suggestions-make.html')


def check_suggestion_args(args):
    orderby = request.args.get('orderby', 'recent')
    page = request.args.get('page', '1')
    query_string = request.args.get('query_string', '')

    if orderby not in ['recent', 'accepted', 'alpha', 'user']:
        return False
    elif not page.isdigit():
        return False
    elif query_string not in (string.ascii_letters + string.digits + '%'):
        return False

    return {orderby: orderby, page: int(page), query_string: unquote(query_string)}


@app.route('/suggestions/view')
def view_suggestions():
    check_result = check_suggestion_args(request.args)

    if check_result:
        suggestions = db.get_suggestions(**check_result)
        results = '\n'.join((map(lambda x: x.to_html(), suggestions)))
    else:
        results = '<span class=\"invalid-query\">Invalid query</span>'

    return render_template('suggestions-view.html', results=results)


@app.route('/suggestions/view_api')
def view_suggestions_html():
    check_result = check_suggestion_args(request.args)

    if check_result:
        suggestions = db.get_suggestions(**check_result)
        return '\n'.join((map(lambda x: x.to_html(), suggestions)))
    else:
        return '<span class=\"invalid-query\">Invalid query</span>'


@app.route('/docs/spec')
def lang_spec():
    html = md.load_markdown('/spec.md')

    html = html.replace('language-whirlwind', 'language-javascript', 3)

    return render_template('docs-spec.html', doc=html)
