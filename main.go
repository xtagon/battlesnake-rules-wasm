package main

import (
  "encoding/json"
  "syscall/js"
  "github.com/BattlesnakeOfficial/rules"
)

func createInitialBoardState(this js.Value, args []js.Value) interface{} {
  ruleset := rules.StandardRuleset{}

  width := args[0].Int()
  height := args[1].Int()
  snakeIDsJson := args[2].String()

  var snakeIDs []string

  err := json.Unmarshal([]byte(snakeIDsJson), &snakeIDs)

  if err != nil {
    return err.Error()
  }

  initialState, err := ruleset.CreateInitialBoardState(int32(width), int32(height), snakeIDs)

  if err != nil {
    return err.Error()
  }

  initialStateJson, err := json.Marshal(initialState)

  if err != nil {
    return err.Error()
  }

  return string(initialStateJson)
}

func createNextBoardState(this js.Value, args[]js.Value) interface{} {
  ruleset := rules.StandardRuleset{}

  prevStateJson := args[0].String()
  movesJson := args[1].String()

  var prevState rules.BoardState
  var moves []rules.SnakeMove

  err := json.Unmarshal([]byte(prevStateJson), &prevState)

  if err != nil {
    return err.Error()
  }

  err = json.Unmarshal([]byte(movesJson), &moves)

  if err != nil {
    return err.Error()
  }

  nextState, err := ruleset.CreateNextBoardState(&prevState, moves)

  if err != nil {
    return err.Error()
  }

  nextStateJson, err := json.Marshal(nextState)

  if err != nil {
    return err.Error()
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
