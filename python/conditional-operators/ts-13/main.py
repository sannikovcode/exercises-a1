a = int(input())
b = int(input())
c = int(input())
v = int(input())

x = a
if b < x:
    x = b
if c < x:
    x = c
if v < x:
    x = v

print(x)
