admin=\$$avian \$$ghost \$$yngling

all:
	mkdir -p bin data
	${MAKE} bin speaker

bin: bin/logdms      \
     bin/logmentions \
     bin/mention     \
     bin/mkthread    \
     bin/msg         \
     bin/nickname    \
     bin/readmsg     \
     bin/username

speaker: speaker/listen.sh

bin/logdms: src/logdms.go
	go build -o bin/logdms src/logdms.go

bin/logmentions: src/logmentions.go
	go build -o bin/logmentions src/logmentions.go

bin/mention: src/mention.go
	go build -o bin/mention src/mention.go

bin/mkthread: src/mkthread.go
	go build -o bin/mkthread src/mkthread.go

bin/msg: src/msg.go
	go build -o bin/msg src/msg.go

bin/nickname: src/nickname.go
	go build -o bin/nickname src/nickname.go

bin/readmsg: src/readmsg.go
	go build -o bin/readmsg src/readmsg.go

bin/username: src/username.go
	go build -o bin/username src/username.go

speaker/listen.sh: speaker/listen.sh.in
	sed "s/\$$admin/${admin}/" speaker/listen.sh.in > speaker/listen.sh

clean:
	rm bin/* speaker/listen.sh
