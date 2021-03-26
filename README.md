# grpc-user-api

This is a significant evolution of @blushi's [original Golang-based user-api](https://github.com/resonatecoop/user-api-old)

The changes are so significant a new repo was created, but a lot of code lives on from that repo.

It builds on that work in several important ways:

- drops Twirp framework in favour of [GRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) which has gained significant traction
- implements full OpenAPIV2 workflow - write interfaces in protobufs and generate the code stubs, then implement them.
- exposes full Swagger UI automatically (currently 100% not seamless, WIP)
- implements full RBAC using native Golang Interceptors (arguably better than using Twirp Handlers)
- RBAC is based on User role and interface access config in the config file
- built with Go modules for dependency management

It is WIP

## Running

Running `main.go` starts a web server on https://0.0.0.0:11000/. You can configure
the port used with the `$PORT` environment variable, and to serve on HTTP set
`$SERVE_HTTP=true`.

```
$ go run main.go
```

An OpenAPI UI is served on https://0.0.0.0:11000/.

### Running the standalone server

If you want to use a separate gRPC server, for example one written in Java or C++, you can run the
standalone web server instead:

```
$ go run ./cmd/standalone/ --server-address dns:///0.0.0.0:10000
```

## Getting started

After cloning the repo, there are a couple of initial steps;

1. Install the generate dependencies with `make install`.
   This will install `buf`, `protoc-gen-go`, `protoc-gen-go-grpc`, `protoc-gen-grpc-gateway`,
   `protoc-gen-openapiv2` and `statik` which are necessary for us to generate the Go, swagger and static files.
1. If you forked this repo, or cloned it into a different directory from the github structure,
   you will need to correct the import paths. Here's a nice `find` one-liner for accomplishing this
   (replace `yourscmprovider.com/youruser/yourrepo` with your cloned repo path):
   ```bash
   $ find . -path ./vendor -prune -o -type f \( -name '*.go' -o -name '*.proto' \) -exec sed -i -e "s;github.com/merefield/grpc-user-api;yourscmprovider.com/youruser/yourrepo;g" {} +
   ```
1. Finally, generate the files with `make generate`.

Now you can run the web server with `go run main.go`.

## Maintenance

Interfaces are designed in
`proto/` directory. See https://developers.google.com/protocol-buffers/
tutorials and guides on writing protofiles.

Once that is done, regenerate the files using
`make generate`. This will mean you'll need to implement any functions in
`server/`, or else the build will fail since your struct won't
be implementing the interface defined by the generated file in `proto/example.pb.go`.
