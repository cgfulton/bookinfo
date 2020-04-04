# bookinfo-prompt

![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)

![Go Report Card](https://goreportcard.com/badge/github.com/cgfulton/bookinfo-prompt)

![Go](https://github.com/cgfulton/bookinfo-prompt/workflows/Go/badge.svg)

An interactive cli featuring auto-complete using [go-prompt](https://github.com/c-bata/go-prompt) inspired by [kube-promtp](https://github.com/c-bata/kube-prompt).

bookinfo-prompt accepts commands needed to configure and deploy the Istio Bookinfo example. 

And you can integrate other commands via pipe (`|`).

```
>>> get pod | grep web
web-1144924021-2spbr        1/1     Running     4       25d
web-1144924021-5r1fg        1/1     Running     4       25d
web-1144924021-pqmfq        1/1     Running     4       25d
```

## Goal

Hopefully support following commands enough to configure, orchestrate, and deploy the Bookinfo example.

* [x] `get`            Display one or many resources
* [x] `describe`       Show details of a specific resource or group of resources
* [x] `create`         Create a resource by filename or stdin
* [x] `replace`        Replace a resource by filename or stdin.