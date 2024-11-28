package main

import "go-every/FunctionalOptions/file"

func main() {
	err := file.New("./tmp/empty.txt")
	if err != nil {
		panic(err)
	}

	err = file.New(
		"./tmp/file.txt",
		file.WithUID(1000),
		file.WithContents("Lorem Ipsum Dolor Amet"),
	)

	if err != nil {
		panic(err)
	}
}
