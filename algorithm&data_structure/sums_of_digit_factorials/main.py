import timeit


def factorial(n):
    result = 1
    n = int(n)

    for i in range(1, n+1):
        result = result * i
    return result


def f(n):
    result = 0
    for i in range(len(str(n))):
        result = result + factorial(str(n)[i])
    return result


def sf(n):
    result = 0
    f_result = f(n)
    for i in range(len(str(f_result))):
        result += int(str(f_result)[i])
    return result


def g(n):
    i = 1
    while True:
        if sf(i) == n:
            return i
        i += 1


def sg(n):
    result = 0
    for i in range(1, n+1):
        g_result = g(i)
        for j in range(len(str(g_result))):
            result += int(str(g_result)[j])
        # print("g({}) : {}".format(i, g_result))
    return result


start = timeit.default_timer()
print(sg(30))
stop = timeit.default_timer()

print('Time: ', stop - start)
