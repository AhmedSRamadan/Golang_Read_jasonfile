package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Elements struct {
    Elements []Element `json:"elements"`
}


type Element struct {
    User   string `json:"user"`
    Color   string `json:"color"`
}

func main() {
    arg := os.Args
    jsonFile, err := os.Open("file.json")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    var elements Elements
    json.Unmarshal(byteValue, &elements)
    for i := 0; i < len(elements.Elements); i++ {
        if (elements.Elements[i].Color) == arg[1] {
          fmt.Println("User Name: " + elements.Elements[i].User)
          fmt.Println("User Color: " + elements.Elements[i].Color)
        }
    }
}