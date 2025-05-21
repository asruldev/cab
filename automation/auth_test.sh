#!/bin/bash
# jika belum ada brew install jq
# lakukan: chmod +x auth_test.sh
# Login dan dapatkan token
response=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}')

token=$(echo "$response" | jq -r '.token')

if [ "$token" == "null" ] || [ -z "$token" ]; then
  echo "Login gagal, token tidak ditemukan"
  exit 1
fi

echo "Token didapat: $token"

# Akses endpoint protected dengan token
curl -X GET http://localhost:8080/protected \
  -H "Authorization: Bearer $token"

echo
