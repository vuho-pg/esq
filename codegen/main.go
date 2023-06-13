package main

func main() {
	if err := fullTextQueryBuild(); err != nil {
		panic(err)
	}
}
