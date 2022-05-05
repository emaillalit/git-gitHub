import re

regex = re.compile(r'\d\d\d-\d\d\d-\d\d\d\d')

mo = regex.search('my number is 415-555-4242')

#print(mo.group())

# ? = zero or one
# * = zero or more
# + = one or more
# {3} = match 3 times
# {3,5} = match min 3 to 5.
string1 = "my number is 415-555-4242"
regex = re.compile(r'(\d\d\d)-(\d\d\d-\d\d\d\d)')
listA = regex.findall(string1)
print(listA)

regexip = re.compile(r'''
(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
  (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
  (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
  (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)''', re.IGNORECASE | re.DOTALL | re.VERBOSE)

regexip2 = re.compile(r'''
(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
  (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
  (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
  (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)''', re.VERBOSE)
#string2 = ('10.11.12.1 182.192.12.234')

string2 = '10.11.12.1 182.192.12.234 192.168.1.22 172.168.1.33 18.11.23.23'

#listB = regexip.search(string2)
#print(listB.group())
listB = regexip2.search(string2)
print(listB.group())


