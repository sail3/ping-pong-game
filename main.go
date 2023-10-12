package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	var b = initializeBoard()
	fmt.Println(b)

	if err := keyboard.Open(); err != nil {
		log.Fatal("error: cannot open keyboard: ", err)
	}

	defer keyboard.Close()

	runeReceiverChan := make(chan rune, 1)
	keyReceiverChan := make(chan keyboard.Key, 1)
	timer := time.NewTicker(time.Millisecond * 150)

	go func() {
		for {
			r, key, err := keyboard.GetSingleKey()
			if err == nil {
				runeReceiverChan <- r
				keyReceiverChan <- key
			}
		}
	}()

	for {
		select {
		case <-timer.C:
			clearConsole()
			b.MoveBall()
			if b.hasWinner {
				fmt.Printf("%s Win the Game!!\n%s", b.turn, b)
				return
			}
			fmt.Printf("PlayerA use keys (W) and (S) || PlayerB use keys (up) and (down)\n%s\n", b)
		case r := <-runeReceiverChan:
			k := <-keyReceiverChan
			if r == 'q' {
				os.Exit(0)
			}
			b.HandlePlayerB(k)
			b.HandlePlayerA(r)
		}
	}
}
