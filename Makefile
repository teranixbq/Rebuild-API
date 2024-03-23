
test:
# running on directory
	go test -v -cover

cover:
# running on directory with view out
	go test -coverprofile=cover.out && go tool cover -html=cover.out

coveralls:
# running on features directory
	go list ./... | grep service | xargs -n1 go test -cover