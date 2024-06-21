projectName = "example"

# build with verison infos
versionDir = "github.com/erpc-go/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
gitBranch = $(shell git symbolic-ref --short -q HEAD)
buildAuthor = $(shell if [ "`git config user.name`" != "" ]; then git config user.name; else hostname; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
goVersion = $(shell go version | awk -F" " '{print $$3}')
platform = $(shell go version | awk -F" " '{print $$4}')

ldflags="-w -X ${versionDir}.name=${projectName} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.gitCommit=${gitCommit} \
		-X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.buildAuthor=${buildAuthor} \
		-X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.goVersion=${goVersion} \
		-X ${versionDir}.platform=${platform} -X ${versionDir}.gitBranch=${gitBranch}"

all: gotool
	@go build -o ${projectName} -v -ldflags ${ldflags} .

clean:
	rm -f ${projectName}

gotool:
	gofmt -w .
	go vet .

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"

.PHONY: clean gotool help jce test bench
