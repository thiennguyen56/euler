
import sys
import math


def get_list_prime(n):
    i = 2
    maxi = 0
    result = []
    while i*i <= n:
        while n % i == 0:
            n /= i
            maxi = i
            result.append(maxi)
        i += 1
    if n > maxi:
        result.append(int(n))
    return result


def get_result(n):
    dict_prime = {
    }
    for i in range(1, n + 1):
        list_prime = get_list_prime(i)
        # print(i, list_prime)
        for i in set(list_prime):
            count = list_prime.count(i)
            if i not in dict_prime:
                dict_prime.update({
                    i: count
                })
            elif count > dict_prime.get(i):
                dict_prime[i] = count
    total = 1
    for k, v in dict_prime.items():
        total *= math.pow(k, v)
    return int(total)


t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    print(get_result(n))
