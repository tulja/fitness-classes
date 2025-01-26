# fitness-classes

## Setup Instructions 

```
git clone https://github.com/tulja/fitness-classes 
cd fitness-classes
go mod tidy 
go run .
```

### Sample Output 

```
PS D:\fitness-classes> go run .
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /classes                  --> main.GetGymClasses (4 handlers)
[GIN-debug] GET    /classes/:id              --> main.GetGymClassByID (4 handlers)
[GIN-debug] GET    /classes/info/:class_name --> main.GetGymClassByClassName (4 handlers)
[GIN-debug] POST   /classes                  --> main.CreateGymClass (4 handlers)
[GIN-debug] DELETE /classes/:id              --> main.DeleteGymClassByID (4 handlers)
[GIN-debug] GET    /bookings                 --> main.GetBookings (4 handlers)
[GIN-debug] GET    /bookings/:id             --> main.GetBookingByID (4 handlers)
[GIN-debug] GET    /bookings/info/:booking_name --> main.GetBookingByBookingName (4 handlers)
[GIN-debug] POST   /bookings                 --> main.CreateBooking (4 handlers)
[GIN-debug] DELETE /bookings/:id             --> main.DeleteBookingByID (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on localhost:9090
[GIN] 2025/01/26 - 17:18:40 | 200 |            0s |       127.0.0.1 | GET      "/bookings"
[GIN] 2025/01/26 - 17:18:46 | 200 |            0s |       127.0.0.1 | GET      "/classes"
[GIN] 2025/01/26 - 17:19:01 | 200 |            0s |       127.0.0.1 | GET      "/classes/info/Yoga"
[GIN] 2025/01/26 - 17:19:11 | 200 |            0s |       127.0.0.1 | GET      "/classes/info/Pilates"
[GIN] 2025/01/26 - 17:19:26 | 400 |            0s |       127.0.0.1 | POST     "/bookings"
[GIN] 2025/01/26 - 17:19:46 | 201 |            0s |       127.0.0.1 | POST     "/bookings"
[GIN] 2025/01/26 - 17:21:25 | 200 |            0s |       127.0.0.1 | GET      "/bookings"
[GIN] 2025/01/26 - 17:21:35 | 200 |            0s |       127.0.0.1 | GET      "/bookings/info/tbd12"
[GIN] 2025/01/26 - 17:22:55 | 200 |       549.5µs |       127.0.0.1 | DELETE   "/bookings/1"
[GIN] 2025/01/26 - 17:23:09 | 200 |            0s |       127.0.0.1 | GET      "/bookings"
```


## Usage

### Create new class
```
curl localhost:9090/classes \
--header 'Content-Type: application/json' \
--data '{"class_name": "Zumba", "capacity":50,"start_date":"01-04-2025", "end_date":"15-04-2025"} '
```

Output 
```
{
    "id": "3",
    "class_name": "Zumba",
    "start_date": "01-03-2025",
    "end_date": "15-03-2025",
    "capacity": 50
}
```

### List all classes 
```
curl localhost:9090/classes
```
Output
```
C:\Users\DELL>curl localhost:9090/classes
[
    {
        "id": "1",
        "class_name": "Pilates",
        "start_date": "25-01-2025",
        "end_date": "31-01-2025",
        "capacity": 56
    },
    {
        "id": "2",
        "class_name": "Yoga",
        "start_date": "01-02-2025",
        "end_date": "05-02-2025",
        "capacity": 17
    },
    {
        "id": "3",
        "class_name": "Zumba",
        "start_date": "01-03-2025",
        "end_date": "15-03-2025",
        "capacity": 50
    }
]
```

### Get class by Id 
Below example uses 1 as id
```
 curl localhost:9090/classes/1
```

Output
```
{
    "id": "1",
    "class_name": "Pilates",
    "start_date": "25-01-2025",
    "end_date": "31-01-2025",
    "capacity": 56
}
```

### Get class by class_name
Below example uses "Yoga" as class_name
```
curl localhost:9090/classes/info/Yoga
```
Output
```
{
    "id": "2",
    "class_name": "Yoga",
    "start_date": "01-02-2025",
    "end_date": "05-02-2025",
    "capacity": 17
}
```

### Delete Class

Requires id of class. Below example uses 1 as id
```
curl --request DELETE localhost:9090/classes/1
```

Output
```
"Deleted id 1 Successfully!"
```



### Create new booking 
```
curl localhost:9090/bookings \
--header 'Content-Type: application/json' \
--data '{"booking_name":"abc","class_name":"Yoga","date":"02-02-2025"}'
```

Output 
```
{
    "id": "1",
    "class_name": "Yoga",
    "date": "02-02-2025",
    "booking_name": "abc"
}
```

### List all bookings
```
curl localhost:9090/bookings
```
Output
```
[
    {
        "id": "2",
        "class_name": "Yoga",
        "date": "02-02-2025",
        "booking_name": "tbd12"
    },
    {
        "id": "1",
        "class_name": "Yoga",
        "date": "02-02-2025",
        "booking_name": "abc"
    }
]
```

### Get booking by Id 
Below example uses 1 as id
```
 curl localhost:9090/bookings/1
```

Output
```
{
    "id": "1",
    "class_name": "Yoga",
    "booking_name": "abc",
    "date": "02-02-2025",
}
```

### Get booking by booking_name
Below example uses "xyz" as booking_name
```
curl localhost:9090/bookings/info/xyz
```
Output
```
{
    "id": "2",
    "class_name": "Yoga",
    "booking_name": "xyz",
    "date": "02-02-2025",
}
```

### Delete Booking

Requires id of booking. Below example uses 1 as id
```
curl --request DELETE localhost:9090/bookings/1
```

Output
```
"Deleted booking id 1 Successfully!"
```

## Testing

### Tools
Postman, cURL

