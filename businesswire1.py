
import re

with open('logs.txt', 'r') as fd:
    listB = fd.readlines()

regcom = re.compile(r'\.com')
regnet = re.compile(r'\.net')
regio = re.compile(r'\.io')
seen = {}
for i in listB:
    i = i.strip()
    if regcom.search(i) or regnet.search(i) or regio.search(i):
        if i not in seen:
            seen[i] = 1
        else:
            seen[i] += 1

prin(seen)
        

