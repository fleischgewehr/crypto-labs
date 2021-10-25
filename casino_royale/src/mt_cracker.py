import datetime as dt
import functools
import time
import typing as t

from numpy.random import Generator, MT19937

from src.cracker import Cracker


class MtCracker(Cracker):
    registration_date: dt.datetime
    initial_value: int
    generator: Generator

    def __init__(self, registration_date: dt.datetime, initial_value: int) -> None:
        self.registration_date = registration_date
        self.initial_value = initial_value

    @staticmethod
    def find_seed(seed: int, initial_value: int) -> bool:
        generator = Generator(MT19937(seed))

        return generator.random_raw() == initial_value

    @staticmethod
    def search(
        from_: int, to: int, oracle: t.Callable[[int], bool]
    ) -> t.Optional[int]:
        assert from_ <= to, '"from" value is GT "to" value'
        for candidate in range(from_, to + 1):
            if oracle(candidate):
                return candidate

    def crack(self) -> None:
        delta = 16_000
        from_ = int(time.mktime(self.registration_date.timetuple()))
        to = from_ + delta * 2
        oracle = functools.partial(self.find_seed, initial_value=self.initial_value)

        seed = self.search(from_, to, oracle)
        assert seed is not None, 'Unable to find seed'

        generator = Generator(MT19937(seed))
        # skip already checked initial value
        generator.random_raw()
        self.generator = generator

    def __iter__(self) -> t.Iterator[int]:
        self.crack()
        for _ in range(10):
            # TODO: to be tested
            yield self.generator.random_raw()
