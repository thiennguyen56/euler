with open("text.txt", "r") as f:
    total = 0
    lines = f.readlines()
    for i in lines:
        total += int(i.strip())
    print(str(total)[:10])
