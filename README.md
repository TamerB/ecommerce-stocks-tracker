# ecommerce-stocks-tracker
This solution keeps an accurate list of products and their stocks up to date


# E-commerce Stocks Tracker service


## Overview

This service keeps an accurate list of products and their stocks up to date.
It provides the following endpoints:
```
/readyz                                                         GET
/healthz                                                        GET

/products/{sku}                                                 GET
/products/{sku}/stocks                                          GET

/products/{sku}/stocks/{country}                                PUT
```

## Developer builds

#### Generating Server using Swagger

+ This project uses ***Swagger version: v0.26.0***.
+ go-swagger configuration commands are stored in `/Makefile`.
+ To check or install ***go-swagger***, run `make check_install`.
+ To generate server with ***go-swagger*** run `make generate_server`.


#### Builds

From the project's root directory:

```
make generate_server
go get -u all
go mod tidy

make build
```

## Running

```bash
#!/bin/sh

export PORT=9090
export DB_DRIVER=postgres
export DB_SOURCE=postgresql://<username>:<password>@localhost:5432/<db-name>?sslmode=<enable-or-disable>

make run
```
To run the build with `http` connection only, please run `.bin/config-service --scheme http`.


### Arguments
#### `-v`
Verbose mode, configures the service to output debug level log entries.

### Environment Variables

#### `PORT`
Ports which the service will be listening on to `http` requests.

#### `DB_DRIVER`
Database driver.
#### `DB_SOURCE`
Database source.
### `DB_CONTAINER_NAME`
DB Container name. This is used by Makefile to Run DB Docker container, create DB, and drop DB.


## Curl

After running the application, please go to `<http-or-https>://<host>:<port>/docs` for example CURL requests.

Example: `http://localhost:9090/docs`

### GET Curl requests on terminal:

```
curl -i http://<host>:<port-or-tls-port>/readyz
curl -i http://<host>:<port-or-tls-port>/healthz

curl -i http://<host>:<port>/products/<product-sku>
curl -i http://<host>:<port>/products/<product-sku>/stocks
```
If you are using self signed TLS certificate, you should add `-H` argument for the request to pass.

#### Examples:
```
curl -i http://localhost:9090/readyz
curl -i http://localhost:9090/healthz
curl -i http://localhost:9090/products/b6d7e2a7fa99
```

## Test

To run tests, from the project's root directory, run `make test` in terminal.