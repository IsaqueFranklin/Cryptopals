def main():
    input1 = decode_hex("1c0111001f010100061a024b53535009181c")
    input2 = decode_hex("686974207468652062756c6c277320657965")

    decoded = xor_it(input1, input2)
    print(encode_hex(decoded))

def decode_hex(hex_string):
    return bytes.fromhex(hex_string)

def encode_hex(input_bytes):
    return input_bytes.hex()

def xor_it(input1, input2):
    # Checking if the inputs are equal
    if len(input1) != len(input2):
        raise ValueError("The inputs have different lengths, impossible to apply XOR.")

    # Creating a buffer to hold the result value
    result = bytearray(len(input1))

    # Iterating through the values of both inputs
    for i in range(len(input1)):
        result[i] = input1[i] ^ input2[i]

    return result

if __name__ == "__main__":
    main()
