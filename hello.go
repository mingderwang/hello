package main

import (
  "log"

  "github.com/codegangsta/martini"
  "github.com/mipearson/logstasher"
  "github.com/mipearson/rfw"
)

func main() {
  m := martini.Classic()

  logstashLogFile, err := rfw.Open("hello.log", 0644)
  if err != nil {
    log.Fatalln(err)
  }
  defer logstashLogFile.Close()
  m.Use(logstasher.Logger(logstashLogFile))

  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}

