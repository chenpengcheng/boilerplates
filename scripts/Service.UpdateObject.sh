curl -s -X "POST" "http://localhost:8080/rpc" -H "Content-Type: application/json; charset=utf-8" -d '{
    "jsonrpc": "2.0",
    "method": "Service.UpdateObject",
    "id": "1",
    "params": [
        {
            "object": {
                "id": 1,
                "name": "Updated Object",
                "status": "UPDATED",
                "description": "An updated object"
            }
        }
    ]
}' |jq .
