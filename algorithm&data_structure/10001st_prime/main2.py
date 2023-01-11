x = 110000
primes = [True]*x
p = 2
while p <= x**.5:
    if primes[p]:
        for i in range(p+p, x, p):
            primes[i] = False
    p += 1
primeNos = []
for i in range(2, len(primes)):
    if primes[i]:
        primeNos.append(i)
