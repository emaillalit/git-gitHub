import re, pyperclip

phoneRegex = re.compile(r'''
# 415-555-5500
(
((\d\d\d) | (\(\d\d\d\)))?
(\s|-)
\d\d\d
-
\d\d\d\d
(((ext(\.)?\s)|x) 
(\d{2,5}))?
)
''', re.VERBOSE)

emailRegex = re.compile(r'''
# some.+_thing@(\d{2,5}))?.com
[a-zA-Z0-9_.+]+
@
[a-zA-Z0-9_.+]+

''', re.VERBOSE)

text = pyperclip.paste()

extracphone = phoneRegex.findall(text)
extracemail = emailRegex.findall(text)