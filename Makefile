gen:
	touch tmp.go && rm -f $$(ls | grep .go | grep -v test.go) && cd codegen && go run *.go && cd .. && go fmt *.go