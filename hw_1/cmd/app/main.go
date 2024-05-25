package main

import (
	"flag"
	"fmt"
	"hw_1/internal/app"
	"log"
)

func main() {
	log.Println("starting..")

	filename := flag.String("filename", "problems.csv", "Имя загружаемого файла с тестом")
	mix := flag.Bool("mix", false, "Перемешать вопросы")

	flag.Parse()
	fmt.Printf("Upload file: %s", *filename)

	if err := run(*filename, *mix); err != nil {
		log.Fatal(err)
	}
}

func run(filename string, mix bool) error {
	app := app.NewApp(filename, mix)
	return app.Start()
}
