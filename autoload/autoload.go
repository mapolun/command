package autoload

import (
	"command/cmd"
	"github.com/ctfang/command"
)

func New(console command.Console) {
	console.AddCommand(cmd.Foo{})
}
