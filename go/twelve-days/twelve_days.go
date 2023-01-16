package twelve

import (
	"fmt"
	"strings"
)

const DAYS_OF_CHRISTMAS = 12

func Song() string {

	var song strings.Builder

	for i := 1; i <= DAYS_OF_CHRISTMAS; i++ {
		song.WriteString(Verse(i)+"\n")
	}
	return song.String()
}

func Verse(day int) string {
	gifts_list := []string{
		"a Partridge in a Pear Tree.",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming"}
	stanza := "On the %s day of Christmas my true love gave to me: %s"
	ordinal_days := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth"}
	var gifts strings.Builder
	var sb strings.Builder

	if day == 1 {
		fmt.Fprintf(&sb, stanza, "first", gifts_list[0])
	} else {
		for x := len(gifts_list[0:day]) - 1; x > 0; x-- {
			gifts.WriteString(gifts_list[x] + ", ")
		}
		gifts.WriteString("and "+gifts_list[0])
		fmt.Fprintf(&sb, stanza, ordinal_days[day-1], gifts.String())
		gifts.Reset()
	}

	return sb.String()
}