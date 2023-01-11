def get_result(n):
    for number in range(n, 100000, -1):
        standard = str(number)
        reverse = str(number)[::-1]
        if standard == reverse:
            for i in range(100, 1000):
                if number % i == 0 and 100 <= (number // i) <= 999:
                    print(i)
                    return number


t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    print(get_result(n))
