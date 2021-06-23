# HTTP server/client

To start http server:
```
go run main.go server
```

### Test it using Curl
```
curl --request GET 'http://localhost:9999/headers' --header 'Content-Type: application/json' --header 'Cookie: __cfduid=d29e90877b57f1ff5efc5f28db538c32e1616600580' 
```

```
curl --request POST 'http://localhost:9999/foo' --header 'Content-Type: application/json' --data-raw '{"email": "marcos.quesada+5@holded.com", "pass": "Marcos"}' 
```

```
curl --request POST 'http://localhost:9999/bar' --header 'Content-Type: application/json' --data-raw '{"email": "marcos.quesada+5@holded.com", "pass": "Marcos"}' 
```

### Use http client from command

```
go run main.go client
```