package main

import "github.com/AndrewJDGalga/keep_awake/keepawake"

func main() {
	awake := keepawake.New()
	if err := awake.Start(); err != nil {
		panic(err)
	}

}
