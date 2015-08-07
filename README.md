## gitio [![GoDoc](https://godoc.org/github.com/xlab/gitio?status.svg)](https://godoc.org/github.com/xlab/gitio)

This is a Go client for [git.io].
Read more about [git.io] here: https://github.com/blog/985-git-io-github-url-shortener

[git.io]: http://git.io

### Installation
```
$ go get github.com/xlab/gitio/...
```

### Usage
```
$ gitio -h
Usage: gitio <long url>
  -c, --code=""        A custom code for the short link, e.g. http://git.io/mycode
  -f, --force=false    Try to shorten link even if the custom code has been used previously.
```

### License

[MIT](http://xlab.mit-license.org)
