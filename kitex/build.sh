#!/bin/sh

echo "enter the test file's name:"
read testcase

# server = "./server/kitex_server.go"
# client = "./client/kitex_client.go"


# Execute the client
go run "./client/kitex_client.go" $testcase

