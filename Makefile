PACKAGE := urawesome
GOPATH := $(PWD)
export GOPATH

# VERSION := 0.0.0
# BUILD := `git rev-parse HEAD`
# LDFLAGS += -ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
LDFLAGS := 

clean:
	rm -rfI build
run:
	go run $(LDFLAGS) main.go

build/%.a: $(wildcard src/%/*.go)
	go build -o $@ $(LDFLAGS) $<

bin/%: %.go
	go build -o $@ $(LDFLAGS) $<

ur: build/urawesome/ur.a

main: bin/main ur

