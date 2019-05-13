.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./listbuckets/list-buckets
	
build:
	GOOS=linux GOARCH=amd64 go build -o listbuckets/listbuckets ./listbuckets

run:
	sam local invoke listbuckets -e s3event.json