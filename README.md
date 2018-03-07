# Pipet: snippet manager

[![CircleCI](https://circleci.com/gh/dbalan/pipet/tree/master.svg?style=svg)](https://circleci.com/gh/dbalan/pipet/tree/master)

Pipet is a set of commands to store and retrieve snippets of text. Depends on
[fzf](https://github.com/junegunn/fzf) for search.

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
see: https://github.com/dbalan/pipet/releases


## Configuration
Pipet looks for config file `.pipet.yaml` in the home directory.

### Sample config
```yaml
document_dir: "<directory-where-files-are-stored>" # default is ~/snippets
editor_binary: "absolute path to editor you want to use" # default is $EDITOR environment variable
```

## Usage

[![asciicast](https://asciinema.org/a/W6tv7bN9z76EAlZJZDS025JwU.png)](https://asciinema.org/a/W6tv7bN9z76EAlZJZDS025JwU)

  - pipet new : create a new snippets
  - pipet search : search through current snippets (only titles and tags for now)
  - pipet edit id : edit a snippet by id
  - pipet show id : show a snippet
  - pipet echo id : like show, but only prints the snippet data.
  - pipet list : list all snippets
  - pipet configure: TBD

## TODO
  - [x] finish configure command
  - [x] hacking docs
  - [x] circleci build
  - [x] binary downloads
  - [x] make public
  - [ ] Search full text, with a flag to search command
  - [ ] Try to abstract snippet id from operations.
  - [ ] Tests, would like more tests.

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

