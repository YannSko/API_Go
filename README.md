# API Project Overview:
This API provides functionalities for managing houses, with features like CRUD operations, JWT-based authentication, rate limiting, caching, and more. It allows users to perform actions such as retrieving houses, adding new ones, updating existing ones, and deleting them.

Key Features and Functions:
1. Database Connection (PostgreSQL):
PostgreSQL Database: The API connects to a PostgreSQL database (mydb) using the pgx driver to manage houses data.
CRUD Operations on Houses:
Create: Add new houses to the database.
Read: Fetch houses from the database with filtering and pagination.
Update: Modify details of existing houses.
Delete: Remove houses from the database.
2. JWT-based Authentication:
Login Route (/login):
Users can login by providing a username and password. If valid, they receive a JWT token for subsequent requests.
JWT Middleware:
A middleware function is applied to protected routes to validate the JWT token passed in the Authorization header.
3. Role-Based Access Control (RBAC):
Role-based Rate Limiting:
Admin users have a higher rate limit for API requests compared to regular users.
Rate limits are dynamically set based on the role (admin or user).
4. Rate Limiting:
Rate Limiting Middleware:
Users and admins have different rate limits.
Rate limits are implemented using the ratelimit package, with different limits set for admins (200 requests/minute) and regular users (50 requests/minute).
5. Advanced Filtering & Pagination:
Filter by Price and Bedrooms:
Users can filter the list of houses by price range and number of bedrooms via query parameters.
Pagination:
Houses can be fetched in paginated form, with page and pageSize query parameters controlling the results.
6. Input Sanitization:
Sanitize Input:
Input data (e.g., house details) is sanitized before validation to prevent potential security risks such as XSS or SQL injection.
7. Caching with Redis:
Redis Caching:
Redis is used for caching the results of house queries to speed up frequent requests and reduce database load.
Cache invalidation and TTL (time-to-live) are managed for optimal performance.
8. Graceful Shutdown:
Graceful Shutdown:
The server gracefully handles shutdown signals, ensuring active requests are completed before the server stops.
9. Testing:
Unit Testing:
Unit tests can be added to validate the functionality of the API's key features (e.g., CRUD operations, JWT authentication).
End-to-End Testing:
Testing the entire API flow to ensure proper interactions between components.
API Endpoints:
Authentication:

POST /login: Log in to get a JWT token.
# Houses Management (Protected Routes):

GET /v1/houses: Fetch houses with optional filters (price range, bedrooms) and pagination.
POST /v1/houses: Add a new house (requires authentication).
PUT /v1/houses/:id: Update an existing house (requires authentication).
DELETE /v1/houses/:id: Delete a house (requires authentication).
Utility:

GET /test: A simple test route to check if the API is running.

# Technologies Used:
Gin: A web framework for building the API.
PostgreSQL: Relational database used for storing houses data.
JWT: JSON Web Tokens for authentication.
Redis: Caching layer to speed up frequently requested data.
Validator: For input validation and ensuring correct data format.
Rate Limiting: To prevent abuse of the API by limiting the number of requests from each user.
Graceful Shutdown: Properly handling server shutdowns to ensure ongoing requests are completed.

# Environment Variables:
DB_HOST: The address of the PostgreSQL database.
DB_USER: The username for the PostgreSQL database.
DB_PASSWORD: The password for the PostgreSQL database.
DB_NAME: The name of the PostgreSQL database.