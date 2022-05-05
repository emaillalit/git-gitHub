import yaml

with open('data.yaml', 'r') as fh:
    data = yaml.load(fh, Loader=yaml.FullLoader)

print(data)
