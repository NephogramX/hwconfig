.PHONY: run

run:
	go env -w GOOS=windows
	go env -w GOARCH=amd64
	go run ./main.go

clean:
	rm build -r