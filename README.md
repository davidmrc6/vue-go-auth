# JWT Authentication Example with Vue+Go
Simple JWT authentication example with Go on the backend and Vue on the frontend.
\
To get started, clone the repository with
```bash
git clone https://github.com/davidmrc6/vue-go-auth.git
```
and set up your database connection by following these steps:
1. Navigate to `./backend`
2. Create a `.env` file (in the `backend` folder)
3. Set up your locally hosted PostgreSQL and fill in these environmental variables inside the `.env` file:
```.env
DB_HOST=localhost
DB_PORT=5432
DB_USER=<your database username>
DB_PASSWORD=<your database password>
DB_NAME=<your database name>
DB_SSLMODE=disable
```
Then, start both the frontend and the backend servers by first navigating to the root directory and running (on separate terminals):
```bash
cd frontend/
npm run dev
```
```bash
cd backend/
go run main.go
```
