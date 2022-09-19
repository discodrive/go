package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined logger
    // including the log entry prefix and a flag
    // to disable printing the time, source file and 
    // line number. 
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Dave", "Gladys", "Stephen"}

    // Request greetings for each of the names
    messages, err := greetings.Hellos(names)

    // If an error was returned, print it to the console and
    // exit the program
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}
