package notes

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"os"
	"slices"
)

type Notes []Note

type NotesError struct {
	error
	ErrorValue        string
	InvalidIdError    string
	EmptyNotesError   string
	InvalidGroupError string
}

var notesError = NotesError{
	ErrorValue:        "",
	InvalidIdError:    "invalid ID",
	EmptyNotesError:   "you have no notes",
	InvalidGroupError: "group not fonde",
}

func (ne *NotesError) SetErrorValue(e string) {
	ne.ErrorValue = e
}

func (ne *NotesError) Error() string {
	return fmt.Sprintf("notes: %s", ne.ErrorValue)
}

// var nerr = NoteError{}
func (n *Notes) checkNotes() error {
	if n == nil || len(*n) == 0 {
		return &NotesError{ErrorValue: notesError.EmptyNotesError}
	}

	return nil
}

func (n *Notes) GetNote(id int) (*Note, bool) {
	for _, note := range *n {
		if note.Id == id {
			return &note, true
		}
	}
	return nil, false
}

func (n *Notes) AddNote(note Note) {
	note.Id = len(*n)
	*n = append(*n, note)
}

func (n *Notes) ResetId() {
	for i := range *n {
		(*n)[i].Id = i
	}
}

func (n *Notes) DeleteNote(id int) error {
	if note, ok := n.GetNote(id); ok {
		index := note.Id
		*n = slices.Delete(*n, index, index+1)
		n.ResetId()
		return nil
	}

	return &NotesError{ErrorValue: notesError.InvalidIdError}
}

func (n *Notes) DeleteGroup(group string) error {
	groupExist := false
	for i, note := range *n {
		if note.Group == group {
			*n = slices.Delete(*n, i, i+1)
			groupExist = true
		}
	}

	if !groupExist {
		return &NotesError{ErrorValue: notesError.InvalidGroupError}
	} else {
		return nil
	}
}

func (n *Notes) ShowAll() {
	err := n.checkNotes()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("All notes:")
	for _, note := range *n {
		note.Show()
	}
}

func (n *Notes) ShowNote(id int) {
	note, _ := n.GetNote(id)
	note.Show()
}

func (n *Notes) ShowGroup(group string, groups []string) error {
	if n == nil || len(*n) == 0 {
		fmt.Println("You have no notes!!")

		return &NotesError{ErrorValue: notesError.EmptyNotesError}
	}

	if !slices.Contains(groups, group) {
		fmt.Printf("Group \"%s\" not found\n", group)

		return &NotesError{ErrorValue: notesError.InvalidGroupError}
	}

	fmt.Println("Group: ", group)
	for _, note := range *n {
		if note.Group == group {
			note.Show()
		}
	}

	return nil
}

func (n *Notes) UpdateNote(id int, newNote string) error {
	for i := range *n {
		if (*n)[i].Id == id {
			(*n)[i].Note = newNote
			return nil
		}
	}

	return &NotesError{ErrorValue: notesError.InvalidIdError}
}

func (n *Notes) WriteToJson(jfile string) error {
	file, err := os.OpenFile(jfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	ml := jsontext.Multiline(true)
	wi := jsontext.WithIndent("    ")
	ne := jsontext.NewEncoder(file, ml, wi)

	ne.WriteToken(jsontext.BeginArray)

	for i, note := range *n {
		note.Id = i
		b, _ := json.Marshal(note)
		ne.WriteValue(b)
	}

	ne.WriteToken(jsontext.EndArray)

	return nil
}
