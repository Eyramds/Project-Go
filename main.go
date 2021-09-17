package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// var ports []string

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

var playerBoard Board
var player Player

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

func (board *Board) hitBoard(p Position) bool {
	for _, boat := range board.boats {
		if boat.hitBoat(p) {
			return true
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

func (boat *Boat) hitBoat(p Position) bool {
	for _, position := range boat.positions {
		if position.posX != p.posX || position.posY != p.posY {
			if !boat.checkBoatPositionHitStatus(p) {
				position.state = true
				return true
			}
		}
	}
	return false
}

func (boat *Boat) checkBoatPositionAvalaible(p Position) bool {
	for _, position := range boat.positions {
		return position.posX != p.posX || position.posY != p.posY
	}
	return false
}

func main() {

	http.HandleFunc("/hit", HitHandler)
	port := os.Args[2]
	username := os.Args[1]
	go addPort(username, port)
	err := http.ListenAndServe(":"+os.Args[1], nil)
	fmt.Println(err)
	start()
}

func start() {
	reader := bufio.NewReader(os.Stdin)

	file, err := os.Open("ports.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var playerInfo []string

	for scanner.Scan() {
		playerInfo = append(playerInfo, scanner.Text())
	}

	file.Close()

	fmt.Println("sur quel joueur souhaiez vous tiré ? Selectionné un port")

	for _, player := range playerInfo {
		fmt.Println("%s\n", player)
	}

	selectedPort, _ := reader.ReadString('\n')
	selectedPort = strings.TrimSpace(strings.TrimSuffix(selectedPort, "\n"))

	fmt.Println("sur quel position voulez vous tiré sur L'axe X")

	posX, _ := reader.ReadString('\n')

	fmt.Println("sur quel position voulez vous tiré sur L'axe Y")

	posY, _ := reader.ReadString('\n')

	posX = strings.TrimSpace(strings.TrimSuffix(posX, "\n"))
	posY = strings.TrimSpace(strings.TrimSuffix(posY, "\n"))

	data := url.Values{
		"posX": {posX},
		"posY": {posY},
	}

	resp, err := http.PostForm("http://localhost:"+selectedPort+"/hit", data)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("%v\n", resp.Body)
	// resp.Body.Read("etat")

}

func HitHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:

		// http.ServeFile(w, req, "add.html")

	case http.MethodPost:

		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}

		// for key, value := range req.PostForm {
		// 	fmt.Println(key, "=>", value)
		// }

		if a := req.FormValue("posX"); a != "" && req.FormValue("posY") != "" {

			x, _ := strconv.Atoi(req.FormValue("posX"))
			y, _ := strconv.Atoi(req.FormValue("posY"))

			position := Position{
				posX: x,
				posY: y,
			}

			if playerBoard.hitBoard(position) {
				fmt.Fprintln(w, "position: %d %d toucher", position.posX, position.posY)
			}

			fmt.Fprintln(w, "position: %d %d toucher", position.posX, position.posY)

			// http.ServeFile(w, req, "add.html")
		}
	}
}

func addPort(player string, port string) {

	file, err := os.OpenFile("ports.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.Write([]byte(player + ":" + port + "\n")); err != nil {
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
