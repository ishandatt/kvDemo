To run:

`go run main.go -c=./kv_server.conf`

Example requests:

`curl -vvv -X PUT localhost:8000/kv -d '{"key": "org/name", "value": "ishan"}'`

`curl -vvv -X GET localhost:8000/kv -d '{"key": "org/name"}'`

`curl -vvv -X DELETE localhost:8000/kv -d '{"key": "org/name"}'`

