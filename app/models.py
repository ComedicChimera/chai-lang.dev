from app.config import PAGE_SIZE
from app import db

import string
from random import choice
from datetime import datetime

from sqlalchemy import desc

class Suggestion(db.Model):
    id = db.Column(db.String(32), primary_key=True, unique=True)
    title = db.Column(db.String(80), nullable=False)
    date = db.Column(db.DateTime, nullable=False)
    author = db.Column(db.String(80), nullable=False)
    accepted = db.Column(db.Boolean, nullable=False)
    body = db.Column(db.String, nullable=False)


def get_suggestions(orderby, page):
    if orderby == 'recent':
        return Suggestion.query.order_by(desc(Suggestion.date)).limit(PAGE_SIZE).all()


def add_suggestion(title, author, body):
    id = random_string(32)

    while len(Suggestion.query.filter_by(id=id).all()):
        id = random_string(32)
    
    s = Suggestion(id, title, datetime.today(), author, False, body)

    db.session.add(s)
    db.session.commit()


def to_html(model):
    if isinstance(model, Suggestion):
        return """<div class="suggestion"><div class="title-box"><span class="title">{model.title.python_type}</span>
        <span class="author">{model.author.python_type}</span><span class="date">{model.date.python_type.strftime('%b %d, %Y')}</span>
        <div class="body">{model.body.python_type}</div>
        """


def random_string(length): return ''.join(choice(string.ascii_letters + string.digits) for _ in range(length))


db.create_all()
