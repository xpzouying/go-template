package cli

import (
	"fmt"

	"github.com/xpzouying/go-cmd-project-template/internal/constant"
	"github.com/xpzouying/go-cmd-project-template/log"

	cli "github.com/urfave/cli/v2"
)

var cliLogger = log.MustNewLogger("info").Sugar().Named("cli-app")

func startAction(c *cli.Context) error {
	_, err := InitConfigAndComponents()
	if err != nil {
		cliLogger.Fatalf("Failed to initialize config and components: %v", err)
	}

	cliLogger.Info("Starting...")
	return nil
}

func CreateCliApp() *cli.App {

	cli.VersionPrinter = func(c *cli.Context) {
		println("Welcome!")
		fmt.Printf("Version: %s\n", constant.Version)
		fmt.Printf("Build Time: %s\n", constant.BuildTime)
		fmt.Printf("Git Commit: %s\n", constant.GitRevision)
	}

	app := cli.NewApp()
	app.Name = "go-template-project"
	app.Flags = RootFlags
	app.Version = constant.Version
	app.Usage = "go-template-project is a template project for go command line application."
	app.Action = startAction

	return app
}
