### BruteForce Password hijacking

To recreate the attack, please follow the steps mentioned below.  

1. Open the ``BruteForce.py`` file.
2. Insert the email of a user for which you'd like to perform bruteforce attack.
3. Save the ``BruteForce.py`` file.
4. Run:
```
python BruteForce.py
```

For submission purposes, We have performed a brute-force attack on the Email ``anishkarnik@iitgn.ac.in``, and the password length is 3 for this case.

By running the script, we got the password in about 3 minutes.  
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114857798/488b3c0d-d322-4e46-88cb-7aadbc2f7c69)

### Note:-
- One can add more letters in the script if you think the password may contain special characters.
- Also length of the password can be increased if the password is not in the given range.

# Defence against BruteForce attacks:-

-  Rate limiters:-
   - Every time a request exceeds the rate limit, the view function will not get called, and instead, a 429 HTTP error will be raised[1].
   - To Perform BruteForce attacks, Adversary has to try all the possible permutations, for which he/she needs to send thousands of requests to the server.

- We applied rate limiters to our website as a defence against this attack.
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114857798/f051cc3d-8a7b-4a91-a30e-63ef652055ce)

After applying rate limiters:-
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114857798/e57b8561-af10-44cd-b9d1-6a86b084bf0a)



### References;-
[1] https://flask-limiter.readthedocs.io/en/stable/





