#!/usr/bin/python

"""
listA = ['hello', 'hell', 'help', 'helicopter', 'doesnotmatch']

listB = []
for i in listA:
    first_character = i[0:1]
    listB.append(first_character)


for i in listB:
    for j in listA:
        if i == j[0:1]:
"""

listC = [3, 4, 7, 8, 10, 3, 4]
input2 = 7
i  = 0
while i < len(listC):
    print(i)
    output = ''
    for j in listC:
        output = j + listC[i]
        print(output)
    print('------------------------')
    i += 1
