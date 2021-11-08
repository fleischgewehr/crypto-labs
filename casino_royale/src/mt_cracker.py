import datetime as dt
import functools
import time
import typing as t

from src.cracker import Cracker
from src.mt import MT19937


class MtCracker(Cracker):
    registration_date: dt.datetime
    initial_value: int
    generator: MT19937

    def __init__(self, registration_date: dt.datetime, initial_value: int) -> None:
        self.registration_date = registration_date
        self.initial_value = initial_value

    @staticmethod
    def find_seed(seed: int, initial_value: int) -> bool:
        generator = iter(MT19937(seed))

        return next(generator) == initial_value

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

        generator = iter(MT19937(seed))
        # skip already checked initial value
        next(generator)
        self.generator = generator

    def __iter__(self) -> t.Iterator[int]:
        self.crack()
        for idx, value in enumerate(self.generator):
            if idx == 10:
                break
            # TODO: to be tested
            yield value
