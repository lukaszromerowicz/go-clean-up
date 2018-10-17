# go-clean-up
Clean up your directories with Go ✨

go-clean-up iterates over the files in your given directory and arranges them in new directories using file's `Last modified` value.

# Installation
1. Get `go-clean-up` first

`go get github.com/lukaszromerowicz/go-clean-up`

2. To add it to your path run

`cd $GOPATH/src/github.com/lukaszromerowicz/go-clean-up`

`GOBIN="/usr/local/bin" go install`

# Usage

`go-clean-up /root my_new_shiny_dir`

Where:
* `/root` is the path to directory you want to clean up
* `my_new_shiny_dir` is the name of directory you want go-clean-up to put your files in; defaults to `cleanup`
