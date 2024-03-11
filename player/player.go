// The player package is used to set and get player data
package player

import (
	"bufio"
	"fmt"
	"log"
)

// The PlayerFieldLimit is given as a security measure for now. It should probably
// be replaced with a better solution.
const (
	PlayerFieldMaxCharacters = 50
	PlayerFieldMinCharacters = 1
)

// PlayerInfo holds the player information
type PlayerInfo struct {
	Name  string
	Loves string
	Fears string
}

// Make a new player and return it
func Make(s *bufio.Scanner) PlayerInfo {

	var name string
	var loves string
	var fears string
	// loop until the user provides the correct length for each field.
	for {
		ans, err := getPlayerInfo(s)

		if err != nil {
			log.Fatal("error shouldn't happen on getPlayerInfo")
		}

		name = ans[0]
		loves = ans[1]
		fears = ans[2]

		// check if too long
		if len(name) > PlayerFieldMaxCharacters || len(loves) > PlayerFieldMaxCharacters || len(fears) > PlayerFieldMaxCharacters {
			// Create custom error to return here
			continue
		}

		// check if any fields are empty
		if len(name) < PlayerFieldMinCharacters || len(loves) < PlayerFieldMinCharacters || len(fears) < PlayerFieldMinCharacters {
			// Create custom error to return here
			continue
		}
		break
	}

	return PlayerInfo{
		Name:  name,
		Loves: loves,
		Fears: fears,
	}
}

// GetPlayerInfo returns a PlayerInfo
// TODO: When would this return an error?
func (p PlayerInfo) GetPlayerInfo() (PlayerInfo, error) {
	return p, nil
}

func getPlayerInfo(s *bufio.Scanner) ([3]string, error) {

	var count int
	var ans [3]string
	fmt.Println("What is your name?")
	fmt.Print("> ")
	if s.Scan() {
		ans[0] = s.Text()
	}

	for {
		switch count {
		case 0:
			fmt.Println("What is/are your greatest loves?")
			fmt.Print("> ")
			if s.Scan() {
				ans[1] = s.Text()
			}
			count++
		case 1:
			fmt.Println("What is/are your greatest fears?")
			fmt.Print("> ")
			if s.Scan() {
				ans[2] = s.Text()
			}
			count++
		default:
			return ans, nil
		}
	}
}
