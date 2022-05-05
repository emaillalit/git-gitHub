def div42by(divideBy):
    try:
        return 42 / divideBy
    except ZeroDivisionError:
        print ('Error: You tried to divide by zero.')

print(div42by(2))
print(div42by(12))
print(div42by(0))
print(div42by(1))

numCats = raw_input('How many cat do you have?')
try:
    if int(numCats) >= 4:
        print('Thats alot of cats.')
    elif int(numCats) < 0:
        print('You entered a negative number.')
    else:
        print('Thats not alot of cats.')
except ValueError:
    print('You did not enter a number.')
