from app import app


@app.route('/')
def index():
    return app.send_static_file('html/index.html')


@app.route('/docs')
def docs():
    return app.send_static_file('html/docs.html')
