# Heroku Go Example

This is a simple example of a zero-dependency Go API application running on Heroku

## Components

```
heroku-go-example         -- Repository / package name
├── Makefile              -- Targets for local development (optional)
├── Procfile              -- Defines a single `web` process
├── README.md             -- This file
├── bin
│   └── heroku-go-example -- The locally-built binary (gitignored)
├── main.go               -- The main Go application
└── vendor
    └── vendor.json       -- Buildpack identification / Heroku configuration [1]
```

## Running Locally

When the application is deployed, the [Heroku Go buildpack](https://github.com/heroku/heroku-buildpack-go) will run `go install .` to build and install your application to `/app/bin`.  We want to mimic this in local development, so the `run` target in the `Makefile` will:

1. Compile `main.go` to `bin/heroku-go-example` -- this is important because Go will use the pacakge name (`heroku-go-example`) when creating the binary.
1. Prepend the `./bin` directory to the `$PATH` and run `heroku local`.  This allows the `web` process in the `Procfile` to find the `heroku-go-example` binary.

This is all done for you by running `make` since `run` is the default target.

## Deploying

```
$ heroku apps:create heroku-go-example
$ git push heroku master
```

## Testing

```
GET:  curl -s https://heroku-go-example.herokuapp.com | jq .
POST: curl -s -d '{"key":"value"}' https://heroku-go-example.herokuapp.com | jq .
```

## Links

* [1] https://github.com/kardianos/govendor
