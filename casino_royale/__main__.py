if __name__ == '__main__':
    import random

    from src.client import TriToporaClient


    client = TriToporaClient(id=random.randint(1, 100_000), host='94.217.177.249')
    client.create_account()
