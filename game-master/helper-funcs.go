package gamemaster

import (
	"math/rand"
)

// Generates a random musical key for a scale or chord and returns its name
func getRandomScaleOrChord() string {
	// Define slices of musical keys and scale/chord types
	keys := []string{"C", "D", "E", "F", "G", "A", "B", "C#", "D#", "F#", "G#", "A#"}
	scales := []string{"Major", "Melodic Minor", "Harmonic Minor", "Natural Minor", "Dorian", "Phrygian", "Lydian", "Mixolydian", "Aeolian", "Locrian"}
	// Alternatively, for chords, you could use something like:
	chords := []string{"Maj7", "min7", "7", "min7b5", "dim7", "Maj9", "min9", "9"}

	// Randomly select an element from each slice
	key := keys[rand.Intn(len(keys))]
	scale := scales[rand.Intn(len(scales))]
	chord := chords[rand.Intn(len(chords))]

	// Return the concatenated string
	if rand.Intn(2)%2 == 0 {
		return key + " " + scale
	}
	return key + " " + chord
}
