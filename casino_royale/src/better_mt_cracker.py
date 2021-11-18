import typing as typ

from src.cracker import Cracker
from src.mt import MT19937


class BetterMt(MT19937):
    def __init__(self, state: typ.List[int]) -> None:
        self.index = len(state) + 1
        self.state = state


class BetterMtCracker(Cracker):
    state: typ.List[int]
    generator: BetterMt

    def __init__(self, state: typ.List[int]) -> None:
        assert len(state) == 624, 'Initial state must contain 624 records'
        self.state = state

    def crack(self) -> None:
        self.generator = BetterMt(self.state)

    def __iter__(self) -> typ.Iterator[int]:
        self.crack()
        for value in self.generator:
            yield value
