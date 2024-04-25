package main

import (
    "encoding/base64"
    "fmt"
)

func main() {

    data := "abc123!?$*&()'-=@~"

    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("Encoded text: ",sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec)) 
}
