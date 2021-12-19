import base64
import typing as typ

from src.misc import ASCII_RANGE, bxor, find_vigenere_key_length


"""
Well, and now you see that it’s absolutely pointless to try to invent your own “crypto” with expectations that if 
someone doesn’t know the algorithm (even binary base64, which is pretty stupid) it’s impossible for him to decipher. 
You’re reading this now in plain text so that argument clearly falls short.

Now to the actual tasks. All of them are graded based on the code you write, so no point in stealing deciphered text from your classmates.
1. Write a piece of software to attack a single-byte XOR cipher which is the same as Caesar but with xor op.
"""
def attack_single_byte_xor(ciphertext: str) -> typ.Tuple[bytes, bytes]:
    best_candidate = None
    max_letters_count = 0
    key = None
    for i in range(256):
        candidate_key = i.to_bytes(1, byteorder='big')
        keystream = candidate_key * len(ciphertext)
        candidate_message = bxor(ciphertext, keystream)
        letters_count = sum([x in ASCII_RANGE for x in candidate_message])
        if letters_count > max_letters_count or best_candidate is None:
            key = candidate_key
            best_candidate = candidate_message
            max_letters_count = letters_count

    return best_candidate, key


"""
Now try a repeating-key XOR cipher. 
E.g. it should take a string \xe2\x80\x9chello world\xe2\x80\x9d and, given the key is \xe2\x80\x9ckey\xe2\x80\x9d, 
xor the first letter \xe2\x80\x9ch\xe2\x80\x9d with \xe2\x80\x9ck\xe2\x80\x9d, 
then xor \xe2\x80\x9ce\xe2\x80\x9d with \xe2\x80\x9ce\xe2\x80\x9d, 
then \xe2\x80\x9cl\xe2\x80\x9d with \xe2\x80\x9cy\xe2\x80\x9d, 
and then xor next char \xe2\x80\x9cl\xe2\x80\x9d with \xe2\x80\x9ck\xe2\x80\x9d again, 
then \xe2\x80\x9co\xe2\x80\x9d with \xe2\x80\x9ce\xe2\x80\x9d and so on. 
You may use an index of coincidence, Hamming distance, Kasiski examination, statistical tests or whatever method you feel would show the best result.
"""
def attack_repeating_key_xor(ciphertext: str) -> typ.Tuple[bytes, bytes]:
    keysize = find_vigenere_key_length(ciphertext)

    key = bytes()
    message_parts = list()
    for i in range(keysize):
        msg, msg_key = attack_single_byte_xor(bytes(ciphertext[i::keysize]))
        key += msg_key
        message_parts.append(msg)

    # rebuild original message by combining bytes
    message = bytes()
    # iterate over range of max message part length
    for i in range(max(map(len, message_parts))):
        message += bytes([part[i] for part in message_parts if len(part) >= i + 1])

    return message, key
