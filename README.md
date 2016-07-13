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
<strong>NOTE: url must include the full url, http://etc... </strong>

### Decode
```
curl localhost:8080/decode/pnel5aK

>>> http://dillbuck.com

```

### Redirect
```
curl localhost:8080/pnel5aK

>>>  301 -> http://dillbuck.com

```

## Getting started

Docker
```
docker run -v <host path>:/data -e "SHAWTY_STORAGE_PATH=/data" -e "SHAWTY_HASH_SALT=<salt>" -d -p 8080:8080 gimanzo/shawty
```

Old School way:
```
go get github.com/gimanzo/shawty

export SHAWTY_STORAGE_PATH=<your local storage path>
export SHAWTY_SALT=<your salt>
//optional
export GOOGLE_TRACKER_ID=<your google analytics trackingid>

go build
./shawty
```

### Building the docker file

```
docker build -t gimanzo/shawty .
```


### Can I use it in production?

You need to implement a storage that can scale beyond one application server.