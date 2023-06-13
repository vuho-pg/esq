package main

func main() {
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
	//match all
	//span
	//specialized
	//term level
}
