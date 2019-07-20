import app.config as config

import sys, os
from flask import Flask
from flask_sqlalchemy import SQLAlchemy


app = Flask('__main__')
app.config['SQLALCHEMY_DATABASE_URI'] = config.DB_URI

db = SQLAlchemy(app)


def run_app(args):
    from app import views

    app.run(**args)
