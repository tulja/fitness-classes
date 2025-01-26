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
