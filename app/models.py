from app.config import PAGE_SIZE
from app import db


class Suggestion(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(80), nullable=False)
    date = db.Column(db.String(10), unique=True, nullable=False)
    author = db.Column(db.String(80), nullable=False)
    body = db.Column(db.String(1000), nullable=False)


def get_suggestions(orderby, page):
    return ["test"]


def to_html(model):
    return '<div class=\"suggestion\">{model}</div>'


db.create_all()
