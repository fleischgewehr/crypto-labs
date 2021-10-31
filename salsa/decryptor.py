import string
import typing as t
from collections import Counter


def xor(a: bytes, b: bytes) -> t.List[int]:
    return [i ^ j for i, j in zip(a, b)]


def find_key(ciphertexts: t.List[bytes]) -> t.List[int | None]:
    ciphertexts = sorted(ciphertexts, key=len)
    key = []
    max_length: int | None = None

    for ciphertext in ciphertexts:
        if max_length is not None:
            ciphertexts = [c[max_length:] for c in ciphertexts]
        max_length = len(ciphertext)

        key += _find_remaining(ciphertexts)

    return key


def _find_remaining(ciphertexts: t.List[bytes]) -> t.List[int | None]:
    shortest_cipher = min(len(text) for text in ciphertexts)
    key = [None for _ in range(shortest_cipher)]

    for idx, ciphertext in enumerate(ciphertexts):
        counter = Counter()

        for inner_idx, inner_ciphertext in enumerate(ciphertexts):
            if idx == inner_idx:
                continue
            counter.update(punctuation(xor(ciphertext, inner_ciphertext)))

        for idx, count in counter.items():
            if count == len(ciphertexts) - 1:
                key[idx] = ord(' ') ^ ciphertext[idx]

    return key


def punctuation(text: t.List[int]) -> Counter:
    c = Counter()
    for idx, char in enumerate(text):
        if char == 0x00 or chr(char) in string.ascii_letters:
            c[idx] += 1

    return c


def recover_text(ciphertexts: t.List[bytes], key: t.List[int | None]) -> t.List[str]:
    result = []
    for ciphertext in ciphertexts:
        l = [chr(i ^ j) if i is not None else '_' for i, j in zip(key, ciphertext)]
        result.append(''.join(l))

    return result


def attack(ciphertexts: t.List[bytes]) -> None:
    key = find_key(ciphertexts)
    # key_to_print = 
    text = recover_text(ciphertexts, key)
    with open('decrypted_ciphertexts.txt', 'w') as file:
        file.write('\n'.join(text))


if __name__ == '__main__':
    with open('ciphertexts.txt', 'r') as file:
        ciphertexts = [bytes.fromhex(line.rstrip()) for line in file]

    attack(ciphertexts)
