# gRPC Long-lived Streaming

Basic PubSub example, clients get subscribed to a topic, as a result they open a stream from where they receive server updates.

## Protocol & gRPC generation

To compile the proto file, run the following command from the `protos` folder:

```
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative longlived.proto
```


## Running the server

Navigate to `/server` folder and run the following:

```
$ go build server.go
$ ./server
```



## Running the client(s)

The client process emulates several clients (total flag).

```
$ go build client.go 
$ ./client --total=XXXX
```

