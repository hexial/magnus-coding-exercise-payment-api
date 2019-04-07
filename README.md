# Magnus Coding Excercise Payment API

## Build the system

```bash
./build.sh
```

Will run all unit tests and build docker images

Uses go build tag ```// +build unit``` for the tests


## Run the system

```bash
./run.sh
```

Will start all docker containers

## Integration Testing

```bash
./test-integration.sh
```

Will execute ```build.sh``` then ```run.sh```; when the system is running another docker container will start that will do all integration testing against the running system. 

Uses go build tag ```// +build integration``` for the tests
