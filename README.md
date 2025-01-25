# imvinhnguyen

code for imvinhnguyen.com

## Status

[![Go-test](https://github.com/sudovinh/imvinhnguyen/actions/workflows/go-test.yml/badge.svg)](https://github.com/sudovinh/imvinhnguyen/actions/workflows/go-test.yml)
[![goreleaser](https://github.com/sudovinh/imvinhnguyen/actions/workflows/release.yml/badge.svg)](https://github.com/sudovinh/imvinhnguyen/actions/workflows/release.yml)

## Getting Started

This project uses [devbox](https://github.com/jetify-com/devbox) to manage its development environment.

Install devbox:

```sh
curl -fsSL https://get.jetpack.io/devbox | bash
```

Start the devbox shell:

```sh
devbox shell
```

Run a script in the devbox environment:

```sh
devbox run <script>
```

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
PORT=<PORT NUMBER> make run
```

Live reload the application:

```bash
PORT=<PORT NUMBER> make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```

## Notes

Created using [Go-blueprint by Melkey](https://github.com/Melkeydev/go-blueprint)
