testShort:
	go test -count=1 -tags longTests ./...

testAll:
	go test -count=1 ./...
