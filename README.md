# Advent Of Code 2023

## Dev Env

I created a custom Docker container. 

To build it run:

```sh
docker build -t go-advent-code-2023 .
```

To run it use:

```sh
docker run --rm -v "$(pwd)"/src:/go/src -it go-advent-code-2023
```

## Run file 

```sh
go run day_xx/day_xx.go
```

## Run Tests

```sh
go test -v ./dayxx
```