import base64
import binascii

from src.attacks import attack_repeating_key_xor, attack_single_byte_xor


if __name__ == '__main__':
    with open('xor_attacks/src/single-byte-xor.txt', 'r') as f:
        raw_cipher = f.read()
    cipher = binascii.unhexlify(raw_cipher)
    res1 = attack_single_byte_xor(cipher)

    with open('xor_attacks/src/repeating-key-xor.txt', 'r') as f:
        raw_cipher = f.read()
    cipher = base64.b64decode(raw_cipher)
    res2 = attack_repeating_key_xor(cipher)

    print('#' * 10, ' Decoded single-byte XOR cipher: ', '#' * 10)
    print(res1)
    print('#' * 10, ' Decoded repeating-key XOR cipher: ', '#' * 10)
    print(res2)
