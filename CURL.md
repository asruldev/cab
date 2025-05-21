My Curl
====

## Login
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"email":"user@example.com","password":"secret123"}'
```

## Register
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret123"
  }'
```

## Protected
```bash
curl -X GET http://localhost:8080/protected \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6"
```