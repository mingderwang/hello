package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()

	logstashLogFile, err := rfw.Open("hello.log", 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer log.stashLogFile.Close()
	m.Use(logstasher.Logger(logstashLogFile))
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
