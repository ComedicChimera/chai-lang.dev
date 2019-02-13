from flask import Flask
import sys, os


app = Flask('__main__')


def run_app(args):
    from app import views

    app.run(**args)