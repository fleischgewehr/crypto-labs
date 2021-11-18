import datetime as dt
import random

from dateutil import parser

from src.client import TriToporaClient
from src.misc import Gamemode
from src.mt_cracker import MtCracker


def rob_mt_casino(client: TriToporaClient) -> None:
    acc = client.create_account()
    deletion_date = parser.parse(acc['deletionTime'])
    # account has a lifespan of 3 hrs
    registration_date = deletion_date - dt.timedelta(hours=3)
    init_value = client.play(Gamemode.mt)['realNumber']
    cracker = MtCracker(registration_date, init_value)
    balance = client.rob_casino(cracker, Gamemode.mt)
    print(f'Balance after MT robbery: {balance}')


if __name__ == '__main__':
    client = TriToporaClient(id=random.randint(1, 100_000), host='95.217.177.249')
    rob_mt_casino(client)
