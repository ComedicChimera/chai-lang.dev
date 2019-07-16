from app import app
import app.db as db
import app.md as md

from flask import render_template, redirect


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/docs')
def docs():
    return render_template('docs.html')


@app.route('/docs/guide')
def guide_home():
    return redirect('docs/guide/chapter1', 302)


@app.route('/docs/guide/<chapter>')
def guide(chapter):
    html = md.load_markdown(f'/guide/{chapter}.md')

    return render_template('docs-guide.html', content=html, guide_names=md.guide_names)


@app.route('/suggestions')
def suggestions():
    return render_template('suggestions.html')


@app.route('/suggestions/make')
def make_suggestion():
    return render_template('suggestions-make.html')


@app.route('/suggestions/view')
def view_suggestions_no_order():
    return redirect('/suggestions/view?orderby=recent', 302)


@app.route('/suggestions/view?orderby=<orderby>')
def view_suggestions(orderby):
    if orderby in ['recent', 'accepted', 'alpha', 'user']:
        suggestions = db.get_suggestions(orderby, 1)
        results = list(map(lambda x: db.to_html(x), suggestions))
    else:
        results = ['<span class=\"invalid-query\">Invalid query</span>']

    return render_template('suggestions-view.html', results=results)


@app.route('/suggestions/view?orderby=<orderby>&page=<int:page>')
def view_suggestions_by_page(orderby, page):
    if page > 0 and orderby in ['recent', 'accepted', 'alpha', 'user']:
        suggestions = db.get_suggestions(orderby, page)
        results = list(map(lambda x: db.to_html(x), suggestions))
    else:
        results = ['<span class=\"invalid-query\">Invalid query</span>']

    suggestions = db.get_suggestions(orderby, page)

    return render_template('suggestions-view.html', results=results)


@app.route('/docs/spec')
def lang_spec():
    html = md.load_markdown('/spec.md')

    html = html.replace('language-whirlwind', 'language-javascript', 3)

    return render_template('docs-spec.html', doc=html)
