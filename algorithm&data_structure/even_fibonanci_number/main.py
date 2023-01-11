#!/bin/python3

import sys


t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    f1, f2 = 1, 2
    total = 0
    while n > f2:
        if f2 % 2 == 0:
            total += f2
        f1, f2 = f2, f1+f2
    print(total)
