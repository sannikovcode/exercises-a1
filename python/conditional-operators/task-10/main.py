num = int(input())


x1 = num // 1000
x2 = (num // 100) % 10
x3 = (num // 10) % 10
x4 = num % 10

sumx1_x4 = x1 + x4
difx2_x3 = x2 - x3

if sumx1_x4 == difx2_x3:
    print("Да")
else:
    print("Нет")    
    