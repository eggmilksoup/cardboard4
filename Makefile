all:
	mkdir -p bin data
	${MAKE} bin

bin: bin/logdms          \
     bin/logmentions     \
     bin/msg             \
     bin/nickname        \
     bin/username

bin/logdms: src/logdms.go
	go build -o bin/logdms src/logdms.go

bin/logmentions: src/logmentions.go
	go build -o bin/logmentions src/logmentions.go

bin/msg: src/msg.go
	go build -o bin/msg src/msg.go

bin/nickname: src/nickname.go
	go build -o bin/nickname src/nickname.go

bin/username: src/username.go
	go build -o bin/username src/username.go

clean:
	rm bin/*
