# Pipet: snippet manager

[![CircleCI](https://circleci.com/gh/dbalan/pipet/tree/master.svg?style=svg)](https://circleci.com/gh/dbalan/pipet/tree/master)

Pipet is a set of commands to store and retrieve snippets of text. Depends on
[fzf](https://github.com/junegunn/fzf) for search.

## Versioning
Follows [semantic versioning](https://semver.org/spec/v2.0.0.html).

## Installation
There are multiple ways to get pipet.
1. As a go package (provided you have a working go setup)
```
go get github.com/dbalan/pipet
cd $GOPATH/src/github.com/dbalan/pipet/
go build
go install # installs to $GOPATH/bin
```

2. As a binary release
see: https://github.com/dbalan/pipet/releases. On *nix you need to set executable permission (`chmod +x pipet`)

## Configuration
Pipet looks for config file `.pipet.yaml` in the home directory. `pipet init` command can generate a new config.

### Sample config

```yaml
document_dir: "<directory-where-files-are-stored>" # default is ~/snippets
editor_binary: "absolute path to editor you want to use" # default is $EDITOR environment variable
```

## Usage

[![asciicast](https://asciinema.org/a/MR8G05JXEIVY1AvKDrfKNjIEy.png)](https://asciinema.org/a/MR8G05JXEIVY1AvKDrfKNjIEy)

```
Usage:
  pipet [command]

Available Commands:
  init        Configure pipet
  new         Creates a new snippet and opens editor to edit content
  show        Show snippet
  delete      Remove snippet from storage (this is irreversible!)
  edit        Edit snippet data
  help        Help about any command
  list        List all snippets

Flags:
      --config string   config file (default is $HOME/.pipet.yaml)
  -h, --help            help for pipet
  -t, --toggle          Help message for toggle

```

## TODO
  - [x] finish configure command
  - [x] hacking docs
  - [x] circleci build
  - [x] binary downloads
  - [x] make public
  - [ ] Ability to search full text, with a flag to search command
  - [x] Try to abstract snippet id from operations, one way to do this is to move id's optional for commands and jump to a search interface in case IDs are not specified.
  - [ ] Tests, would like more tests.
  - [ ] Add an archive flag for delete, the data is not deleted, but is not exposed unless user turns on another flag.

## Hacking
0. Uses [dep](https://golang.github.io/dep/) for dependency management.
```bash
# Clone repo
WORKDIR=$GOPATH/src/github.com/dbalan/pipet
git clone git@github.com:dbalan/pipet $OWORKDIR
cd $WORKDIR
dep ensure  # dep needes to be installed ofcourse
go build
```
1. Pipet is mostly written to suit my workflow, which means it is quite limited.
   I would be happy to accept pull requests to improve and change workflows, but
   please open an Issue explaining what and why you want to change something
   before opening a PR.

## Thanks
Pipet takes a heavy inspiration from [pet](https://github.com/knqyf263/pet) and other projects.

