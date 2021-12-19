import binascii

from src.attacks import attack_single_byte_xor


if __name__ == '__main__':
    with open('xor_attacks/src/single-byte-xor.txt', 'r') as f:
        raw_cipher = f.read()
    cipher = binascii.unhexlify(raw_cipher)
    res = attack_single_byte_xor(cipher)

    print('#' * 10, ' Decoded single-byte XOR cipher: ', '#' * 10)
    print(res)
