// The player package is used to set and get player data
package player

import (
	"errors"
	"fmt"
)

// The PlayerFieldLimit is given as a security measure for now. It should probably
// be replaced with a better solution.
const (
	PlayerFieldLimit = 50
)

// PlayerInfo holds the player information
type PlayerInfo struct {
	name  string
	loves string
	fears string
}

func Make(name, loves, fears string) (PlayerInfo, error) {
	if len(name) > PlayerFieldLimit || len(loves) > PlayerFieldLimit || len(fears) > PlayerFieldLimit {
		return PlayerInfo{}, errors.New(fmt.Sprintf("Player inputs are too long. Only %d characters are allowed for each field", PlayerFieldLimit))
	}
	return PlayerInfo{
		name:  name,
		loves: loves,
		fears: fears,
	}, nil
}

func (p PlayerInfo) GetPlayerInfo() (PlayerInfo, error) {
	return p, nil
}
