# apache-thrift
Using Apache Thrift with Go

## Installation guide
Make sure you have thrift installed in your machine, if you don't, follow the steps here:https://thrift.apache.org/download

### Generating thrift files
Run the command bellow to generate thrift code 

``make thrift-gen``

## Testing
Before testing, make sure there's nothing running on your port :8080

### Server
In a separated terminal, run:

``make server``

### Client
Open another terminal tab and run

``make client``

This command will send a request to the RPC server

