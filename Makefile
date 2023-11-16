build:
	go build -o bin/rmq .

producer: build
	./bin/rmq -producer=true -qname=testing -msg="test message"

consumer: build
	./bin/rmq -producer=false -qname=testing -msg="test message"