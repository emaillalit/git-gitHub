n = 5
while n > 0:
    print n
    n = n - 1
print 'Blassoff'
print n


n = 0
while n > 0:
    print 'Lather'
    print 'Rinse'

print 'Dry off'


while True:
    line = raw_input('>')
    if line == 'done':
        break
    print line
print 'Done!'

while True:
    line = raw_input('>')
    if line[0] == '#':
        continue
    if line == 'done':
        break
    print line
print 'Done!'

for i in range(10):
    print i
print 'Blastoff!'

friends = ['joseph', 'glen', 'sally']

for friend in friends:
    print 'Happy New Year:', friend
print 'Done!'

s = 'Monty Python'
print s[0:4]
print s[6:7]
print s[6:20]
print s[-1:-7:-1]
print s[:]
print s[::-1]
