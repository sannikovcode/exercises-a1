num = int(input())

c = num % 10          # третья цифра
b = (num // 10) % 10  # вторая цифра  
a = num // 100        # первая цифра

# Вариант 1: Через строки (проще)
abc = f"{a}{b}{c}"
acb = f"{a}{c}{b}"
bac = f"{b}{a}{c}"
bca = f"{b}{c}{a}"
cab = f"{c}{a}{b}"
cba = f"{c}{b}{a}"

print(abc)
print(acb)
print(bac)
print(bca)
print(cab)
print(cba)
