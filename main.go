package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "math/rand"
  "syscall/js"
  "github.com/BattlesnakeOfficial/rules"
)

type optsForInit struct {
  Seed int64
  RulesetName string
  RulesetParams rulesetParams
  Height int32
  Width int32
  SnakeIDs []string
}

type optsForNext struct {
  Seed int64
  RulesetName string
  RulesetParams rulesetParams
  PreviousState rules.BoardState
  SnakeMoves []rules.SnakeMove
}

type optsForGameOver struct {
  RulesetName string
  RulesetParams rulesetParams
  BoardState rules.BoardState
}

// This struct holds params available to ANY supported ruleset. This is an
// anti-pattern and is really just a lazy way to avoid figuring out how to
// appease Go's lack of union types. When a ruleset is intantiated in
// makeRuleset, only the params that it supports will be used.
type rulesetParams struct {
  // Applies to Standard, Solo, and Constrictor:
  FoodSpawnChance int32
  MinimumFood int32

  // Royale and Squad params not yet supported
}

func makeRuleset(name string, params rulesetParams) (rules.Ruleset, error) {
  standard := rules.StandardRuleset{
    FoodSpawnChance: params.FoodSpawnChance,
    MinimumFood: params.MinimumFood,
  }

  switch name {
  case "standard":
    return &standard, nil
  case "solo":
    return &rules.SoloRuleset{StandardRuleset: standard}, nil
  case "constrictor":
    return &rules.ConstrictorRuleset{StandardRuleset: standard}, nil
  }

  return nil, errors.New("unsupported ruleset")
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

  ruleset, err := makeRuleset(opts.RulesetName, opts.RulesetParams)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  rand.Seed(opts.Seed);

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

  ruleset, err := makeRuleset(opts.RulesetName, opts.RulesetParams)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  rand.Seed(opts.Seed);

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

func isGameOver(this js.Value, args[]js.Value) interface{} {
  if len(args) != 1 || args[0].Type() != js.TypeString {
    fmt.Println("unexpected arguments")
    return js.Null()
  }

  var opts optsForGameOver

  err := json.Unmarshal([]byte(args[0].String()), &opts)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  ruleset, err := makeRuleset(opts.RulesetName, opts.RulesetParams)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  gameOverBool, err := ruleset.IsGameOver(&opts.BoardState)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  gameOverJson, err := json.Marshal(gameOverBool)

  if err != nil {
    fmt.Println(err.Error())
    return js.Null()
  }

  return string(gameOverJson)
}

func makeExports() map[string]interface{} {
  exports := make(map[string]interface{})

  exports["createInitialBoardState"] = js.FuncOf(createInitialBoardState)
  exports["createNextBoardState"] = js.FuncOf(createNextBoardState)
  exports["isGameOver"] = js.FuncOf(isGameOver)

  return exports
}

func main() {
  js.Global().Set("BattlesnakeRules", makeExports())

  // The main function must continue running, or else the exported functions
  // will no longer be available to use.
  <-make(chan bool)
}
