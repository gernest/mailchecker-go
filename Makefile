
test:
	@go test
	
deps:
	@go get -u github.com/jteeuwen/go-bindata/...

bindata:
	@go-bindata -pkg="mailchecker" list.json