# Pipet: snippet manager

Pipet is a set of commands to store and retrieve snippets of text.

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
TBD


## Configuration
Pipet looks for config file `.pipet.yaml` in the home directory.

### Sample config
```yaml
document_dir: "<directory-where-files-are-stored>" # default is ~/snippets
fzf: "<path-to-fzf-binary>" # default is /usr/bin/fzf
```

## Usage
  - pipet new : create a new snippets
  - pipet search : search through current snippets (only titles and tags for now)
  - pipet edit id : edit a snippet by id
  - pipet show id : show a snippet
  - pipet echo id : like show, but only prints the snippet data.
  - pipet list : list all snippets
  - pipet configure: initial config

## TODO
  - [ ] finish configure command
  - [ ] Search full text, with a flag to search command
  - [ ] Try to abstract snippet id from operations.
  - [ ] Tests, would like more tests.

## Changes
1. Pipet is mostly written to suit my workflow, which means it is quite limited.
   I would be happy to accept pull requests to improve and change workflows, but
   please open an Issue explaining what and why you want to change something
   before opening a PR.

## Thanks
Pipet takes a heavy inspiration from pet and other projects.

