curl -s -X "POST" "http://localhost:8080/rpc" -H "Content-Type: application/json; charset=utf-8" -d '{
    "jsonrpc": "2.0",
    "method": "Service.CreateObject",
    "id": "1",
    "params": [
        {
            "object": {
                "name": "New Object",
                "status": "CREATED",
                "description": "A new object"
            }
        }
    ]
}' |jq .
