import requests

res = requests.get("https://localhost:8443", verify='ca/ca.crt')
print(res.text)
