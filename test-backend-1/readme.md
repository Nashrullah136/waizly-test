# User Management API

- This API can add, update, read, dan delete user
- You can see example API in request_example.postman_collection
- This API have migration and data seed for deployment, you can activate migrate and data seeder by set `MIGRATE` and `SEED` to true in .env (you can look at .env.example)
- You can set default user and password for initial admin in .env
- This API only support for MySql Database
- This API have authentication and authorization feature. Using JWT token to authentication