# golambda

* install sam local: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html
* sam init --runtime go
* go get ./...
* make build
* sam local generate-event s3 > s3event.json
* sam local invoke list-buckets
* make build && sam local invoke "listbuckets" -e s3event.json