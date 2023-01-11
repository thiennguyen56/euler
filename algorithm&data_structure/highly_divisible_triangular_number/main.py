
terms = []

step = 1
start = 0
list_factor = []
while len(list_factor) <= 500:
    terms.append(start + step)
    list_factor = [i for i in range(
        1, start + step + 1) if (start + step) % i == 0]

    step += len(terms) + 1
    print(terms[-1])
    print(list_factor)
