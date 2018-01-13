default: clean
	@export GOPATH=$$GOPATH:$$(pwd) && go install runner
edit:
	@export GOPATH=$$GOPATH:$$(pwd) && atom .
edit2:
	@export GOPATH=$$GOPATH:$$(pwd) && code .
run: default
	@bin/runner
	@echo ""
clean:
	@rm -rf bin
setup:
	go get -u github.com/aws/aws-sdk-go/...
