package notes

import (
	"fmt"
	"strings"
	"time"
)

type Note struct {
	Id    int
	Note  string
	Group string
	Time  string
}

func NewNote(note, group string) *Note {
	t := time.Now().Format(time.DateTime)
	return &Note{
		Id:    0,
		Note:  note,
		Group: group,
		Time:  t,
	}
}

func (n *Note) Show() {
	line := strings.Split(n.Note, " ")
	if len(line) > 5 {
		firstLine := strings.Join(line[:7], " ")
		secondLine := strings.Join(line[7:], " ")
		fmt.Printf(
			"%d) %s\n   %-35s at:%s \tgroup( %s )\n",
			n.Id, firstLine, secondLine, n.Time, n.Group)
		fmt.Println()
	}
	fmt.Printf(
		"%d) %-35s at:%s \tgroup( %s )\n",
		n.Id, strings.Join(line, " "), n.Time, n.Group)
	fmt.Println()
}
