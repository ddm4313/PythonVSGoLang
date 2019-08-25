import requests
import time

combo = open("/home/den/go_projects/src/hello/test.txt", "r").readlines()

def main():
    start_time = time.time()
    for line in combo:
        username, password = line.split(":")
        t = {"username": "fsafsafas", "password": "fasfsafas"}
        r = requests.post("https://accounts.spotify.com/password/login", data=t)
        print(r.text)
    elapsed_time = time.time() - start_time
    print(elapsed_time)
        
main()
