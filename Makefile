run: bin/heroku-go-example
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/heroku-go-example: main.go
	go build -o bin/heroku-go-example main.go

clean:
	rm -rf bin
