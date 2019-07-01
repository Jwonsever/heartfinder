# Heartfinder
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FJwonsever%2Fheartfinder.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FJwonsever%2Fheartfinder?ref=badge_shield)


This is a project built using Go, Postgres, and Create-React-App.  It's a "blog" like website letting people post pictures and locations for the Hearts in San Francisco project.

https://sfghf.org/events/hearts-in-sf/

## How can I build it?

It's supposed here that you have `Golang`, `Node.JS` and `yarn` instllaed on your computer. First of all you need to creare new React application. It's not delivered as part of the source code to be sure that latest version of `create-react-app`. So, clone the project:

```
 mkdir -p $GOPATH/src/github.com/jwonsever/heartfinder
 cd $GOPATH/src/github.com/jwonsever/heartfinder
 git clone https://github.com/jwonsever/heartfinder.git .
```

Ok, now we have our project. Let's install it's dependencies and build it:

```
 cd $GOPATH/src/github.com/jwonsever/heartfinder/ui
 yarn install
 PUBLIC_URL=http://127.0.0.1:9999/ui/build yarn build
```

Ok, now we have things to play with. Let's prepare our server for build. It has only one dependency right now:

```
 go get gopkg.in/alecthomas/kingpin.v2
```

Let's build:

```
 cd $GOPATH/src/github.com/jwonsever/heartfinder
 go build -o bin/heartfinder cmd/main.go
```

## How I can run it?

Just run server you built:

```
 cd $GOPATH/src/github.com/jwonsever/heartfinder
 bin/heartfinder --listen 127.0.0.1:9999 --build ui/build --db "YOUR DB INFO"
```

And visit the web page:

```
 open http://127.0.0.1:9999
```

## License

Please, take a look at the [LICENSE](https://github.com/jwonsever/heartfinder/blob/master/LICENSE) file for details about this aspect of the project.


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FJwonsever%2Fheartfinder.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FJwonsever%2Fheartfinder?ref=badge_large)