import json

f = open("../plaintext/plaintext_array.json")
data = json.load(f)
f.close()
print(data)

f = open("../plaintext/plaintext_object.json")
data = json.load(f)
f.close()
print(data)
