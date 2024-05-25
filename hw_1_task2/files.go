package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	// "github.com/urfave/cli/v2"
)

func DeleteFile(c *cli.Context) error {
	fmt.Println("deleting..")
	err := os.Remove(c.String("filename"))
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(c *cli.Context) error {
	fmt.Println("creating..")
	_, err := os.Create(c.String("filename"))
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(c *cli.Context) error {
	fmt.Println("reading..")
	file, err := os.Open(c.String("filename"))
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}
	fmt.Println("Stats:")
	fmt.Println("Name: ", fi.Name())
	fmt.Println("Size:", fi.Size())
	fmt.Println("IsDir:", fi.IsDir())
	fmt.Println("Mode:", fi.Mode())

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "files"
	app.Usage = "methods with files"

	app.Commands = []cli.Command{
		{
			Name:        "delete",
			HelpName:    "delete",
			Action:      DeleteFile,
			ArgsUsage:   "",
			Usage:       "Delete selected file",
			Description: "delete file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "filename",
					Usage: "filename to delete",
				},
			},
		},

		{
			Name:        "create",
			HelpName:    "create",
			Action:      CreateFile,
			ArgsUsage:   "",
			Usage:       "create file with name",
			Description: "create file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "filename",
					Usage: "filename to create",
				},
			},
		},

		{
			Name:        "read",
			HelpName:    "read",
			Action:      ReadFile,
			ArgsUsage:   "",
			Usage:       "read selected file",
			Description: "read file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "filename",
					Usage: "filename to read",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
