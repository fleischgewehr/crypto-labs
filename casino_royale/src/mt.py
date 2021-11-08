import typing as typ


class MT19937:
    """https://en.wikipedia.org/wiki/Mersenne_Twister"""

    # 32 bit implementation (word size)
    w = 32
    # has state of 624 capacity (degree of recurrence)
    n = 624
    # middle word, an offset used in the recurrence relation
    # defining the series x, 1 ≤ m < n
    m = 397
    # separation point of one word, or the number of bits
    # of the lower bitmask, 0 ≤ r ≤ w − 1
    r = 31
    # coefficients of the rational normal form twist matrix
    a = 0x9908b0df
    # additional Mersenne Twister tempering bit shifts/masks
    u, d, l = 11, 0xffffffff, 18
    # TGFSR(R) tempering bitmasks (b, c)
    # & TGFSR(R) tempering bit shifts (s, t)
    s, b = 7, 0x9d2c5680
    t, c = 15, 0xefc60000
    # forms another parameter to the generator, though not part
    # of the algorithm proper. The value for f for MT19937
    # is 1812433253 and for MT19937-64 is 6364136223846793005
    f = 1812433253

    @property
    def lower_mask(self) -> int:
        return (1 << self.r) - 1

    @property
    def upper_mask(self) -> int:
        return (~self.lower_mask) & self.d

    def __init__(self, seed: int) -> None:
        self.state = [seed]
        self.idx = self.n + 1
        for i in range(1, self.n):
            val = (
                self.f * (self.state[i - 1] ^ (self.state[i - 1] >> (self.w - 2))) + i
            ) & self.d
            self.state.append(val)

    def twist(self) -> None:
        for i in range(self.n):
            x = (
                (self.state[i] & self.upper_mask)
                + (self.state[(i + 1) % self.n] & self.lower_mask)
            )
            # print(f'{i=}: {x=}')
            xA = x >> 1
            if (x % 2) != 0:
                # print('lowest bit is equal to 1')
                xA = xA ^ self.a

            self.state[i] = self.state[(i + self.m) % self.n] ^ xA
        
        self.idx = 0

    def __iter__(self) -> typ.Iterator[int]:
        while True:
            if self.idx >= self.n:
                self.twist()
            
            y = self.state[self.idx]
            y = y ^ ((y >> self.u) & self.d)
            y = y ^ ((y << self.s) & self.b)
            y = y ^ ((y << self.t) & self.c)
            y = y ^ (y >> 1)

            self.idx += 1
            yield y & self.d
