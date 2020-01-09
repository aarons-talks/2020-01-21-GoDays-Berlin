# Welcome Modules Using Gopher!

To see slides, make sure you have [caddy v1](https://caddyserver.com/v1/)
and run this:

```console
$ cd slides
$ caddy
```

Outline:

- What are Go Modules & Really brief history
- What changed in Go 1.13
  - Public proxy
  - Sum DB
    - go.sum and trust on first use
- Impact of the sum DB on you
- Impact for private projects
- What knobs to turn
  - `GOPRIVATE`
  - `GONOSUMDB`
  - `GOPROXY`
- Changing the defaults
  - Alternatives to the public proxy
  - Running your own private proxy
    - And its impact on the sum DB
