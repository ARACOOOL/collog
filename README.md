# collog

# Launch

```bash
server.exe -h=":8080" -jk="qwerty" -dsn="root:qwerty@tcp(192.168.99.100)/logs"
```

# Usage

## Store a log

`POST /logs`

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

`GET /logs`

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