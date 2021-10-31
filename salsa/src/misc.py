import string
import typing as t
from collections import Counter


def xor(a: bytes, b: bytes) -> t.List[int]:
    return [i ^ j for i, j in zip(a, b)]


def find_key(ciphertexts: t.List[bytes]) -> t.List[int | None]:
    """
    Find key by using XOR attack where k is the reused pad:
    xor(p1, k) = c1
    xor(p2, k) = c2
    considering xor(a, a) = 0 & xor(a, xor(b, c)) = xor(xor(a, b), c)
    then:
    xor(c1, k) = xor(xor(p1, k), k), where c1 = xor(p1, k)
    then:
    xor(p1, xor(k, k)) => xor(p1, 0) = p1

    also:
    xor(c1, c2) == xor(xor(p1, k), xor(p2, k)) == xor(xor(p1, p2), xor(k, k)) =>
    => xor(c1, c2) == xor(p1, p2)
    """
    # sort ciphertexts from shorter to longer
    ciphertexts = sorted(ciphertexts, key=len)
    key = []
    max_length: int | None = None

    for ciphertext in ciphertexts:
        if max_length is not None:
            # cut off already attacked chunk of ciphertexts
            ciphertexts = [c[max_length:] for c in ciphertexts]
        max_length = len(ciphertext)

        key += _find_remaining(ciphertexts)

    return key


def _find_remaining(ciphertexts: t.List[bytes]) -> t.List[int | None]:
    shortest_cipher = min(len(text) for text in ciphertexts)
    # fill list with "N" None values, where N is the shortest ciphertext length
    # because we want to XOR only equal parts of ciphertexts
    key = [None for _ in range(shortest_cipher)]

    for idx, ciphertext in enumerate(ciphertexts):
        counter = Counter()

        for inner_idx, inner_ciphertext in enumerate(ciphertexts):
            # no need to xor equal ciphertexts
            if idx == inner_idx:
                continue
            # save found punctuation
            counter.update(punctuation(xor(ciphertext, inner_ciphertext)))

        for idx, count in counter.items():
            # if we met space in every inner ciphertext then it's present
            # at given index at the outer ciphertext
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
        # recover text by XORing key and given ciphertext and fill it with
        # underscore if the key was not found at given position
        l = [chr(i ^ j) if i is not None else '_' for i, j in zip(key, ciphertext)]
        result.append(''.join(l))

    return result
