package scale

import (
	"container/ring"
	"strings"
)

type Tonic struct {
	note string
	sharp bool
	flat bool
}
type Interval string
type ScaleType []string

// Assumes only valid tonics and intervals are provided
func Scale(tonic string, interval string) ScaleType {
	note := Tonic{note: tonic}
	var scaleLength int

	if interval == "" {
		scaleLength = 12
	} else {
		scaleLength = len(interval)
	}

	notesSharp := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	sharpNotesRing := ring.New(len(notesSharp))
	for i := 0; i < sharpNotesRing.Len(); i++ {
		sharpNotesRing.Value = notesSharp[i]
		sharpNotesRing = sharpNotesRing.Next()
	}

	notesFlat := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	flatNotesRing := ring.New(len(notesFlat))
	for i := 0; i < flatNotesRing.Len(); i++ {
		flatNotesRing.Value = notesFlat[i]
		flatNotesRing = flatNotesRing.Next()
	}

	//choose ring based on tonic
	var noteRing *ring.Ring
	if note.isAccidental() {
		if note.flat {
			noteRing = flatNotesRing
		} else {
			noteRing = sharpNotesRing
		}
	} else {
		switch note.note {
		case
			"C", "a", "G","D", "A", "E", "B", "e", "b":
			noteRing = sharpNotesRing
		default:
			noteRing = flatNotesRing
		}
	}

	//spin ring until we get to the correct starting point
	for i := 0; i < noteRing.Len(); i++ {
		if strings.Title(note.note) == noteRing.Value.(string) {
			break
		}
		noteRing = noteRing.Next()
	}

	scale := ScaleType{}
	scale = append(scale, noteRing.Value.(string))

	if scaleLength == 12 {
		for i := 0; i < 11; i++ {
			noteRing = noteRing.Next()
			scale = append(scale, noteRing.Value.(string))
		}

	} else {
		steps := strings.Split(interval, "")
		// for x in interval
		for _, step := range steps[:len(steps)-1] {
			var rotations int
			switch step {
			case "m":
				rotations = 1
			case "M":
				rotations = 2
			case "A":
				rotations = 3
			}

			for i := 0; i < rotations; i++ {
				noteRing = noteRing.Next()
			}
			scale = append(scale, noteRing.Value.(string))
		}
	}

	return scale
}

func (t *Tonic) isAccidental() bool {
	splitTonic := strings.Split(t.note, "")
	if len(splitTonic) > 1 {
		if splitTonic[1] == "b" {
			t.flat = true
			return true
		} else {
			t.sharp = true
			return true
		}
	}
	return false
}
