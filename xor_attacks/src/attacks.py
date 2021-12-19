import typing as typ

from src.misc import ASCII_RANGE, bxor


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
