import re

message = 'Call me 415-555-1011 tomorrow, or at 415-555-9999'

preg = re.compile(r'\d\d\d-\d\d\d-\d\d\d\d')

mo = preg.search(message)
print(mo.group())

fall = preg.findall(message)
print(fall)
