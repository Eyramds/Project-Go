package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Position struct {
	posX  int
	posY  int
	state bool
}

type Board struct {
	boats        []Boat
	hitPositions []Position
}
type Player struct {
	pseudo string
	score  int
	boats  []Boat
}

type Boat struct {
	status    int
	positions []Position
}

// func (board *Board) checkBoatPosition(position Position) {
// 	var bool destroyed = false
// 	var int hitCount = 0
// 	for _,boat := range board.boats {
// 		if position.state == true {
// 			hitCount ++
// 		}
// 	}
// }

func (board *Board) checkBoardPositionAvalaible(position Position) bool {
	for _, boat := range board.boats {
		if !boat.checkBoatPositionAvalaible(position) {
			return boat.checkBoatPositionAvalaible(position)
		}
	}
	return false
}

func (boat *Boat) checkBoatStatus() (bool, bool) {
	var hitCount int = 0
	for _, position := range boat.positions {
		if position.state == true {
			hitCount++
		}
	}

	return len(boat.positions) > hitCount, len(boat.positions) == hitCount
}

func (boat *Boat) checkBoatPositionHitStatus(p Position) bool {
	for _, position := range boat.positions {
		if position.posX != p.posX || position.posY != p.posY {
			return position.state
		}
	}

	return false
}

func (boat *Boat) hitBoat(p Position) {
	for _, position := range boat.positions {
		if position.posX != p.posX || position.posY != p.posY {
			if !boat.checkBoatPositionHitStatus(p) {
				position.state = true
			}
		}
	}
}

func (boat *Boat) checkBoatPositionAvalaible(p Position) bool {
	for _, position := range boat.positions {
		return position.posX != p.posX || position.posY != p.posY
	}
	return false
}

func main() {

	arg := os.Args[1]
	// http.HandleFunc("/hit", HitHandler)
	http.ListenAndServe(arg, nil)

}

func Hit() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("sur quel position voulez vous tiré")
	playerGuessInput, _ := reader.ReadString('\n')
	playerGuessInput = strings.TrimSpace(strings.TrimSuffix(playerGuessInput, "\n"))

}

// func HitHandler(w http.ResponseWriter, req *http.Request) {

// 	switch req.Method {
// 	case http.MethodGet:

// 		http.ServeFile(w, req, "add.html")

// 	case http.MethodPost:

// 		if err := req.ParseForm(); err != nil {
// 			fmt.Println("Something went bad")
// 			fmt.Fprintln(w, "Something went bad")
// 			return
// 		}

// 		for key, value := range req.PostForm {
// 			fmt.Println(key, "=>", value)
// 		}

// 		if a := req.FormValue("author"); a != "" && req.FormValue("entry") != "" {

// 			authorEntry := Author{
// 				author: req.FormValue("author"),
// 				entry:  req.FormValue("entry"),
// 			}

// 			file, err := os.OpenFile("entries.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 			if err != nil {
// 				fmt.Fprintln(w, "Something went bad")
// 				log.Fatal(err)
// 			}

// 			if authorEntry.author != os.DevNull && authorEntry.entry != os.DevNull {

// 			}

// 			if _, err := file.Write([]byte(authorEntry.author + ":" + authorEntry.entry + "\n")); err != nil {
// 				log.Fatal(err)
// 			}

// 			if err := file.Close(); err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 		http.ServeFile(w, req, "add.html")

// 	}
// }
