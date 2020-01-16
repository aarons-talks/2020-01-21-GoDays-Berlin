# GoDays 2020 Private Code Demo

This is similar to the [basic public code demo](../basic-public), except that it uses a private repository as one of its dependencies: [`github.com/arschles/godays2020private`](https://github.com/arschles/godays2020private).

This new private repository provides a new `ImagesHandler` function that eliminates a lot of redundant code in [`cat_handler.go`](./cat_handler.go) and [`dog_handler.go`](./dog_handler.go).

In order to use it, we'll need to set a `GOPRIVATE` environment variable to tell the modules system to _not_ use `proxy.golang.org` to download this private repository:

```console
$ go env -w GOPRIVATE=github.com/arschles/godays2020private
```

>Note: this will not work unless you have access to my repository! There's unfortunately no way (that I know of) to give read access to just this private repository. I recommend using a private repository of your own.

Then, clear your local modules disk cache and build run the server, just as you did in the [`basic-public` demo](../basic-public):

```console
$ sudo rm -rf $(go env GOPATH)/pkg/mod
$ go run .
```

## Other Notes

You can set this `GOPRIVATE` environment variable to a list and/or wildcard module names. For example: 

```console
$ go env -w GOPRIVATE="github.com/myprivate/*,github.com/otherprivate/*`
```

See more on this and other related environment variables [on the GitHub wiki](https://github.com/golang/go/wiki/Modules#recent-changes) and the [golang.org documentation](https://golang.org/cmd/go/#hdr-Module_configuration_for_non_public_modules).
