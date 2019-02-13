FROM python:latest

ADD . .

CMD ["pip", "install", "-r", "requirements.txt"]

EXPOSE 5000

RUN python index.py

