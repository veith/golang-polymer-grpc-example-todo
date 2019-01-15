# The Server Project

## Start the servers
Run
```bash
$ go run cmd/grpcserer/grpcserver.go
$ go run cmd/transcoder/transcoder.go
$ go run cmd/reverseproxy/reverseproxy.go 

```

## To log in you have to..
**[POST] localhost:5555/api/auth**
```json5
{
 "username": "mario",
 "password": "party"
}
```

## API Documentation
Just paste the swagger files from ../proto/doc/*.swagger.json to [editor.swagger.io](https://editor.swagger.io/).

