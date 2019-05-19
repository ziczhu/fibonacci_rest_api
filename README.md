# Fibonacci Rest API
---
A Fibonacci REST Web Service in Golang

## Overview
This project provides a simple API to generate fibonacci sequence.

```
GET /api/v1/fibonacci/:n

n should be >= 0 and smaller than the predefined limit(see below)
```

Success Response:
```
GET /api/v1/fibonacci/5
// HTTP 200 OK
{
    status: 200,
    data: [0, 1, 1, 2, 3]
}
```

Error Response:
```
GET /api/v1/fibonacci/abc
// HTTP 400 Bad Request
{
    status: 400,
    error_code: 40000,
    error_msg: "invalid number 'abc'"
}
```

## Setup

### 1. Using Docker & Docker Compose
The easiest way to running this project is using docker.
Please make sure you already install Docker & Docker Compose.

1.Clone to your local.
```
git clone https://github.com/ziczhu/fibonacci_rest_api.git ${your_folder}
```

2.Build the images
```
make build
```

3.Start the project

```
make start
```
At this point, you can visit the api at:
http://localhost/api/v1/fibonacci/:n

4.Watch the logs

```
make logs
```

5.Shutdown the service
```
make stop
```

### 2.Run service in your local machine.

Please make sure your already installed the tools below:
* golang >= 1.11
* dep >= 0.5

1.Clone to your local.
```
git clone https://github.com/ziczhu/fibonacci_rest_api.git ${your_folder}
```

2.Install the dependency.
```
dep ensure
```

3.Build the service.
```
go build -o fibonacci .
```

4.Run the service.
```
PORT=80 ./fibonacci start
```
At this point, you can visit the api at:
http://localhost/api/v1/fibonacci/:n.
And the log is in stdout.

5.Shutdown the service.
```
ctrl-c or just kill the process.
```

## Configuration

There are some global env we can utilize to config
this service.

| name                | type   | default | description                                                                                                                                                                                                                                                                |
|---------------------|--------|---------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ENV                 | string | dev     | Tweak behavior in different env, options: 'dev' | 'prod'.                                                                                                                                                                                                                  |
| PORT                | int    | 8080    | The PORT for the Golang service                                                                                                                                                                                                                                            |
| MAX_FIB_INPUT       | int    | 10000   | The max input of the fibonacci sequence, since the memory cost is huge, we need to setup the limit according to server.                                                                                                                                                    |
| INIT_FIB_CACHE_SIZE | int    | 1000    | The init size of the cached sequence, to reduce to computation, there will be a cached sequence in the memory, INIT_FIB_CACHE_SIZE will step up the initial size of the cache before starting.                                                                             |
| MAX_FIB_CACHE_SIZE  | int    | 5000    | MAX_FIB_CACHE_SIZE defines the max size of the cached sequence, a 10000-length fibonacci sequence(using big.Int internally) allocates more than 50MiB memory, so we won't grow up the cache endlessly. For the huge numbers, just calculate and return or setup the limit. |

You can either update docker-compose.yml if you're using docker,
or passed as env like `MAX_FIB_INPUT=1000 ./fibonacci start`.

## Testing

### Using Docker
`make docker-test` running the tests in docker container.
```
make docker-test
```

### Local Testing
`make test` running the tests in your machine,
please make sure you install the dependency first.
```
make test
```
