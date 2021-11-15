
# to-localhost

[![Go](https://github.com/qbantek/to-localhost/actions/workflows/go.yml/badge.svg)](https://github.com/qbantek/to-localhost/actions/workflows/go.yml)
[![CodeQL](https://github.com/qbantek/to-localhost/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/qbantek/to-localhost/actions/workflows/codeql-analysis.yml)


This barebones Go app helps developers to redirect a callback URL to localhost.

The idea was taken from https://tolocalhost.com/.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.17 or newer 
and the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) installed.

```sh
$ git clone https://github.com/qbantek/to-localhost.git
$ cd to-localhost
$ go build -o bin/to-localhost -v .
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
