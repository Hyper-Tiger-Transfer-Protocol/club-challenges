require 'json'
require 'pp'

json = File.read('../plaintext/plaintext.json')
obj = JSON.parse(json)

pp obj