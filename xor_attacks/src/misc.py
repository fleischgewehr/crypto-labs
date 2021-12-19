
import typing as typ

ASCII_RANGE = list(range(97, 122)) + [32]


def bxor(a: typ.List[int], b: typ.List[int]) -> bytes:
    return bytes([x ^ y for x, y in zip(a, b)])


def find_hamming_distance(a: int, b: int) -> int:
    return sum(bin(byte).count('1') for byte in bxor(a, b))


def score_vigenere_key_size(candidate_key_size: int, ciphertext: str) -> int:
    slice_size = 2 * candidate_key_size

    # calculate number of samples from given ciphertext
    nb_measurements = len(ciphertext) // slice_size - 1

    # lowest score == more likely
    score = 0
    for i in range(nb_measurements):

        s = slice_size
        k = candidate_key_size
        slice_1 = slice(i * s, i * s + k)
        slice_2 = slice(i * s + k, i * s + k * 2)

        score += find_hamming_distance(ciphertext[slice_1], ciphertext[slice_2])

    # normalize score to avoid biasing towards long key sizes
    score /= candidate_key_size
    score /= nb_measurements

    return score


def find_vigenere_key_length(
    ciphertext: str, min_length: int = 2, max_length: int = 30
) -> int:
    key = lambda x: score_vigenere_key_size(x,ciphertext)
    return min(range(min_length, max_length), key=key)
