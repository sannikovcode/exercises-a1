num = int(input())
last_digit = num % 10
first_digit = num // 10

print('Искомое число =', last_digit * 10 + first_digit)
