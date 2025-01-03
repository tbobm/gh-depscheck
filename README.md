# gh depscheck

[![Release](https://github.com/tbobm/gh-depscheck/actions/workflows/release.yaml/badge.svg)](https://github.com/tbobm/gh-depscheck/actions/workflows/release.yaml)

Find outdated dependencies in Github Actions Workflows.

## Features

- Compare version used by Github Actions Worfklows with latest available releases

## Building depscheck

### Container image

```bash
$ docker build -t depscheck:local .
```

### From source

```bash
$ go build -o ./depscheck ./cmd/
```
