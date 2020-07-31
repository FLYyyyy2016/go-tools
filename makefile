.PHONY: test

test:
		go test ./api/
		go test ./tools/fflag/
		go test ./tools/http/
		go test ./tools/schedule/
		go test ./tools/validate/