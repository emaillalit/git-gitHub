names = ['joseph', 'simi', 'mathew']
stuff = ['slice', 'bread', 'tomato']

print(names)
print(stuff)

some = [1, 91, 21, 10, 6]

for e in names:
    print('Each element -> %s' %e)
for e in stuff:
    print('Each element in stuff -> {}'.format(e))
print(max(some))
print(min(some))
print(sum(some))
print(len(some))
print(sum(some)/len(some))

# Slice func.
str1 = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
print(str1[1:8])
print(str1[8:20])
print(str1[:])
print(str1[::-1])
print(str1[-6:-3])
print(str1[-3:-8:-1])

print(str1[-1])
"""
numlist = []
while True:
    inp = input('Enter a number: ')
    if inp == 'done':
        break
    val = int(inp)
    numlist.append(val)
average = sum(numlist)/len(numlist)
print('Average: {}'.format(average))
"""
# Read a file

allwords = []
with open('local.py', 'r') as fh:
    for line in fh: 
        line1 = line.strip()
        print('Each line -> {}'.format(line1))
        if not line1.startswith('print'):
            continue
        words = line1.split()
        allwords.append(words)
print(allwords)


# Dictionaries
purse = {'money': 12, 'tissues': 75, 'candy': 3}
print(purse)
print('ADD')
purse['candy'] = purse['candy'] + 2
print(purse)

counts = {}
names = ['csev', 'cwen', 'csav', 'zquian', 'cwen']
for name in names:
    if name not in counts:
        counts[name] = 1
    else:
        counts[name] = counts[name] + 1
print(counts)
for k, v in counts.items():
    if v > 1:
        print('most %s '%(k))


counter = {}
statement = input('Enter a line of txt: ')
words = statement.split()
for word in words:
    counter[word] = counter.get(word,0) + 1
print(counter)

for k, v in counter.items():
    print ('Key -> %s, Value -> %s' %(k, v))

# Tuples

x = ('Glenn', 'Sally', 'Joseph')
print(x)
for i in x:
    print('Each el in x tuples -> %s' %(i))

d = {}
d['csev'] = 2
d['cwen'] = 4
for k, v in d.items():
    print(f'key - {k} value - {v}')
tups = d.items()
print(f'tups - {tups}')
d['amy'] = 8
print(tups)

print(sorted([ (v, k) for k,v in d.items() ] ))
