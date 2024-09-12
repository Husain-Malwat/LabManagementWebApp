# Laboratory Management System

![Screenshot 2024-08-12 150443](https://github.com/user-attachments/assets/7eaedbbb-9224-4938-8556-2d9c71b45f81)

## Overview

The **Laboratory Management System** is an open-source web application designed to streamline manage laboratory operations, create and enroll for courses,and perform bookings efficiently. This application provides features for managing lab inventory, managing lab users, making lab bookings, issue/return any resource. Built with Flask and MySQL, the system incorporates Golang for handling concurrency, ensuring that multiple users and operations can be managed simultaneously without performance degradation.

## Features

- **Inventory Management:** Track and manage lab equipment, chemicals, and supplies.
- **User Management:** Manage user roles and permissions within the lab environment.
- **Google Authentication:** Secure login using Google OAuth.
- **Auto-Generated Emails:** Automatically send emails for issuing and returning lab resources.
- **Lab Bookings:** Reserve lab equipment and resources for specific times.
- **Course Management:** Create and enroll in courses designed for students.
- **Concurrency Support:** Golang is utilized to efficiently manage concurrent operations.
- **Responsive Design:** Accessible on various devices, including desktops, tablets, and smartphones.

## Tech Stack

- **Backend:**
  - ![Flask](https://img.shields.io/badge/Flask-000000?style=for-the-badge&logo=flask&logoColor=white)
  - ![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) 
- **Database:**
  - ![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
- **Frontend:**
  - ![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
  - ![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
  - ![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/lab-management-system.git
cd lab-management-system
```

### 2. Create and activate a virtual environment:
```bash
python3 -m venv venv
source venv/bin/activate  # On Windows use `venv\Scripts\activate` 
```
### 4. Install the required dependencies:
```bash
pip install -r requirements.txt
```
### 5. Set up the MySQL database:
- Create a new database for the application.
- Update the database configuration in config.py.
### 6. Run the application:
```bash
flask run
```
### 7. Run the Golang concurrency service:
```bash
go run concurrency_service.go
```
### 8. Access the application:
- Open your browser and go to http://127.0.0.1:5000/.

## Usage

- **Admin Panel:**
  - Accessible to users with admin roles for managing inventory, users, and experiments.

- **Lab Users:**
  - Can log in to view and manage experiments, update their progress, and view reports.

## Contributing

We welcome contributions! Please follow these steps:

- Fork the repository.
- Create a new branch (`git checkout -b feature-branch`).
- Make your changes.
- Commit your changes (`git commit -m 'Add some feature'`).
- Push to the branch (`git push origin feature-branch`).
- Open a Pull Request.

## License

- This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
