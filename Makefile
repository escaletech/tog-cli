TEST ?= ./...

GOCMD=$(if $(shell which richgo),richgo,go)

install:
	go install ./cmd/tog

test:
	ENV=test $(GOCMD) test -v $(TEST)

test-watch:
	reflex -s --decoration=none -r \.go$$ -- make test TEST=$(TEST)
	ENV=test $(GOCMD) test -v $(TEST)

release:
	@bash -c "$$(curl -s https://raw.githubusercontent.com/escaletech/releaser/master/tag-and-push.sh)"
