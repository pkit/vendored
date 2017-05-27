CONTAINER=nickg/misspell

.git/hooks/pre-commit: scripts/pre-commit.sh
	cp -f scripts/pre-commit.sh .git/hooks/pre-commit
.git/hooks/commit-msg: scripts/commit-msg.sh
	cp -f scripts/commit-msg.sh .git/hooks/commit-msg

hooks: .git/hooks/pre-commit .git/hooks/commit-msg

build: hooks
	go install ./cmd/misspell
	gometalinter \
		 --vendor \
		 --deadline=60s \
	         --disable-all \
		 --enable=vet \
		 --enable=golint \
		 --enable=gofmt \
		 --enable=goimports \
		 --enable=gosimple \
		 --enable=staticcheck \
		 --enable=ineffassign \
		 --exclude=/usr/local/go/src/net/lookup_unix.go \
		 ./...
	go test .

test:
	go test .

# the grep in line 2 is to remove misspellings in the spelling dictionary
# that trigger false positives!!
falsepositives: /scowl-wl
	cat /scowl-wl/words-US-60.txt | \
		grep -i -v -E "payed|Tyre|Euclidian|nonoccurence|dependancy|reenforced|accidently|surprize|dependance|idealogy|binominal|causalities|conquerer|withing|casette|analyse|analogue|dialogue|paralyse|catalogue|archaeolog|clarinettist|catalyses|cancell|chisell|ageing|cataloguing" | \
		misspell -debug -error
	cat /scowl-wl/words-GB-ise-60.txt | \
		grep -v -E "payed|nonoccurence|withing" | \
		misspell -locale=UK -debug -error
#	cat /scowl-wl/words-GB-ize-60.txt | \
#		grep -v -E "withing" | \
#		misspell -debug -error
#	cat /scowl-wl/words-CA-60.txt | \
#		grep -v -E "withing" | \
#		misspell -debug -error

bench:
	go test -bench '.*'

clean:
	rm -rf dist/ bin/
	go clean ./...
	git gc --aggressive

ci:
	docker run --rm \
		-v $(PWD):/go/src/github.com/client9/misspell \
		-w /go/src/github.com/client9/misspell \
		${CONTAINER} \
		make build falsepositives

docker-build:  ## build a test test image
	docker build -t ${CONTAINER} .

docker-pull:  ## pull latest test image
	docker pull ${CONTAINER}

docker-console:  ## log into the test image
	docker run --rm -it \
		-v $(PWD):/go/src/github.com/client9/misspell \
		-w /go/src/github.com/client9/misspell \
		${CONTAINER} sh

.PHONY: ci console docker-build bench
