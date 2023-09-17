# Fiber CRM
Fiber CRM is a simple Customer Relationship Management (CRM) system built with Go and the Fiber web framework. This project allows you to manage leads, including creating, retrieving, updating, and deleting lead information via a RESTful API. It uses an SQLite database for data storage and is designed to be easy to set up and use.
# Features
1. Create Leads: Add new leads to the CRM system, including their name, company, email, and phone number.
2. Retrieve Leads: Fetch a list of all leads or retrieve a specific lead by their unique ID.
3. Update Leads: Update the information of existing leads, such as their name, company, email, or phone number.
4. Delete Leads: Delete leads from the CRM by their ID.
# Getting Started
Follow these steps to set up and run the Fiber CRM project on your local machine:
1. Clone the Repository:
git clone https://github.com/yourusername/fiber-crm.git
cd fiber-crm
2. Install Dependencies:
go get -u github.com/gofiber/fiber
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/sqlite
3. Initialize the Database:
go run main.go migrate
4. Start the Application:
go run main.go
The application will be available at http://localhost:3505.
# API Endpoints
1. GET /api/v1/leads: Fetch a list of all leads.
2. GET /api/v1/leads/:id: Retrieve a specific lead by ID.
3. POST /api/v1/leads: Create a new lead.
4. PUT /api/v1/leads/:id: Update an existing lead by ID.
5. DELETE /api/v1/leads/:id: Delete a lead by ID.
# Acknowledgments
1. [Fiber](https://github.com/gofiber/fiber) - A fast and Express.js-compatible web framework for Go.
2. [GORM](https://gorm.io/) - The fantastic ORM library for Golang.
3. [SQLite](https://www.sqlite.org/index.html) - A C-language library for implementing a small, fast, self-contained, high-reliability, full-featured, SQL database engine.
# Contributing
Contributions to Fiber CRM are welcome! If you'd like to contribute, please follow these guidelines:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure the code is properly formatted.
4. Write tests for new features or bug fixes.
5. Submit a pull request with a clear description of your changes.
