# generate-go-embed4assets
Generate go embed file for assets (.json,.yaml,.yml)

## Download and Install

### Release builds

[Download](https://github.com/f9n/generate-go-embed4assets/releases) the latest release from GitHub.

### Install using Homebrew

Using [Homebrew](https://brew.sh), you can install using the my Homebrew tap:

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
$ generate-go-embed4assets --help
Usage of generate-go-embed4assets:
  -directory string
    where to find assets (default "gen/go")
  -file-formats string
    evaluate these file format types (default ".json,.yaml,.yml")
  -version
    print the current version
$ generate-go-embed4assets -version
version v0.1.0, gitCommit: ..., buildDate: ..., buildUser: ...

$ generate-go-embed4assets -directory gen -file-formats=.json,.yml
[+] Generating gen/swagger/apidocs.swagger.json.embed.go
```
