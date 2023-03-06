# generate-go-embed4assets
Generate go embed file for assets (.json,.yaml,.yml)

## Download and Install

### Release builds

[Download](https://github.com/f9n/generate-go-embed4assets/releases) the latest release from GitHub.

### Install using Homebrew

Using [Homebrew](https://brew.sh), you can install using the My personal Homebrew tap:

```shell
brew install f9n/homebrew-tap/generate-go-embed4assets
```

### Install from Source


To install the `generate-go-embed4assets` command line tool, run:

```shell
go install github.com/f9n/generate-go-embed4assets/cmd/generate-go-embed4assets@latest
```

## Usage

```shell
$ generate-go-embed4assets version

$ generate-go-embed4assets -directory gen/go -file-formats=.json,.yml

```