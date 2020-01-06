# gamezop

### Build and run queuing-service
```bash
export SQSAPIVersion='2012-11-05'
export SQSQueueURL='https://sqs.us-east-2.amazonaws.com/*******/gamezop'
export AWSRegion=us-east-2
export MONGO_CONNECT_STRING="mongodb://localhost/test"
export REDIS_HOST='localhost'
export REDIS_PORT=6379


#### build it and run it 
go build -o ./bin/queue -i ./cmd/queue/ && ./bin/queue
```


### Build and run http-service
```bash
export AWSRegion=us-east-2
export MONGO_CONNECT_STRING="mongodb://localhost/test"
export HTTP_PORT=8000

#### build it and run it 
go build -o ./bin/http -i ./cmd/http/ && ./bin/http
```
