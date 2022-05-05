import random

print("Hello, what is your good name? ")
name = input()
secretNumber = random.randint(1, 20)
print('Well %s Iam thinking the number between 1 and 20.' %name)

for i in range(1, 7):
    print('Take a guesss.')
    guess = int(input())
    if guess < secretNumber:
        print('You guess low.')
    elif guess > secretNumber:
        print('you guess high.')
    else:
        break
if guess == secretNumber:
    print('Good job %s ! You guessed my number in %s try.' %(name, i))
else:
    print('Nope. The number I was thinking of was %s ' %secretNumber)
