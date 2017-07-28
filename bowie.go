package main

import (
	"fmt"
	"os"
	"os/exec"
)

var upstream = "upstream"

func main() {
	fetch(upstream)

	checkout("master")
	resetHard(upstream, "master")

	argsWithoutProg := os.Args[1:]

	if (len(argsWithoutProg) == 0) {
		printIncorrectUsage("Commit SHA to cherry-pick must be specified")
		return;
	}

	commitIsh := argsWithoutProg[0]

	if (len(argsWithoutProg) == 1) {
		printIncorrectUsage("Branches to cherry-pick to must be specified")
		return;
	}

	for _, branch := range argsWithoutProg[1:] {
		checkout(branch)
		resetHard(upstream, branch)
		cherryPick(commitIsh)
		push(upstream)
	}
}

func printIncorrectUsage(message string) {
	fmt.Println(message)
	fmt.Println()
	fmt.Println("usage: bowie <commit-ish> <branch1> [<branch2>]")
	fmt.Println("example usage: bowie a1b2c3d4 6.x 6.0")
}

func fetch(remote string) {
	cmd := "git"
	args := []string{"fetch", remote}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func checkout(branch string) {
	cmd := "git"
	args := []string{"checkout", branch}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cherryPick(commitIsh string) {
	cmd := "git"
	args := []string{"cherry-pick", commitIsh}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func push(remote string) {
	cmd := "git"
	args := []string{"push", remote}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func resetHard(remote string, branch string) {
	cmd := "git"
	args := []string{"reset", "--hard", fmt.Sprintf("%s/%s", remote, branch)}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}