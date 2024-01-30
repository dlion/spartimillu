## Spartimillu

Application Level Load Balancer written in Go

## Author
Domenico Luciani   
https://domenicoluciani.com


## Spartimillu Server - Load Balancer
It represents the load balancer which is in charge of redirecting requests towards other servers   
To run it:
1. Set the address of your server instances into the main
2. Set your healthcheck endpoint
3. Set your healtcheck time frame
4. `go build`
5. `./spartimillu`

### Test the load balancer
`curl http://localhost/ --output -`

## bg - Dummy background services    
This directory contains the folders which will represent 2 services, to start them `cd bg` and into two separates terminal instances:
1. server8080: `python -m http.server 8080 --directory server8080`
2. server8081: `python -m http.server 8081 --directory server8081`

## Tests

The project is composed by unit and integration tests, to run all of them: `go test ./...`.
1. unit tests: `server/server_test.go`
2. integration tests: `client/client_test.go`

## Idea
This is the implementation of the [WC Coding Challenge](https://codingchallenges.fyi/challenges/challenge-load-balancer) with Go, following those requirements step-by-step.    
Any idea or feedback to make it better are always welcomed!