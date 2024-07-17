.PHONY: .test
.test:
	$(info Running tests...)
	go test ./sprint-2/...

.PHONY: test
test: .test ## run unit tests