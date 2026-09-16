# Goa Examples

[![Slack](https://img.shields.io/badge/slack-gophers-orange.svg?style=flat)](https://gophers.slack.com/messages/goa/)

This repository contains examples of microservices implemented using
[Goa](https://github.com/goadesign/goa). Each example focuses on a specific topic as indicated by
the directory name. The [cellar](https://github.com/goadesign/examples/tree/main/cellar) example
provides a complete implementation of a simple microservice.
The [retry](https://github.com/goadesign/examples/tree/main/retry) example shows how generated
HTTP and gRPC clients safely retry temporary failures.

The v3.31.0 examples accompany Goa v3.31.0 and require Go 1.26 or later.
This Goa release includes intentional generated-source and transport changes.
When adapting an existing application, read the
[upgrade guide](https://github.com/goadesign/goa/blob/v3.31.0/UPGRADING.md)
and use matching Goa and examples versions. Regenerate the complete `gen`
directory and update handwritten interceptors, multipart decoders, and command
starters as described there.

The samples in each directory serve as templates, and you can clone them using the `gonew` command, as explained in [this blog post](https://go.dev/blog/gonew).

```shell
$ go install golang.org/x/tools/cmd/gonew@latest
$ gonew goa.design/examples/basic@latest github.com/<your_repo>/basic
$ cd basic
```

A [fully instrumented example](https://github.com/goadesign/clue/tree/main/example/weather) of a
system consisting of multiple Goa microservices is included in the [Clue](https://github.com/goadesign/clue) repo.

To get started with the Goa Clue example, you can use the gonew command to clone it into your own repository:

```shell
$ gonew github.com/goadesign/clue/example/weather github.com/<your_repo>/weather
$ cd weather
```
Please follow the README in the Clue repository for more details on running and testing the Goa Clue example.

As you study each example consider contributing back by providing better or more complete docs,
adding clarifying comments to code or fixing any error you may run across!
