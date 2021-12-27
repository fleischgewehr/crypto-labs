import random
import secrets
from functools import lru_cache
from string import punctuation

from src import constants as const


@lru_cache
def _get_common_words():
    with open('pwd_generator/words.txt', 'r') as file:
        return file.readlines()


@lru_cache
def _get_frequent_passwords():
    with open('pwd_generator/frequent.txt', 'r') as file:
        return file.readlines()


@lru_cache
def _get_common_passwords():
    with open('pwd_generator/common.txt', 'r') as file:
        return file.readlines()


def get_frequent_password():
    return random.choice(_get_frequent_passwords()).strip()


def get_common_password():
    return random.choice(_get_common_passwords()).strip()


def get_random_password():
    return secrets.token_urlsafe(10)


def get_combined_password():
    words = _get_common_words()

    w1 = random.choice(words).strip()
    w2 = random.choice(words).strip().capitalize()
    
    return f'{w1}{w2}{random.choice(punctuation)}{random.randrange(0, 2000)}'


def generate_passwords(n):
    passwords = []
    for _ in range(int(const.N * const.COMMON)):
        passwords.append(get_common_password())
    for _ in range(int(const.N * const.FREQUENT)):
        passwords.append(get_frequent_password())
    for _ in range(int(const.N * const.RANDOM)):
        passwords.append(get_random_password())
    for _ in range(int(const.N * const.COMBINED)):
        passwords.append(get_combined_password())

    return passwords
