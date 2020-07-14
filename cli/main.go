package main

import (
	"fmt"
	"os"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("app", "1.0.0")    //这里指定了APP名字和版本
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"foo": fooCommandFactory,    //定义foo命令和工厂
		"bar": barCommandFactory,    //定义bar命令和工厂
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}

//foo命令工厂
func fooCommandFactory() (cli.Command, error) {
	fmt.Println("I am foo command")
	return new(FooCommand), nil
}
//bar命令工厂
func barCommandFactory() (cli.Command, error) {
	fmt.Println("I am bar command")
	return new(BarCommand), nil
}
//foo命令实现
type FooCommand struct{}

// Help should return long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (f *FooCommand) Help() string {
	return "help foo"
}

// Run should run the actual command with the given CLI instance and
// command-line arguments. It should return the exit status when it is
// finished.
//
// There are a handful of special exit codes this can return documented
// above that change behavior.
func (f *FooCommand) Run(args []string) int {
	fmt.Println("Foo Command is running")
	return 1
}

// Synopsis should return a one-line, short synopsis of the command.
// This should be less than 50 characters ideally.
func (f *FooCommand) Synopsis() string {
	return "foo command Synopsis"
}

//bar命令实现
type BarCommand struct{}

// Help should return long-form help text that includes the command-line
// usage, a brief few sentences explaining the function of the command,
// and the complete list of flags the command accepts.
func (b *BarCommand) Help() string {
	return "help bar"
}

// Run should run the actual command with the given CLI instance and
// command-line arguments. It should return the exit status when it is
// finished.
//
// There are a handful of special exit codes this can return documented
// above that change behavior.
func (b *BarCommand) Run(args []string) int {
	fmt.Println("bar Command is running")
	return 1
}

// Synopsis should return a one-line, short synopsis of the command.
// This should be less than 50 characters ideally.
func (b *BarCommand) Synopsis() string {
	return "bar command Synopsis"
}