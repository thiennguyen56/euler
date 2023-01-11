def unit_sum(n):
    if n == 0:
        return 0
    return n*(n+1) // 2


t = int(input().strip())
for a0 in range(t):
    n = int(input().strip()) - 1
    print(3 * unit_sum(n // 3) + 5 * unit_sum(n // 5) - 15 * unit_sum(n // 15))
