from app import app
import app.models as db
import app.md as md

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
            obj = request.get_json()

            try:
                db.add_suggestion(obj['title'], obj['name'], obj['email'], obj['body'])
            except IndexError:
                abort(422)
            except Exception:
                abort(500)

            return 200
        else:
            abort(415)
    else:
        return render_template('suggestions-make.html')


@app.route('/suggestions/view')
def view_suggestions():
    orderby = request.args.get('orderby', 'recent')

    if orderby in ['recent', 'accepted', 'alpha', 'user']:
        suggestions = db.get_suggestions(orderby, 1)
        results = '\n'.join((map(lambda x: db.to_html(x), suggestions)))
    else:
        results = '<span class=\"invalid-query\">Invalid query</span>'

    return render_template('suggestions-view.html', results=results)


@app.route('/docs/spec')
def lang_spec():
    html = md.load_markdown('/spec.md')

    html = html.replace('language-whirlwind', 'language-javascript', 3)

    return render_template('docs-spec.html', doc=html)
