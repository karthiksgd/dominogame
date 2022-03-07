package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Boneyard or Pile
var Tile = [][]int{
	{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6},
	{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
	{2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6},
	{3, 3}, {3, 4}, {3, 5}, {3, 6},
	{4, 4}, {4, 5}, {4, 6},
	{5, 5}, {5, 6},
	{6, 6},
}

// Players Listing
var players []string

// Player Deck
var deck map[string][][]int

// Unique Random Check
var tileStatus = make(map[int]bool)

// Boneyard Check all double Input Tiles
var doubles [][]int

// Boneyard Check all double Input Tiles
var bigDouble = make(map[string][]int)

// DoublyList Array
var doublyPrintList [][]int

// Tile Struct
type tile struct {
	data []int
	prev *tile
	next *tile
}

// DoubleLinked List
type doublyLinkedList struct {
	len  int
	tail *tile
	head *tile
}

func main() {

	fmt.Println("Welcome to Domino Game")

	fmt.Println("Pile : ", Tile)

	//Define Players

	////Number of Players Playing
	fmt.Println("Input no: of Player Playing ? ")

	read_Players := bufio.NewReader(os.Stdin)
	conv_Players, _ := read_Players.ReadString('\n')
	no_Players, _ := strconv.Atoi(strings.TrimSpace(conv_Players))

	////Get Name of those who are playing
	for i := 0; i < no_Players; i++ {

		fmt.Printf("Input Player %v Name : ", i+1)

		read_Players = bufio.NewReader(os.Stdin)
		conv_Players, _ = read_Players.ReadString('\n')
		players = append(players, conv_Players)

	}

	// Set Status for Tiles
	for i := 0; i < 28; i++ {
		tileStatus[i] = true
	}

	// Distribute 7 Tiles to Each Players
	deck = make(map[string][][]int)

	for _, val := range players {

		////Get Random Tiles
		for i := 0; i < 7; i++ {

			for {

				randNo := getRandomInt()

				if tileStatus[randNo] {

					deck[val] = append(deck[val], Tile[randNo])
					tileStatus[randNo] = false
					break

				} else {

					tileStatus[randNo] = false
				}
			}

		}
	}

	// Each Player's Deck
	for k, val := range deck {

		fmt.Printf("\n%v\n-----------\n-> %v \n\n", k, val)
	}

	// Find Biggest Doubles from each Players
	for _, val := range deck {

		for i := 0; i < 7; i++ {

			if val[i][0] == val[i][1] {

				doubles = append(doubles, val[i])

			}
		}
	}

	// Find the Biggest Double

	BigD := [][]int{{0, 0}}

	for _, val := range doubles {

		if val[0] > BigD[0][0] {

			BigD[0] = val
		}
	}

	// Player holding biggest Double

	for k, val := range deck {

		for i := 0; i < 7; i++ {

			if val[i][0] == BigD[0][0] && val[i][1] == BigD[0][1] {
				bigDouble[k] = val[i]
			}
		}
	}

	// Player with Biggest Double
	starting_index := 0

	for k, val := range bigDouble {
		fmt.Printf("\nPlayer with Highest Double is %v with Value %v \n", k, val)

		opener := k

		for index, val1 := range players {
			if val1 == opener {
				starting_index = index
			}
		}
	}

	// Playtime
	fmt.Printf("\n\n <--------- Let's Play ---------> \n\n")

	//--------------------------------------------------------------------------------------------------------------------------------//
	// Player with Biggest Double starts first

	doublyList := initDoublyList()

	starter := starting_index + 1
	fmt.Println(players[starting_index])

	fmt.Printf("Starter Tile : %v\n", BigD)
	doublyList.AddLeftNode(BigD[0])

	fmt.Print("End to End Domino Chain : <- ")

	// Update DoublyPrintList
	err := doublyList.PrintList()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(" -> ")

	// Remove Tile from Opener
	for index, val := range deck[(players[starting_index])] {
		if val[0] == BigD[0][0] && val[1] == BigD[0][1] {

			deck[(players[starting_index])] = append(deck[(players[starting_index])][:index], deck[(players[starting_index])][index+1:]...)
		}
	}

	for {

		if starter == len(players) {
			starter = 0
		}

		// Next Player
		fmt.Printf("\n\nNext Player is %v\n", players[starter])
		fmt.Printf("\n\nPlayer Deck Consist of : %v\n\n", deck[(players[starter])])

		var tilePick int
		var tilePlacement string

		// List Tiles with Position and Placement
		for i, val := range deck[(players[starter])] {
			fmt.Println("Enter ", i+1, "to Pick Tile ", val)

		}

		fmt.Println("Enter 0 if you wanna pass the turn.")
		fmt.Println("Your Pick is :")
		fmt.Scanln(&tilePick)

		if tilePick == 0 {
			starter++
			fmt.Printf("\nPlayer %v has passed to next player \n", players[starter])
			continue
		}

		// Placement
		fmt.Println("Where you wanna  place the tile (r/l) ? ")
		fmt.Scanln(&tilePlacement)

		// Place the Tile in the Domino Chain
		if tilePlacement == "r" {

			doublyList.AddRightNode(deck[(players[starter])][tilePick-1])

			fmt.Print("End to End Domino Chain : <- ")

			// Update DoublyPrintList
			err := doublyList.PrintList()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Print(" -> ")

		} else if tilePlacement == "l" {

			doublyList.AddLeftNode(deck[(players[starter])][tilePick-1])

			fmt.Print("End to End Domino Chain : <- ")

			// Update DoublyPrintList
			err := doublyList.PrintList()
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Print(" -> ")

		}

		// Remove the Tile from Players Deck

		deck[(players[starter])] = append(deck[(players[starter])][:tilePick-1], deck[(players[starting_index])][tilePick:]...)

		// Check whether Tiles runs out

		// Increment the Starter
		starter++

		// break
	}

	// fmt.Printf("Size of doubly linked ist: %d\n", doublyList.Size())

}

// Get Randon Integer
func getRandomInt() int {

	randNo, _ := rand.Int(rand.Reader, big.NewInt(28))

	return int(randNo.Int64())
}

// Initialize Empty List
func initDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}

// Add Tile to Right End
func (d *doublyLinkedList) AddLeftNode(data []int) {
	newNode := &tile{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	// return
}

// Add Tile to Left End
func (d *doublyLinkedList) AddRightNode(data []int) {
	newNode := &tile{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	// return
}

// Print Tile on Boneyard
func (d *doublyLinkedList) PrintList() error {
	if d.head == nil {
		return fmt.Errorf("PrintError: List is empty")
	}
	temp := d.head
	for temp != nil {
		// fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		// doublyPrintList = append(doublyPrintList, temp.data)
		fmt.Print(temp.data)
		// fmt.Print("<->")
		temp = temp.next
	}

	return nil
}

func (d *doublyLinkedList) Size() int {
	return d.len
}
