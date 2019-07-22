import app.config as config

import sys, os, time
from flask import Flask
from flask_sqlalchemy import SQLAlchemy


app = Flask('__main__')
app.config['SQLALCHEMY_DATABASE_URI'] = config.DB_URI
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

# forced delay since docker won't allow me to healthcheck db
time.sleep(1)

db = SQLAlchemy(app)


def run_app(args):
    from app import views

    app.run(host='0.0.0.0', **args)
