# gh depscheck

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
