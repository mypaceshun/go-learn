cmd      = bin/shun-command

build:
	${MAKE} ${cmd}

clean:
	rm -v ${cmd}

format:
	go fmt cmd/*
	go fmt log/*

${cmd}:
	go build -o ${cmd} cmd/*
