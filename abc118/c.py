import math
import itertools
from functools import reduce
# print(math.gcd(1000000000, 1000000000))


def gcd(*numbers):
    return reduce(math.gcd, numbers)

N = int(input())
monsters = list(map(int, input().split()))

result = 10 ** 9 + 1

print(gcd(*monsters))
