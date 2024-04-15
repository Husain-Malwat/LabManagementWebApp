### SQL Injection
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





