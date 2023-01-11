#!/bin/python3

import sys
import math


def get_result(n):
    sum_square_n = (n*(n+1)*(2*n+1)) // 6
    square_sum_n = int(math.pow((n * (n + 1) / 2), 2))
    return square_sum_n - sum_square_n


t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    print(get_result(n))
