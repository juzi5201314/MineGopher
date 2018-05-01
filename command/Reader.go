package command

import (
	"io"
	"bufio"
	"strings"

	"github.com/juzi5201314/minegopher"
	"github.com/juzi5201314/minegopher/utils"
)

func NewCommandReader(inputReader io.Reader) *CommandReader {
	reader := &CommandReader{bufio.NewReader(inputReader), []func(string){}}
	go func() {
		for {
			reader.readLine()
		}
	}()
	return reader
}

type CommandReader struct {
	reader *bufio.Reader
}

func (reader *CommandReader) readLine() {
	command, _ := reader.reader.ReadString(0x0a)
	command = strings.Trim(command, "\n")
	manager := minegopher.GetServer().GetCommandManager()
	args := strings.Split(command, " ")
	commandName := args[0]
	i := 1
	if !manager.IsCommandRegistered(commandName) {
		utils.GetLogger().Error("Command could not be found.")
		return
	}

	for {
		if i == len(args) {
			break
		}
		commandName += " " + args[i]
		i++
	}
	args = args[i:]
	command, _ := manager.GetCommand(commandName)
	command.Execute(minegopher.GetServer(), args)
}