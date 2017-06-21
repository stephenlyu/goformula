# Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# blame: jnml, labs.nic.cz

all: easylang.go tokenizer.go context.go expression.go
	gofmt -l -s -w *.go
	go build $^

run: all
	./easylang

tokenizer.go: tokenizer.l
	golex -o $@ $<

easylang.go: easylang.y
	goyacc -o $@ $<

clean:
	go clean
	rm -f tokenizer.go easylang.go lex.yy.go y.go y.output *~

nuke: clean
	rm -f easylang
