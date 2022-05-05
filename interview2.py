listA = ['c','a','i','o','p','a']
"""
count = {}
listB = []
for i in listA:
    if i not in count:
        count[i] = 1
    else:
        count[i] = count[i] + 1

for k,v in count.items():
    if v >= 2:
        listB.append(k)

print(listB)
"""
count = {}
my_set = {'c','a','i','o','p','a'}

for i in my_set:
    if i in my_set:
        print(i)
    else:
        print('sd {}'.format(i))
