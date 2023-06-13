
gen:
	cd codegen && go run *.go && cd .. && go fmt *.go