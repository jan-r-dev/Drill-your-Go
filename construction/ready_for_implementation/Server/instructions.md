1. Using Go's standard library, create a simple localhost webserver on a port of your choice.
2. Create a simple struct that will hold a list of users. Implementation details are up to you.
* For variety you can also store them in a file instead of in-memory.
3. Create a router for that webserver that will conduct appropriate actions and return HTTP codes for:
* POST to `/api/v1/users` to add a user to the existing list. Body of the request should match your implementation.
* GET to `api/v1/users` to retrieve a current list of users (can be simple csv form, up to you)
* DELETE to `api/v1/users` to delete an existing user using a primary key that matches your implementation.
4. Modify the server by configuring read and write timeouts.