import hashlib

import bcrypt


def get_md5_hash(plain):
    return hashlib.md5(plain.encode('utf-8')).hexdigest()


def get_bcrypt_hash(plain):
    salt = bcrypt.gensalt()
    hash = bcrypt.hashpw(plain.encode('utf-8'), salt)
    
    return hash.decode('utf-8'), salt.decode('utf-8')
