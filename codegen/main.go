package main

func main() {
	if err := enumBuild(); err != nil {
		panic(err)
	}
	if err := interfaceBuild(); err != nil {
		panic(err)
	}
	//compound
	if err := compoundBuild(); err != nil {
		panic(err)
	}
	// fulltext
	if err := fullTextQueryBuild(); err != nil {
		panic(err)
	}
	//geo
	//shape
	//joining
	if err := joiningBuild(); err != nil {
		panic(err)
	}
	//match all
	//span
	//specialized
	//term level
	if err := termBuild(); err != nil {
		panic(err)
	}
}
