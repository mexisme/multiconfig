# `multiconfig` Package

This is a fairly simple package that was extracted from another wrapper binary.

The purpose of the wrapper was to provide config from multiple sources to a child / sub-process via
environment variables, i.e. from [12-factor](https://12factor.net/), the ["Store config in the Environment"
(https://12factor.net/config) aspect.

This package can take config from multiple config files and environment variables.

It can merge them together into a single KV-style string (i.e. `map[string]string`).

It can then convert that map into a array of strings in the same format as `os.Environ()` provides, and
compatible with the "Env" field of the [os/exec.Cmd](https://golang.org/pkg/os/exec/#Cmd) type and the
`envv` arg [syscall.Exec](https://golang.org/pkg/syscall/#Exec) function.

More detail can be found in [godoc.org](https://godoc.org/github.com/mexisme/multiconfig).

## Supported config sources
Currently, it supports:
- Reading enviroment variables provided to the binary / app
- `.env` files
- TOML files
- YAML files
- JSON files

## Caveats
Since the output format is intended to be compatible with `os.Environ()` it's assumed the provided configs
are also flat KV-pairs, with a `string` value.

## Build/Test
This is a library, meant to be imported into another binary, so "build" doesn't really
make sense in this context.

However, there is a [`Makefile`](./Makefile) -- copied from another of my projects, so
needs some tidying for this project -- to wrap certain functionality:

- Tests:
  - `make test` will run the Ginkgo/Gomega based tests
  - `make test-verbose coverage` will also generate coverage report(s)
- Vendorising packages:
  - Currently using Go [`dep`](https://github.com/golang/dep), at least until Go v1.11 is
    released
  - `make vendor` will download dependencies to the local `./vendor/` directory
  - `touch Gopkg.toml; make vendor` should try to update the vendorised packages to the
    next most-appropriate version(s) (depending on the SemVer constraints)

## Background
The primary reason for supporting multiple configs is for providing secrets and credentials to an app:
different groups of secrets were accessible to different teams, but all needed to be made available to the
underlying app

The reasons for choosing to do this via environment variables is this methodology is mature and understood,
and most languages have a library that supports this.

There are plenty of great Secrets management solutions like:
- [GCP KMS Service](https://cloud.google.com/kms/)
- [AWS SSM](https://aws.amazon.com/blogs/compute/managing-secrets-for-amazon-ecs-applications-using-parameter-store-and-iam-roles-for-tasks/)
- [Hashicorp Vault](https://www.vaultproject.io/)
- [Docker Swarm Secrets](https://docs.docker.com/engine/swarm/secrets/)

The reasons for not choosing to build-in support for something like the above is that support across
languages is a bit inconsistent -- not that many are blessed by AWS, for example -- and it makes it more
difficult to work with multiple Secrets services or Cloud providers.

## To Do
See [TODO.org](./TODO.org)
