a = int(input())
b = int(input())
c = int(input())

counter = 0

if a % 2 == 0:
    counter = counter + 1 
if b % 2 == 0:
    counter = counter + 1 
if c % 2 == 0:
    counter = counter + 1 

print("Четные:", counter)