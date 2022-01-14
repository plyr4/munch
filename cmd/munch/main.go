package main

import (
    "github.com/oakmound/oak/v3"
    "github.com/oakmound/oak/v3/scene"
)

func main() {
    oak.AddScene("firstScene", scene.Scene{
        Start: func(*scene.Context) {
            // ... draw entities, bind callbacks ... 
        }, 
    })
    oak.Init("firstScene")
}