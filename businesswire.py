
with open('logs.txt', 'r') as fd:
    listB = fd.readlines()

seen = {}
for i in listB:
    i = i.strip()
    if 'yahoo.com' in i:
        if i not in seen:
            seen[i] = 1
        else:
            seen[i] += 1
    elif 'google.com' in i:
        if i not in seen:
            seen[i] = 1
        else:
            seen += 1
    elif 'google.net' in i:
        if i not in seen:
            seen[i] = 1
        else:
            seen[i] += 1
    elif 'info.io' in i:
        if i not in seen:
            seen[i] = 1
        else:
            seen[i] += 1
    elif 'msn.com' in i:
        if i not in seen:
            seen[i] = 1
        else:
            seen[i] += 1

print(seen)
