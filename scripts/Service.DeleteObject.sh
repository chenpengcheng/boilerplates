curl -s -X "POST" "http://localhost:8080/rpc" -H "Content-Type: application/json; charset=utf-8" -d '{
    "jsonrpc": "2.0",
    "method": "Service.DeleteObject",
    "id": "1",
    "params": [
        {
            "id": 1
        }
    ]
}' |jq .
