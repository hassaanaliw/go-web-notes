tests:
	cd test && go test -v -race
run:
	go run main.go config.go