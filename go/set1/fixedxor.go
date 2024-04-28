package main

import (
  "encoding/hex"
  "fmt"
  "errors"
)

func main() {
 
  //Comentário futuro: o "_" usado como parâmetro duplo nas funções está sendo usado para ignorar o retorno de err que essas funções deveriam retornar. => Lembrar disso em outros projetos, conhecimento útil.

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

  //Checking if the inputs are equal
  if input2 != input2 { 
    return nil, errors.New("The inputs have different lengths, impossible to apply XOR.")
  }

  //Creating a buffer to hold the result value.
  result := make([]byte, len(input1))

  //Iterating trought the values of both inputs.
  for i := 0; i < len(input2); i++ {
    result[i] = input1[i] ^ input2[i]
  }
}
