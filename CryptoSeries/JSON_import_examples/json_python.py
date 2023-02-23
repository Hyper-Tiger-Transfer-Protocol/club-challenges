import json

f = open("../plaintext/plaintext.json")

data = json.load(f)

for string in data:
    print(string)
