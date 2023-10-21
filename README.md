# Goa Examples

[![Slack](https://img.shields.io/badge/slack-gophers-orange.svg?style=flat)](https://gophers.slack.com/messages/goa/)

This repository contains examples of microservices implemented using
[Goa](https://github.com/goadesign/goa). Each example focuses on a specific topic as indicated by
the directory name. The [cellar](https://github.com/goadesign/examples/tree/master/cellar) example
provides a complete implementation of a simple microservice.

The samples in each directory serve as templates, and you can clone them using the `gonew` command, as explained in [this blog post](https://go.dev/blog/gonew).

```shell
$ go install golang.org/x/tools/cmd/gonew@latest
$ gonew goa.design/examples/basic@latest github.com/<your_repo>/basic
$ cd basic
```

A [fully instrumented example](https://github.com/goadesign/clue/tree/main/example/weather) of a
system consisting of multiple Goa microservices is included in the [Clue](https://github.com/goadesign/clue) repo. Please follow the README in the Clue repository for more details on running and testing the Goa Clue example.

As you study each example consider contributing back by providing better or more complete docs,
adding clarifying comments to code or fixing any error you may run across!
