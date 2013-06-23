# serve

Like `python -m SimpleHTTPServer`, plus more!

Serve a directory on the host and port of your choice.

## Usage

```bash
$ serve
Serving '/home/james/' on 'http://127.0.0.1:8000'

$ serve -p 1234
Serving '/home/james/' on 'http://127.0.0.1:1234'

$ serve /usr/share/doc/python3-doc/html
Serving '/usr/share/doc/python3-doc/html/' on 'http://127.0.0.1:8000'

$ serve -h 0.0.0.0
Serving '/home/james/' on 'http://0.0.0.0:8000'
```

## Install

```bash
go get github.com/jamesadney/serve
```
