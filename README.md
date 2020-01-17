# Welcome Modules Using Gopher!

To see slides, make sure you have [caddy v1](https://caddyserver.com/v1/)
and run this:

```console
$ cd slides
$ caddy
```

Press `f` to enter full screen.

## Todo before demos

### [`basic-public`](./basic-public)
You don't need to do anything before the [`basic-public`](./basic-public) demo.

### [`basic-private`](./basic-private)

Before the [`basic-private`](./basic-private) demo, make sure to run:

```console
$ git config --global --add url."git@github.com:".insteadOf "https://github.com/"
```

So that the `go` tool will check out private code from GitHub over SSH.

>This command is in the steps of the [README](./basic-private/README.md) in that demo

### [`athens-private`](./athens-private)

You'll need a GitHub personal access token for this. Create one [here](https://github.com/settings/tokens). When you create it, just enable the `repo` scope and nothing else.

After you have it, store it in an environment variable on the command line. Here's how to do it on Linux/Mac systems:

```console
$ export ATHENS_GITHUB_TOKEN="<your token>"
```

Then you'll be ready to run [that demo](./athens-private/README.md).

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
