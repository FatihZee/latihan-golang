# Golang Practice Project

## About Me
I am currently practicing Golang and developing a structured backend system using MVC architecture.

## Project Overview
This repository contains a simple backend application built with Go, featuring:
- **User Management**: Registration, login, and CRUD operations for users.
- **Product Management**: CRUD operations for products.
- **API Endpoints**: Built with the **Gin** framework for handling requests.
- **Database Integration**: Uses **GORM** with MySQL for data storage.

## Project Structure
```
📂 latihan-golang  
 ┣ 📂 controllers  # Handles request logic  
 ┃ ┣ 📜 userController.go  
 ┃ ┣ 📜 productController.go  
 ┣ 📂 models  # Defines database models  
 ┃ ┣ 📜 user.go  
 ┃ ┣ 📜 product.go  
 ┣ 📂 routes  # API route definitions  
 ┣ 📂 config  # Database connection  
 ┣ 📜 main.go  # Entry point  
```
```

## How to Run
1. Clone this repository:
   ```
   git clone <REPO_URL>
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Run the application:
   ```
   go run main.go
   ```

## Technologies Used
- **Golang**
- **Gin (Web Framework)**
- **GORM (ORM for Go)**
- **MySQL (Database)**

Happy coding! 🚀

