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
see: https://github.com/dbalan/pipet/releases. On *nix you need to set executable permission (`chmod +x pipet`)

## Configuration
Pipet looks for config file `.pipet.yaml` in the home directory. `pipet init` command can generate a new config.

### Sample config

```yaml
document_dir: "<directory-where-files-are-stored>" # default is ~/snippets
editor_binary: "absolute path to editor you want to use" # default is $EDITOR environment variable
```

## Usage

[![asciicast](https://asciinema.org/a/pDumZGUeirlDHdzieWtNB5riL.png)](https://asciinema.org/a/pDumZGUeirlDHdzieWtNB5riL)

```
Store and sprinkle code snippets

Usage:
  pipet [command]

Available Commands:
  delete      Remove snippet from storage (this is irreversible!)
  edit        edit snippet data
  help        Help about any command
  init        Configure pipet
  list        list all snippets
  new         Creates a new snippet and opens editor to edit content
  show        display the snippet

Flags:
      --config string   config file (default is $HOME/.pipet.yaml)
  -h, --help            help for pipet
  -t, --toggle          Help message for toggle

Use "pipet [command] --help" for more information about a command.
```

## TODO
  - [ ] Tests, would like more tests.
  - [ ] Add an archive flag in place of delete (?)

## Hacking
See CONTRIBUTING.md

## Versioning
Follows [semantic versioning](https://semver.org/spec/v2.0.0.html).

## Thanks
Pipet takes a heavy inspiration from [pet](https://github.com/knqyf263/pet) and other projects.

