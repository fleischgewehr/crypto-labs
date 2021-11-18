import typing as t

import requests

from src.cracker import Cracker
from src.misc import Gamemode


class TriToporaClient:
    id: int
    host: str
    balance: int = None

    def __init__(self, id: int, host: str) -> None:
        self.id = id
        self.host = host

    def create_account(self) -> t.Dict:
        url = f'http://{self.host}/casino/createacc'
        resp = requests.get(url, {'id': self.id}).json()
        if err := resp.get('error'):
            raise Exception(err)
        if balance := resp.get('money'):
            self.balance = balance

        return resp

    def play(self, mode: Gamemode, *, number: int = 1, bet: int = 1) -> t.Dict:
        url = f'http://{self.host}/casino/play{mode.value}'
        params = {'id': self.id, 'bet': bet, 'number': number}
        resp = requests.get(url, params).json()
        if err := resp.get('error'):
            raise Exception(err)
        if balance := resp.get('account', {}).get('money'):
            self.balance = balance

        return resp

    def rob_casino(self, cracker: Cracker, mode: Gamemode) -> int:
        try:
            for number in cracker:
                print(self.play(mode=mode, number=number, bet=100))
        except KeyboardInterrupt:
            return self.balance
