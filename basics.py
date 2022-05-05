with open('/var/tmp/fruits.txt', 'r') as fruits:
    content = fruits.read(2)

print(content)

with open("vegetables.txt", "w") as myfile:
    myfile.write("\nCucumber\nOnion\n")
    myfile.write('Garlic')

