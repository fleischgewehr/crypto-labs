
import typing as typ

ASCII_RANGE = list(range(97, 122)) + [32]


def bxor(a: typ.List[int], b: typ.List[int]) -> bytes:
    return bytes([x ^ y for x, y in zip(a, b)])
