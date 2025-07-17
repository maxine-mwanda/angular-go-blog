Prerequisites
Before you begin, ensure you have the following installed:

Go (v1.16+) - Download

Node.js (optional, for frontend tooling) - Download

Git (optional, for cloning) - Download

Step 1: Clone or Download the Project
Option A: Clone via Git
bash
git clone https://github.com/yourusername/blog-app.git
cd blog-app
Option B: Download ZIP
Download the project ZIP from GitHub.

Extract it to a folder (e.g., ~/code/blog-app).

Step 2: Backend Setup (Golang)
1. Navigate to the Backend Directory
bash
cd backend
2. Install Dependencies
bash
go mod tidy
3. Run the Go Server
bash
go run main.go
The server should start on http://localhost:8080.

You should see logs like:

text
2025/07/17 21:03:31 Connected to database
2025/07/17 21:03:31 Server running on port 8080

Step 3: Frontend Setup (AngularJS)
1. Navigate to the Frontend Directory
bash
cd ../frontend
2. (Optional) Install Frontend Dependencies
If you modify frontend files and need tooling (e.g., npm):

bash
npm install
3. Access the Application
Open your browser and go to:

text
http://localhost:8080
You should see the AngularJS blog homepage.

Step 4: Verify Everything Works
Test API Endpoints
List all posts:

text
http://localhost:8080/api/posts
(Should return JSON data.)

Get a single post:

text
http://localhost:8080/api/posts/sample-post
(Replace sample-post with an actual slug from your database.)

Test Frontend Navigation
Click on blog post titles to navigate to individual posts.

Use the back button to return to the homepage.

Troubleshooting
1. If index.html is not loading
Check if the frontend files are in the correct location:

bash
ls ../frontend
(Should show index.html, app/, assets/.)

If the path is wrong, update backend/routes/routes.go:

go
frontendPath := "/your/correct/path/to/frontend"
2. If the API returns 404
Ensure the Go server is running (go run main.go).

Check if blog.db exists in the backend folder.

3. If AngularJS routes don’t work
Clear browser cache (Ctrl+Shift+Del).

Ensure index.html has:

html
<base href="/">

Project Structure
text
/blog-app
├── backend/          # Golang server
│   ├── config/      # Config files
│   ├── controllers/ # API handlers
│   ├── db/          # Database setup
│   ├── models/      # Data models
│   ├── routes/      # URL routing
│   └── main.go      # Server entry point
├── frontend/        # AngularJS frontend
│   ├── app/         # AngularJS app logic
│   ├── assets/      # CSS, images
│   └── index.html   # Main HTML file
└── README.md        # Setup guide
