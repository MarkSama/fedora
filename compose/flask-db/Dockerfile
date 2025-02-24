FROM python:latest

#ENV http_proxy=http://10.0.0.2:3128
#ENV https_proxy=http://10.0.0.2:3128

RUN mkdir app


COPY ./apt.conf /etc/apt/apt.conf
COPY ./app.py /app
COPY ./requirements.txt /app

WORKDIR /app


RUN pip install -r requirements.txt --proxy http://10.0.0.2:3128

RUN apt update -y
RUN apt install default-mysql-client -y

CMD ['python3','/app/app.py']