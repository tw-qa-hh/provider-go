FROM golang

WORKDIR /app

RUN gem install pact-provider-verifier

COPY . /app