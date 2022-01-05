# Notes about `if` statements efficiency

#### Three different `if` statements

Program will check every single statement even if the first statement is true

```go
if playerChoice == "rock" {
  playerValue = ROCK
}

if playerChoice == "paper" {
  playerValue = PAPER
}

if playerChoice == "scissors" {
  playerValue = SCISSORS
}
```

#### Three statements using `if` and `else if` and `else`

If the first statement is true, program will stop checking the rest.

```go
if playerChoice == "rock" {
  playerValue = ROCK
}

else if playerChoice == "paper" {
  playerValue = PAPER
}

else if playerChoice == "scissors" {
  playerValue = SCISSORS
}
```

It is clearly more efficient to use `if esle` statements
