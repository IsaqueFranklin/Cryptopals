package main

import (
    "encoding/base64"
    "fmt"
)

func main() {

    data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

    sEnc := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("Encoded text: ",sEnc)

    sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec)) 
}
