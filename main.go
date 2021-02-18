package main

import (
  "encoding/json"
  "fmt"
  "syscall/js"
  "github.com/BattlesnakeOfficial/rules"
)

type optsForInit struct {
  Height int32
  Width int32
  SnakeIDs []string
}

type optsForNext struct {
  PreviousState rules.BoardState
  SnakeMoves []rules.SnakeMove
}

func createInitialBoardState(this js.Value, args []js.Value) interface{} {
  if len(args) != 1 || args[0].Type() != js.TypeString {
    fmt.Println("unexpected arguments")
    return js.Null()
  }

  var opts optsForInit

  err := json.Unmarshal([]byte(args[0].String()), &opts)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  ruleset := rules.StandardRuleset{}

  initialState, err := ruleset.CreateInitialBoardState(opts.Width, opts.Height, opts.SnakeIDs)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  initialStateJson, err := json.Marshal(initialState)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  return string(initialStateJson)
}

func createNextBoardState(this js.Value, args[]js.Value) interface{} {
  if len(args) != 1 || args[0].Type() != js.TypeString {
    fmt.Println("unexpected arguments")
    return js.Null()
  }

  var opts optsForNext

  err := json.Unmarshal([]byte(args[0].String()), &opts)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  ruleset := rules.StandardRuleset{}

  nextState, err := ruleset.CreateNextBoardState(&opts.PreviousState, opts.SnakeMoves)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  nextStateJson, err := json.Marshal(nextState)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  return string(nextStateJson)
}

func makeExports() map[string]interface{} {
  exports := make(map[string]interface{})

  exports["createInitialBoardState"] = js.FuncOf(createInitialBoardState)
  exports["createNextBoardState"] = js.FuncOf(createNextBoardState)

  return exports
}

func main() {
  js.Global().Set("BattlesnakeRules", makeExports())

  // The main function must continue running, or else the exported functions
  // will no longer be available to use.
  <-make(chan bool)
}
