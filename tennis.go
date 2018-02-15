package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//Player is used as a goroutine in a "tennis" match,
//if the ball is missed the game is over for all goroutines,
//this is done by the closing the clannel otherwise we increment the score
//and "hit back" to the other goroutine.
func Player(name string, score chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		count, open := <-score
		//if channel is closed finish the game
		if !open {
			//The other player closed the channel so this player wins
			fmt.Printf("Player %v won the game \n", name)
			return
		}

		//Create a change to miss the ball and end the game
		//by closing the channel.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %v missed. \n", name)
			close(score)
			return
		}
		fmt.Printf("Player %v hit the ball, the score is %v \n", name, count)
		count++
		score <- count
	}
}
