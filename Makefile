DIR_NAME=dist

build: ${DIR_NAME}
	go mod tidy
	go build -o ${DIR_NAME}/rename-files

clean:
	@rm -rf ${DIR_NAME}

${DIR_NAME}:
	@mkdir -p ${DIR_NAME}
