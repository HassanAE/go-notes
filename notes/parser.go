package notes

import (
	"fmt"
	"strconv"
	"strings"
)

type Commands struct {
	List,
	Add,
	Update,
	Delete,
	Group string
}

type SubCommands struct {
	Group string
	New   string
}

var Cmd = Commands{"list", "add", "update", "delete", "group"}
var SubCmd = SubCommands{"-group", "-new"}

func Parse(cmd string, notes *Notes, groups []string, filePath string) {
	buffer := strings.Split(cmd, " ")
	cmd = buffer[0]
	subCmd := ""
	index := 1
	group := "glob"

	if len(buffer) > 1 {
		if strings.HasPrefix(buffer[1], "-") {
			subCmd = buffer[1]
			index = 2
		}
	}

	line := strings.Join(buffer[index:], " ")
	//fmt.Println(line, subCmd)
command:
	switch cmd {
	case Cmd.List:
		switch {
		case subCmd == SubCmd.Group:
			notes.ShowGroup(line, groups)
		case line == "":
			fmt.Println("show all notes:")
			notes.ShowAll()
		default:
			fmt.Println("invalid command format")
		}

	case Cmd.Add:
		switch {
		case line == "":
			break command
		case subCmd == SubCmd.Group:
			buf := strings.Split(line, ":")
			group = buf[0]
			line = string(buf[1])
		}
		newNote := NewNote(line, group)
		notes.AddNote(*newNote)
		notes.WriteToJson(filePath)

	case Cmd.Delete:
		switch {
		case line == "":
			break command
		case subCmd == SubCmd.Group:
			notes.DeleteGroup(line)
			notes.WriteToJson(filePath)
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%s", err)
			break
		}

		err = notes.DeleteNote(id)
		if err != nil {
			fmt.Println(err)
			break
		}

		notes.WriteToJson(filePath)
	case Cmd.Update:
		temp := strings.Split(line, ":")
		id, _ := strconv.Atoi(temp[0])
		line = string(temp[1])
		notes.UpdateNote(id, line)

		notes.WriteToJson(filePath)

	default:
		fmt.Println("invaled command: ", cmd)
	}
}
