require 'json'
require 'pp'

json = File.read('../plaintext/plaintext_array.json')
list = JSON.parse(json)

pp list

json = File.read('../plaintext/plaintext_object.json')
obj = JSON.parse(json)

pp obj