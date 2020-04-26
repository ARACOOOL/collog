# collog

[![Build Status](https://travis-ci.org/ARACOOOL/collog.svg?branch=master)](https://travis-ci.org/ARACOOOL/collog)

Collog is a standalone server that provides a REST API to collect your logs.

# Launch

```bash
server.exe -h=":8080" -jk="qwerty" -dsn="root:qwerty@tcp(192.168.99.100)/logs"
```

## Parameters
```
h   - server host
jk  - JWT secret key
ld  - Logs directory (for the server logs)
dsn - Mysql data source name
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

## Filter parameters

Filter by a name of source
`/logs?source={sourceName}`

Filter by a name of category
`/logs?category={categoryName}`


Filter by an error level
`/logs?level=400`

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
  "email": "test@test.com",
  "iat": 1516239022
}
```

[JWT documentation](https://jwt.io/)