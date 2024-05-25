package app

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type Application struct {
	Filename string
	Mix      bool
	Data     [][]string

	User User // TODO
}

type User struct {
	Name  string
	Right int
	Wrong int
	Stats float32
}

func NewApp(filename string, mix bool) Application {
	return Application{Filename: filename, Mix: mix}
}

func (a *Application) Start() error {
	log.Println("open file:", a.Filename)

	file, err := os.Open(a.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// asking que
	var answer string
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	a.Data, err = reader.ReadAll()
	if err != nil {
		log.Println("error reading some string in csv-file")
	}

	if a.Mix != false {
		a.ShuffleData()
	}

	for i := 0; i < len(a.Data); i++ {

		if len(a.Data[i][0:len(a.Data[i])-1]) >= 2 {
			fmt.Printf("Вопрос %d: %s | ", i+1, strings.Join(a.Data[i][0:len(a.Data[i])-1], ","))
		} else {
			fmt.Printf("Вопрос %d: %s | ", i+1, strings.Join(a.Data[i][0:len(a.Data[i])-1], " "))
		}

		answer, err = bufio.NewReader(os.Stdin).ReadString('\n')
		answer = strings.TrimSpace(strings.ReplaceAll(answer, " ", ""))

		if strings.ToLower(answer) == strings.ToLower(strings.ReplaceAll(a.Data[i][len(a.Data[i])-1], " ", "")) {
			a.User.Right++
			fmt.Println("Верно!")
		}
	}
	a.User.Wrong = len(a.Data) - a.User.Right
	fmt.Printf("Результаты теста: \n Правильных: %d \n Неправильных: %d", a.User.Right, a.User.Wrong)
	return nil
}

func (a *Application) ShuffleData() {
	rand.Shuffle(len(a.Data),
		func(i, j int) { a.Data[i], a.Data[j] = a.Data[j], a.Data[i] })
}
