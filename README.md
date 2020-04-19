# collog

# Launch

```bash
server.exe -h=":8080" -jk="qwerty" -dsn="root:qwerty@tcp(192.168.99.100)/logs"
```

# Usage

## Store a log

```
POST /logs

Authorization: Bearer {token}
```

```json
{
	"source": "web app",
	"category": "auth",
	"level": 400,
	"message": "Invalid credentials",
	"trace": "{}",
	"payload": {"login": "user22"}
}
```

## Get list of logs

```
GET /logs

Authorization: Bearer {token}
```

**Response**

```json
[
    {
        "id": "19440916-9e24-4ce0-9770-f38eba0d6667",
        "source": "web app",
        "category": "auth",
        "level": 400,
        "message": "Invalid credentials",
        "trace": "{}",
        "payload": {
            "login": "user22"
        },
        "created_at": "2020-04-19 16:06:36"
    }
]
```

# JWT Token

**Headers**

```json
{
  "alg": "HS512",
  "typ": "JWT"
}
```

**Payload**

```json
{
  "name": "test@test.com",
  "iat": 1516239022
}
```