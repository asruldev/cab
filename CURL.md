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
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMyIsImVtYWlsIjoidXNlckBleGFtcGxlLmNvbSIsImlzcyI6ImRldi5hc3J1bC5jYmUiLCJleHAiOjE3NDc4MzcyMjMsImlhdCI6MTc0NzgzNjMyM30.okooF3h2YxLUettqLppmF30eRckPg74oX2E-wTOr9JU"
```