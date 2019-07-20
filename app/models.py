from app.config import PAGE_SIZE
from app import db


def get_suggestions(orderby, page):
    return ["test"]


def to_html(model):
    return '<div class=\"suggestion\">{model}</div>'
