# SnappFood Test
 ### Running The Project
Just run the docker-compose using the following command:
```
docker-compose up -d
```

 ### Usage
 You can use the Postman collection provided.

 ***[Postman Collection For Snappfood Test](./snapfood_test.postman_collection.json)***
 
 ### Details

| Route                       | Method | Header | Body                    | Response           | Description                                                |
|-----------------------------|--------|--------|-------------------------|--------------------|------------------------------------------------------------|
| {{addr}}/order/             | POST   | -      | OrderRequest            | Order              | creates a new order                                        |
| {{addr}}/order/delay        | POST   | -      | OrderDelayReportRequest | Order              | submits a request for delay                                |
| {{addr}}/order/vendor/delay | GET    | -      | -                       | VendorDelayReport  | gives the last week's report on vendors delays             |
| {{addr}}/trip/              | POST   | -      | TripRequest             | Trip               | creates a trip for the order                               |
| {{addr}}/trip/              | PUT    | -      | TripUpdateRequest       | Trip               | updates the status of the trip                             |
| {{addr}}/agent/task         | Get    | -      | -                       | AgentHistory       | gives the agent a new task if he does not have any ongoing |
