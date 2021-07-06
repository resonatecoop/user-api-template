# user-api

This is a significant evolution of @blushi's [original Golang-based user-api](https://github.com/resonatecoop/user-api-old)

The changes are so significant a new repo was created, but a lot of code lives on from that repo.

It builds on that work in several important ways:

- drops Twirp framework in favour of [GRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) which has gained significant traction
- implements full OpenAPIV2 workflow - write interfaces in protobufs and generate the code stubs, then implement them.
- exposes full Swagger UI automatically
- implements full RBAC using native Golang Interceptors (arguably better than using Twirp Handlers)
- RBAC is based on User role and interface access config in the config file
- built with Go modules for dependency management
- adds a CLI for database management and for running the server
- replaces `go-pg` with `bun`
- merges in the models from `resonatecoop\id`

It is WIP, do NOT use this in Production yet!

## Running

Running `go run main.go runserver` starts a web server on https://0.0.0.0:11000/. You can configure
the port used with the `$PORT` environment variable, and to serve on HTTP set
`$SERVE_HTTP=true`.

```
$ go run main.go runserver
```

An OpenAPI UI is served on https://0.0.0.0:11000/.

## Getting started

After cloning the repo, there are a couple of initial steps;

1. Install the generate dependencies with `make install`.
   This will install `buf`, `protoc-gen-go`, `protoc-gen-go-grpc`, `protoc-gen-grpc-gateway`,
   `protoc-gen-openapiv2` and `statik` which are necessary for us to generate the Go, swagger and static files.
2. Install the git submodule(s) with `git submodule update --init` from root directory of the cloned repo
3. Finally, generate the files with `make generate`.

## Dev database setup

* Create user and database as follows (as found in the local config file in `./conf.local.yaml`):

username = "resonate_dev_user"

password = "password"

dbname = "resonate_dev"

```
CREATE DATABASE resonate_dev;

CREATE USER resonate_dev_user WITH PASSWORD 'password';

GRANT ALL PRIVILEGES ON DATABASE resonate_dev TO resonate_dev_user;

```

Add following postgres extensions: "hstore" and "uuid-ossp"

```
\c resonate_dev;
CREATE EXTENSION hstore;
CREATE EXTENSION "uuid-ossp";
```

You can confirm with

```
SELECT * FROM pg_extension;
```

From the root of the user-api:

* Init migrations

```sh
$  go run main.go db -env dev init
```

* Run migrations

```sh
$ go run main.go db -env dev migrate
```

rolling back:

```sh
$ go run main.go db -env dev rollback
```

* Loading default fixtures
```sh
$ go run main.go db -env dev load_default_fixtures
```

* Loading test data (fixtures)
```sh
$ go run main.go db -env dev load_test_fixtures
```

The same can be repeated for the test database (`resonate_test`) substituting `dev` for `test`, e.g:

```sh
$  go run main.go db -env test init
```

## Tests

Ongoing WIP atm, but for example, can be run with:

```sh
$  go test -timeout 30s -run ^TestDeleteUser$ github.com/resonatecoop/user-api/server/users
```

## Running!

Now you can run the web server with `go run main.go runserver`.

This will default to running in dev and without debug output from queries.

However, there are two flags:

- `env` (`dev`, `test` or `prod`, defaults to `dev`)
- `dbdebug` (`true` or `false`, defaults to `false`)

So e.g. `go run main.go runserver -env test -dbdebug true` will run the server on the test DB (defined in the config) with db query debug output on.

The PSN connection strings for dev and test are in conf.local.yaml, but for prod this is built from environement variables:

-	`POSTGRES_NAME` (DB name)
- `POSTGRES_USER` (DB username)
- `POSTGRES_PASS` (DB password)
- `POSTGRES_HOST` (DB host, defaulted `127.0.0.1`)
- `POSTGRES_PORT` (DB port, defaulted `5432`)
- `POSTGRES_SSL` (`enable` or defaulted `disable`)

## Docker!

Build a container with `docker build -t resonateuserapi .`

(to avoid cache, use `--no-cache` option)

Run container with `docker run -p 11000:11000 --network=host -tid resonateuserapi`

Check status with `docker container ls` and `docker logs <image>`

(use sudo as required)

## Working with a reverse-proxy (like nginx)

You need to register a certificate in a pair of files.

This prefix of this file is held in the config file as `cert_name`, default value `uaclient`

The server will then look for two files, suffix's ".pem" and ".key" in the directory provided by 
the environment variable `UACERT_DIR`

In your reverse proxy you will need to refer to these too in order to be able to proxy the service securely.

## Maintenance

Interfaces are designed in
`proto/` directory. See https://developers.google.com/protocol-buffers/
tutorials and guides on writing protofiles.

Once that is done, regenerate the files using
`make generate`. This will mean you'll need to implement any functions in
`server/`, or else the build will fail since your struct won't
be implementing the interface defined by the generated file in `proto/example.pb.go`.
