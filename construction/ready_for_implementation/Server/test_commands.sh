curl -X GET localhost:8080/api/v1/users | jq -e   
curl -X POST -H 'Content-type: application/json' -d '{ "id": "1", "email": "test_user@gmail.com", "fullName": "John Doe" }'  localhost:8080/api/v1/users
curl -X POST -H 'Content-type: application/json' -d '{ "id": "2", "email": "other_one@gmail.com", "fullName": "Jane Lane" }'  localhost:8080/api/v1/users
curl -X DELETE -H 'Content-type: text/json' -d '{ "id": "2"}'  localhost:8080/api/v1/users