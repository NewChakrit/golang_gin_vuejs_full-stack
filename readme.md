curl -X GET http://127.0.0.1:80/api/txn       
curl -X POST http://127.0.0.1:80/api/txn/add -d '{"type":"buy", "ticker":"amzn", "volume":1.25, "price":166.90, "date":"2025-07-01"}
curl -X POST http://127.0.0.1:80/api/txn/update -d '{"id": 1,"type":"buy", "ticker":"amzn", "volume":2.75, "price":166.90, "date":"2025-07-01"}'