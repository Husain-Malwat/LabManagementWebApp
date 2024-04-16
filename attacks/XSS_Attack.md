# XSS Attck
We have performed a cross-site scripting attack on our register page.  

* The vulnerability found here was to be used to insert inline JavaScript. The email we provide 
is being evaluated using a general email ID format. This function is dynamic and changes when updating inputs.
![WhatsApp Image 2024-04-16 at 17 35 47_67862c4c](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/e6615747-4a62-46ab-9d45-0d79a00b3964)

* Now we insert inline javascript to know the password written by the user, and the attacker can view the password even when it's hidden(shown as dots). The attacker could also
register by confirming the password, and then he can log in.
![WhatsApp Image 2024-04-16 at 17 35 30_bda7ec5c](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/be8bb9b2-f851-4085-886c-9c5342c6348b)
![WhatsApp Image 2024-04-16 at 17 35 09_094a3f43](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/d7dfe606-f5b4-4481-924f-029e0a5d03cd)


# Defence Against XSS Attack
To defend against this attack, we use the email address as a string so that the attacker cannot manipulate the internal variables.
![WhatsApp Image 2024-04-16 at 17 36 31_06164759](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/e097ecec-e27f-4865-812a-cf1ba34283ba)

After making changes:
![WhatsApp Image 2024-04-16 at 17 44 38_4a5bcb8c](https://github.com/kaushal-003/LabManagementWebApp/assets/114944809/a549aa5a-95df-40f5-bbd6-fbac1e7411c1)
 
