#!/bin/python3

import sys
from functools import reduce


def get_result(start, n, num, value):
    if start >= len(num) - 1 - n:
        return value
    else:
        second = reduce(lambda x, y: x*y, map(int, str(num)
                        [start: start + n]))
        print(second, start, sep=",")
        return get_result(start+1, n, num, value if value > second else second)


t = int(input().strip())
for a0 in range(t):
    n, k = input().strip().split(" ")
    n, k = [int(n), int(k)]
    num = input().strip()
    # start = reduce(lambda x, y: x*y, map(int, str(num)[:k]))
    print(start)
    # print(get_result(1, k, num, start))
