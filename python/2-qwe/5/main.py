a = int(input())

x3 = a % 10 
x2 = (a // 10) % 10
x1 = a // 100

sumx = x1 + x2 + x3
mullx = x1 * x2 * x3


print("Сумма цифр =", sumx)
print("Произведение цифр =", mullx)