# Bloggy

This is a full-stack web application consisting of a Golang backend API and a Next.js frontend.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Setting Up the Backend](#setting-up-the-backend)
  - [Setting Up the Frontend](#setting-up-the-frontend)
- [Usage](#usage)
- [API Routes](#api-routes)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Before you begin, ensure you have met the following requirements: 

- Go (Golang) installed
- Node.js and npm installed
- MongoDB installed and running
- Your favorite code editor (e.g., Visual Studio Code)

## Getting Started

Follow these steps to set up and run the project on your local machine.

### Setting Up the Backend

1. Clone this repository:

   ```bash
   git clone https://github.com/niyioo/bloggy.git

2. Navigate to the Golang backend folder:

cd my-blog-api

3. Install Go dependencies:

go mod download

4. Update the MongoDB connection settings in the main.go file:

// Update your MongoDB URI based on your MongoDB setup
mongoURL := "mongodb://localhost:27017" // Change this to your MongoDB server URI

5. Start the Golang backend:

go run main.go


### Setting Up the Frontend

1. Navigate to the Next.js frontend folder:

cd my-blog-app

2. Install Node.js dependencies:

npm install

3. Update the API URL in your Next.js app to point to your Golang backend.

const apiUrl = 'http://localhost:8080'; // Update this with your Golang backend URL

4. Start the Next.js development server:

npm run dev


## Usage

Access your Next.js frontend at http://localhost:3000.
Access your Golang backend API at http://localhost:8080.

## API Routes

/api/register - Register a new user.
/api/login - Authenticate a user and get an access token.
/api/blogposts - List and create blog posts.

You can find more API routes and details in the Golang backend code.
# bloggy
