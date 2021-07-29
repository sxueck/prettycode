package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		dst     string
		exclude string
	)

	app := &cli.App{
		Name:  "prettycode",
		Usage: "automatically format the code snippets in the md document for you",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "dst",
				Value:       ".",
				Usage:       "format the corresponding directory",
				Destination: &dst,
			},
			&cli.StringFlag{
				Name:        "exclude",
				Value:       "",
				Usage:       "exclude certain directories",
				Destination: &exclude,
			},
		},
		Action: func(c *cli.Context) error {
			err := filepath.Walk(dst, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
					return err
				}
				if info.IsDir() {
					//fmt.Printf("dir : %+v \n", info.Name())
					return nil
				}

				if match, _ := filepath.Match("*.go", info.Name()); match {
					fmt.Printf("found : %s\n", path)
				}
				return nil
			})

			if err != nil {
				return err
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
