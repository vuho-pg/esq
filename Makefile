gen:
	touch tmp.gen.go && rm -f $$(ls | grep .gen.go) && cd codegen && go run *.go && cd .. && go fmt *.go

rem:
	touch tmp.gen.go && rm -f $$(ls | grep .gen.go)