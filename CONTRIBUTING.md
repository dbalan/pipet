# Hacking
0. Pipet uses [dep](https://golang.github.io/dep/) for dependency management.

```bash
# Clone repo
WORKDIR=$GOPATH/src/github.com/dbalan/pipet
git clone git@github.com:dbalan/pipet $WORKDIR
cd $WORKDIR
dep ensure  # dep needes to be installed ofcourse
go build
```

1. Pipet is mostly written to suit my workflow, which means it is quite limited.
   I would be happy to accept pull requests to improve and change workflows, but
   if the changes are big, please open an Issue explaining what and why you want
   to change something.
