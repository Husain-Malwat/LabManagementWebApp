# SQL Injection
We will perform SQL Injection on the login page of our web app.

We observe how we can proceed with attacking the vulnerability.
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/64beaf7c-419f-40c4-bde1-53135c6c35a2)

The query takes the email and the password and fetches the first result. As we know, an entry would have
a single email and corresponding password. This query would give a single row in general use, so we fetch a single result.  

We inserted our query with SQL code into the email and successfully logged in as a user.  
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/3d837fa2-f53f-47b3-be10-886ddb736c1e)  
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/839132cb-7541-40de-aac2-d953df8b88a0)


Observe what happens if we run this query in MySQL. It fetches the first result of the student's table and logs in to the first user's account.
Thus, we can log in to the first student's account in the student table using SQL Injection.
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/26dd4dcc-cde4-41d9-9f3f-2c1556bb79b7)

# How To Prevent SQL Injection
We implement a simple change in the app.py code. Instead of passing direct values in the executable SQL line, we will pass parameterized values, ensuring 
each attribute gets the same value as given in the parameter, which fails SQL Injection.
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/131673e5-edc9-4677-9d2e-1909fb0bae16) 

Now on giving the faulty email, we get redirected to the same page again:
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/f731d754-81c7-4a2a-b1d5-855f5e5e2ebe)


 



