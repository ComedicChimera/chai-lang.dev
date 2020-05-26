# NOTE: this image should only be built once

FROM python:3
ENV PYTHONUNBUFFERED 1
COPY requirements.txt /
RUN pip install -r requirements.txt
RUN mkdir /src
COPY src /src/
WORKDIR /src
