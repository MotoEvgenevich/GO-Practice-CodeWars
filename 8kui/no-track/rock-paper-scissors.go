package main

import "fmt"

func main() {
	fmt.Println(Rps("rock", "scissors"))
}

func Rps(p1, p2 string) string {
	if p1 == "rock" && p2 == "rock" {
		return "Draw!"
	}
	if p1 == "scissors" && p2 == "scissors" {
		return "Draw!"
	}
	if p1 == "paper" && p2 == "paper" {
		return "Draw!"
	}
	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	}
	if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	}
	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	}
	if p1 == "scissors" && p2 == "rock" {
		return "Player 2 won!"
	}
	if p1 == "rock" && p2 == "paper" {
		return "Player 2 won!"
	}
	if p1 == "paper" && p2 == "scissors" {
		return "Player 2 won!"
	}
	return "Draw!"
}

func RpsCase(p1, p2 string) string {
	switch 
}
