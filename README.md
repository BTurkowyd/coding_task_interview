# Info
- To run the application, Docker is required.
- The application is organized in three folders, and each will be a separate docker image:
    backend (Go)
    frontend (Node.js, React)
    database (PostgreSQL)
- Database consists of two tables: "users" and "bike_renting_system", the latter is populated with three bikes which can be rent. The "users" table is empty.

## Installation
- Open the terminal in the main folder which contains the docker-compose.yml file and run the command docker-compose up. 
-  Optionally you can use -d flag to run it in a detached mode
- After all images are built, they will run automatically and each service will run on the following ports:
    backend -> :3000
    frontend -> :3001
    database -> :5432

# Use
- Upon entering http://localhost:3001/ webpage will move to the login window, which has a hyperlink to a simple registration form (user password, for simplicity e-mail option is omitted). Register user. Upon registration application will direct back to the login window where you can log in.
- Entire application relies on a cookie session, e.g. list of bikes is available only to logged in users.
- Upon login, the page with all bikes is displayed, where users can rent/return bikes. Each bike has a hyperlink to its own simple page displaying all information.
- At the bottom of the page a logout button is present, which clears the user session and redirects to the login page.