Concurrent multi-user access - Multiple users with different roles can access the website
now, but 2 different users cannot access the website from the same browser, if multiple
users have to access the website it has to be different browser or incognito mode, because
the session info is stored at client side hence different users can access the website from
different browsers. This is also present in all the major commercial websites.  

For doing concurrent updates in the database, the database will be consistent because
MySQL uses InnoDB engine which by default does Row-Level locking while doing inserts
and updates. Also in our database the queries are written in such a way that the internal
locking will be sufficient to prevent any race conditions.  

Ref - https://dev.mysql.com/doc/refman/8.3/en/internal-locking.html  

We also tried updating the amount due value concurrently using threading and found that the
final value was consistent with the result that was expected. This proved the fact that InnoDB
maintains the row level locking while updating and inserting.  

### The Query
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/0d6395e2-c528-46c9-aab7-937614622dce)
 

### Before doing concurrent updates:
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/d42e0c6e-e603-4578-8e1b-aeb6adabfe47)  

### After running concurrent script:
![image](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/725b5c1f-fea9-41ba-9189-86bf93ba9994)

The concurrent script takes the value of Amount_Due and does +1 to the value obtained and
writes it into the database. If there would be a problem then the value should be inconsistent
but this shows that InnoDB maintains the consistency between the queries.
To do the query you need to change the database config and also the update query should
be set accordingly where the user id exist





