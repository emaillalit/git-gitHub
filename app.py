import json

def translate(word):
    word = word.lower()
    with open("data.json") as fh:
        data = json.load(fh)
        if word in data:
            return data[word]
        else:
            return "The word doesn't exist. Please double check it."
# Take User input.
user_input = input('Enter a word: ')
print(translate(user_input))