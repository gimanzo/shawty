[![GoDoc](https://godoc.org/github.com/gimanzo/shawty?status.svg)](http://godoc.org/github.com/gimanzo/shawty)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/gimanzo/shawty/master/LICENSE)

# Shawty: URL Shortener

This service encodes URLs using Hashids

It has 3 features:

### Encode
```
curl -XPOST localhost:8080/encode --data "url=http://dillbuck.com"

>>> pnel5aK

```

### Decode
```
curl -XPOST localhost:8080/decode/pnel5aK

>>> http://dillbuck.com

```

### Redirect
```
curl -XPOST localhost:8080/redirect/pnel5aK

>>>  301 -> http://dillbuck.com

```

## Getting started

```
go get github.com/gimanzo/shawty
export SHAWTY_STORAGE_PATH=<your local storage path>
export SHAWTY_SALT=<your salt>
```


### Can I use it in production?

You need to implement a storage that can scale beyond one application server.