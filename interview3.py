import sys
import os

dir_path = input('Enter a directory path: ')
if not os.path.exists(dir_path):
    sys.exit(f"Following {dir_path} doesn't exist. Please enter valid path")

cppfiles = []
for dirpath, dirs, files in os.walk(dir_path):
    for file in files:
        if file.endswith('.cpp'):
            fullpath = os.path.join(dirpath, file)
            with open(fullpath, 'r') as f:
                print('File -> %s' %fullpath)
                for i in f:
                    print(i.strip())
                    
                

