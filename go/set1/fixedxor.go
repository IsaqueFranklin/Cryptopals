package main

import (
  "encoding/hex"
  "fmt"
)

func main() {
  
  input1, _ := decodeHex([]byte(""))

  input2, _ := decodeHex([]byte(""))

  decoded, _ := xorIt(input1, input2)
}

func decodeHex(input []byte) ([]byte, error) {
  
  db := make([]byte, hex.DecodedLen(len(input)))

  _, err := hex.Decode(db, input)
  if err != nil {
    return nil, err
  }

  return db, nil
}

func xorIt(input1, input2 []byte) ([]byte, error) {

}
