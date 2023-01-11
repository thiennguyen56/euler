form = "1_2_3_4_5_6_7_8_9_0"


i = 100000000


def check_number(n):
    if len(str(n)) != len(form):
        return False
    for i in range(0, len(str(n)), 2):
        if form[i] != str(n)[i]:
            print(i)
            return False
    return True


while check_number(i*i) is False:
    print(i)
    i += 10

print(i)
