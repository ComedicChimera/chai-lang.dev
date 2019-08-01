from app.config import PAGE_SIZE
from app import db

import string
from random import choice
from datetime import datetime

from sqlalchemy import desc
from markupsafe import escape

class Suggestion(db.Model):
    id = db.Column(db.String(32), primary_key=True, unique=True)
    title = db.Column(db.String(80), nullable=False)
    date = db.Column(db.DateTime, nullable=False)
    author = db.Column(db.String(80), nullable=False)
    email = db.Column(db.String(100), nullable=False)
    accepted = db.Column(db.Boolean, nullable=False)
    body = db.Column(db.String, nullable=False)

    def __repr__(self):
        return '<Suggestion: {id}>'

    def to_html(self):
        f"""<div class="suggestion"><span class="title">{escape(self.title)}</span>
        <span class="author">By <span class="author-name">{escape(self.author)}</span></span>
        <span class="date">{self.date.strftime('%b %d, %Y')}</span>
        <div class="body">{escape(self.body)}</div></div>
        """


def get_suggestions(orderby, page, query_string):
    if query_string != '':
        ordered_result = Suggestion.query.order_by(query_string in Suggestion.body)
    elif orderby == 'recent':
        ordered_result = Suggestion.query.order_by(desc(Suggestion.date))
    elif orderby == 'accepted':
        ordered_result = Suggestion.query.order_by(Suggestion.accepted)
    elif orderby == 'alpha':
        ordered_result = Suggestion.query.order_by(Suggestion.title)
    elif orderby == 'user':
        ordered_result = Suggestion.query.order_by(Suggestion.author)

    if page > 1 and page * PAGE_SIZE < ordered_result.count():
        return ordered_result.offset(PAGE_SIZE * (page - 1)).limit(PAGE_SIZE).all()
    else:
        return ordered_result.limit(PAGE_SIZE).all()
    


def add_suggestion(title, author, email, body):
    id = random_string(32)

    while len(Suggestion.query.filter_by(id=id).all()):
        id = random_string(32)
    
    s = Suggestion(id=id, title=title, date=datetime.today(), author=author, email=email, accepted=False, body=body)

    db.session.add(s)
    db.session.commit()


def random_string(length): return ''.join(choice(string.ascii_letters + string.digits) for _ in range(length))


db.create_all()
