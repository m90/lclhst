# lclhst

[![Build Status](https://travis-ci.org/m90/lclhst.svg?branch=master)](https://travis-ci.org/m90/lclhst)
[![godoc](https://godoc.org/github.com/m90/lclhst?status.svg)](http://godoc.org/github.com/m90/lclhst)

> Wait for localhost to be ready like what?!?!

Wait for localhost to be ready in order to run tests or anything that relies on localhost responding.

## Installation:

Install the library:
```sh
go get github.com/m90/lclhst
```

Install the command:
```sh
go get github.com/m90/lclhst
```

## Usage

Use the command before running anything that expects localhost to be up:

```sh
lclhst && echo "OK!"
```

The following options are available:

```
Usage of lclhst:
  -port int
    	the port of the application (default 8080)
  -timeout duration
    	timeout for giving up (default 10s)
```

### License
MIT © [Frederik Ring](http://www.frederikring.com)
