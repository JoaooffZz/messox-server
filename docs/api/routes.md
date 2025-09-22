# 📌 Routes API

## 👤 Routes-User:

```diff
(POST) /user/login
```
```bash
curl -L \
  -H "Accept: application/json" \
  -H "Authorization: Bearer <NONE-TOKEN>" \
  https://api:8080/user/login
```
+ request-body -J \
+ {
+   "name": "string",
+   "password": "string"
+ }
+ response-body -J \
+ {
+   "token": "string",
+ }
```plaintext
   status-codes:
        (200) -> "sucess login!"
        (400) -> "headers or reponse-body not is valid"
        (401) -> "not authorized"
        (404) -> "user not found"
        (500) -> "internal error"
```

```diff
(POST) /user/register
```
```bash
curl -L \
  -H "Accept: application/json" \
  -H "Authorization: Bearer <NONE-TOKEN>" \
  https://api:8080/user/register
```
+ request-body -J \
+ {
+   "name": "string",
+   "password": "string"
+ }
+ response-body -J \
+ {
+   "token": "string",
+   "name": "string",
+   "profile": "bytes",
+   "bio": "string"
+ }
```plaintext
   status-codes:
        (201) -> "created user"
        (400) -> "headers or request-body invalid"
        (401) -> "not authorized"
        (409) -> "user already exists (unique violation)"
        (413) -> "string length exceeded"
        (404) -> "resource not found"
        (500) -> "internal error"
```

## 💻 Routes-Server:

```diff
(GET) /server/ping
```
```bash
curl -L \
  -H "Accept: */*" \
  -H "Authorization: Bearer <YOUR-API-KEY?>" \
  https://api:8080/server/ping
```
```plaintext
   status-codes:
        (200) -> "sucess ping"
        (400) -> "headers or request-body invalid"
        (401) -> "not authorized, api key is necessary"
```

```diff
(GET) /server/ws
```
```bash
curl -L \
  -H "Accept: */*" \
  -H "Authorization: Bearer <YOUR-TOKEN>" \
  https://api:8080/server/ws
```
```plaintext
   status-codes:
        (101) -> "Switching Protocols, conection WebSocket sucess!"
        (200) -> "success ping"
        (400) -> "headers invalid or missing"
        (401) -> "not authorized, api key is necessary"
        (403) -> "forbidden, token invalid or expired"
        (426) -> "Upgrade Required, WebSocket connection refused"
        (500) -> "internal server error"
```