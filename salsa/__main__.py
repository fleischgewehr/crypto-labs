from src.misc import find_key, recover_text


if __name__ == '__main__':
    with open('salsa/ciphertexts.txt', 'r') as file:
        ciphertexts = [bytes.fromhex(line.rstrip()) for line in file]

    key = find_key(ciphertexts)
    text = recover_text(ciphertexts, key)

    with open('salsa/decrypted_ciphertexts.txt', 'w') as file:
        file.write('\n'.join(text))
