# EchoGen

EchoGen is a simple Yeoman scaffolding for [Labstack's Echo v3](https://github.com/labstack/echo) web framework.

It uses [Glide](https://github.com/Masterminds/glide) for application's dependencies. Go to [glide.yaml.tmpl](https://github.com/mdouchement/echogen/blob/master/templates/glide.yaml.tmpl) to see current dependencies versions.


## Installation

```sh
$ go get -u github.com/mdouchement/echogen

# Dependencies that you need to develop your application.
$ go get -u github.com/Masterminds/glide
$ go get -u github.com/jteeuwen/go-bindata/...

# Used in Makefile for LiveReload.
$ brew install fswatch
```


## Usage

EchoGen assumes that you have a well configured `$GOPATH` (e.g. `$GOPATH=/go/`) to generates the whole project.

```sh
$ cd $GOPATH/src/github.com/mdouchement
$ echogen --name lss

$ cd lss
$ glide install

# run the server with LiveReload
$ make serve

# or manually
$ go generate && go run myapp.go server -b localhost -p 5000
```


## License

**MIT**


## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request
