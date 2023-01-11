import math

s = 4
m = -1
for a in range(3, (s//3) + 1):
    b = (s**2 - 2*s*a) // (s*2 - 2*a)
    c = s - a - b
    if a**2 + b**2 == c**2 and a*b*c > m:
        m = a*b*c
print(m)
