package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    // This map associates names with values
    messages := make(map[string]string)
    // Loop through the recieved slice of names. Call the Hello
    // function each time
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        // In the map associate the returned message with the name
        messages[name] = message
    }

    return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greetings messages at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi %v, Welcome!",
        "Great to see you, %v!",
        "Hail %v, Well met!",
    }

    // Return a random message by specifying a random index for the slice
    return formats[rand.Intn(len(formats))]
    
}

