import json

with open('data.json', 'r') as fh:
    data = json.load(fh)

print(json.dumps(data, indent=4, sort_keys=True))

