all:
	mkdir -p bin data
	${MAKE} bin

bin: bin/logmentions     \
     bin/msg             \
     bin/nickname

bin/logmentions: src/logmentions.go
	go build -o bin/logmentions src/logmentions.go

bin/msg: src/msg.go
	go build -o bin/msg src/msg.go

bin/nickname: src/nickname.go
	go build -o bin/nickname src/nickname.go

clean:
	rm bin/*
