# GoDays 2020 Private Code Demo

This is similar to the [basic private code demo](../basic-private), except that it uses [Athens](https://docs.gomods/io) to host _all_ of the code that we need, including the [private module](https://github.com/arschles/godays2020private) that we need.

We'll set up this demo by running Athens and setting up our local environment to ensure that we are talking to Athens and not leaking any private module names.

## Running Athens

First, you'll need to run Athens. We're using Docker for this, but there are [alternative installation options](https://docs.gomods.io/install/) that you can use if you'd like.

You'll need to set up Athens to be able to fetch code from the private repositories we're using, proxy the global `sum.golang.org` checksum database for public modules, and _not_ proxy the global checksum database for private modules:

```console
$ docker run \
    -p 3000:3000 \
    -e ATHENS_GONOSUM_PATTERNS="github.com/arschles/godays2020private" \
    -e ATHENS_GITHUB_TOKEN="$ATHENS_GITHUB_TOKEN" \
    gomods/athens:v0.7.1
```

>Make sure that the `ATHENS_GITHUB_TOKEN` environment variable is set to a [GitHub personal access token](https://github.com/settings/tokens) that has access to your private repository. I can't share mine for `github.com/arschles/godays2020private` because sadly I can't create a token with read access for just that repository.

## Configuring Local Environment

Now that Athens is configured and running, it will maintain its own database of module code and it will proxy the sum database for every module except for `github.com/arschles/godays2020private` - our private code. For this module, Athens will return an error.

You'll need to set up the local environment to fetch all modules from Athens (not `proxy.golang.org`), while also handling the case when Athens returns an error when the `go` tool requests the checksum for `github.com/arschles/godays2020private`.

In a new terminal window, run these commands to do your setup:

```console
$ go env -w GOPROXY=http://localhost:3000
$ go env -w GONOSUMDB=github.com/arschles/godays2020private
```

>`GOPROXY` tells the `go` tool to contact Athens for both module code and checksums. `GONOSUMDB` tells the `go` tool to contact Athens for checksums.

## Compile and run

Now that you're all set up, clear your cache and run the server in the same terminal window as the last step (just as you did in the [previous demo](../basic-private)):

```console
$ sudo rm -rf $(go env GOPATH)/pkg/mod
$ go run .
```
