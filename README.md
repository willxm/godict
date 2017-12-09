Chinese-English translation tools in the command line
===
[![Build Status](https://travis-ci.org/willxm/godict.svg?branch=master)](https://travis-ci.org/willxm/godict)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)][license]


[license]: https://github.com/willxm/godict/blob/master/LICENSE

### config
go to http://ai.youdao.com/gw.s
apply for your own app key

replace the key in the file(GOPATH/github.com/willxm/godict/utils/const.go)
thank you

### install
```
go get github.com/willxm/godict
cd GOPATH/src/github.com/willxm/godict
go install
```
### usage

```
$ godict good

n. 好处；善行；慷慨的行为
adj. 好的；优良的；愉快的；虔诚的
adv. 好
n. (Good)人名；(英)古德；(瑞典)戈德

$ godict 你好
hello
hi
```