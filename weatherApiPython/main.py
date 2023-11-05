import requests
import json

with open('config.json') as config_file:
    config= json.load(config_file)

BASE_URL= config['BASE_URL']
API_KEY = config['API_KEY']

city=input("Where do you want to check the weather : ")
request_url=f"{BASE_URL}?appid={API_KEY}&q={city}"

response = requests.get(request_url)

if response.status_code==200:
    data=response.json()
    weather=data['weather'][0]['description']
    print("Weather :", weather)
    temp=round(data['main']['temp']-273.15,2)
    print("Temperature :",temp,"ºC")
else:
    print("An error occured.")