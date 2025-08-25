APP_NAME = btool
TARGET = ./bin/${APP_NAME}

run: build
	${TARGET}

build:
	go build -o ${TARGET}

install: build
	chmod +x ./install.sh
	./install.sh

	
