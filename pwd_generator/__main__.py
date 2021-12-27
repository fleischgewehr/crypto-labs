import csv

from src import constants as const
from src import crypto
from src import misc


if __name__ == '__main__':
    passwords = misc.generate_passwords(const.N)
    print('Passwords were generated')
    with open('md5.csv', 'w') as md5_file, open('bcrypt.csv', 'w') as bcrypt_file:
        md5_writer = csv.writer(md5_file)
        md5_writer.writerow(['Hash'])
        bcrypt_writer = csv.writer(bcrypt_file)
        bcrypt_writer.writerow(['Hash', 'Salt'])
        for password in passwords:
            md5_writer.writerow([crypto.get_md5_hash(password)])
            bcrypt_writer.writerow([*crypto.get_bcrypt_hash(password)])
    print('Hashes were generated')

    while True:
        match input():
            case "common":
                print(misc.get_common_password())
            case "frequent":
                print(misc.get_frequent_password())
            case "random":
                print(misc.get_random_password())
            case "combined":
                print(misc.get_combined_password())
            case "exit":
                break
            case _:
                print("invalid command")
