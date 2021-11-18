import typing as typ

from src.cracker import Cracker


class LcgCracker(Cracker):
    """next = (a * current + c) % m"""
    m: int = 2 ** 32
    a: typ.Optional[int]
    c: typ.Optional[int]
    results: typ.List[int]

    def __init__(
        self, results: typ.List[int], m: int = 2 ** 32, a: int = None, c: int = None
    ) -> None:
        assert len(results) > 2, 'Initial results must contain more than 2 records'
        self.m = m
        self.a = a
        self.c = c
        self.results = results[-3:]

    @property
    def current(self) -> int:
        return self.results[-1]

    def _egcd(self, a: int, b: int) -> typ.Tuple[int, int, int]:
        if a == 0:
            return (b, 0, 1)
        else:
            g, x, y = self._egcd(b % a, a)
            return (g, y - (b // a) * x, x)

    def _modinv(self, b: int, n: int) -> typ.Optional[int]:
        g, x, _ = self._egcd(b, n)
        if g == 1:
            return x % n

    def crack(self) -> None:
        modinv = self._modinv(self.results[1] - self.results[0], self.m)
        self.a = (self.current - self.results[1]) * modinv % self.m
        self.c = (self.results[1] - self.results[0] * self.a) % self.m

    def __iter__(self) -> int:
        self.crack()
        while True:
            next_ = (self.a * self.current + self.c) % self.m
            self.results.append(next_)

            yield next_
