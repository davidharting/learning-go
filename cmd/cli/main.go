package main

import (
	"fmt"

	"github.com/davidharting/learngo/git"
)

func main() {
	fmt.Println("hello world")

	repo, _ := git.NewRepo(".")
	branch, _ := repo.CurrentBranch()
	fmt.Printf("Current branch: %v\n", branch)

	msg, _ := repo.LatestCommit()
	fmt.Printf("\nLatest commit message:\n%v\nAuthor: %v\n", msg.Message, msg.Author)
}
