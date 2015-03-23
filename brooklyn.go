package main

import (
	"os"
	//"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/app"
	"github.com/robertgmoss/brooklyn-cli/command_factory"
	"github.com/robertgmoss/brooklyn-cli/command_runner"
	"github.com/robertgmoss/brooklyn-cli/net"
)

func main() {
	network := net.NewNetwork("http://192.168.50.101:8081", "brooklyn", "Sns4Hh9j7l")
	cmdFactory := command_factory.NewFactory(network)
	cmdRunner := command_runner.NewRunner(cmdFactory)
	metaDatas := cmdFactory.CommandMetadatas()
	theApp := app.NewApp(cmdRunner, metaDatas...)
	theApp.Run(os.Args)
}
