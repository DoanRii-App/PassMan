package main

import (
  "log"

  "github.com/sethvargo/go-password/password"
)

func GenPass(l int, d int, sym int, ulC bool, drC bool) {
  // Generate a password that is 64 characters long with 10 digits, 10 symbols,
  // allowing upper and lower case letters, disallowing repeat characters.
  res, err := password.Generate(l, d, sym, ulC, drC)
  if err != nil {
    log.Fatal(err)
  }
  log.Printf(res)
}
