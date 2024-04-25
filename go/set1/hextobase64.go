package main

import (
    "encoding/base64"
    "encoding/hex"
    "fmt"
)

func main() {

    data := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

    db, err := decodeHex(data)
    if err != nil {
        fmt.Printf("Failed to decode hex: %s", err)
        return
    }

    f := base64Encode(db)
    fmt.Printf("%s", f)
    /*sEnc := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("Encoded text: ",sEnc)

    sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))*/ 
}

func decodeHex(input []byte) ([]byte, error) {
    db := make([]byte, hex.DecodedLen(len(input)))

    _, err := hex.Decode(db, input)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func base64Encode(input []byte) ([]byte) {
    eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
    base64.StdEncoding.Encode(eb, input)

    return eb
}
