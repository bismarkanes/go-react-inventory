# Inventory service can be modeled around three core tables:

## Table Item
```
|----------------+----------+--------------------+---------|
| field          | type     | description        | example |
|----------------+----------+--------------------+---------|
| id             | auto int | primary key of id  |       1 |
| stock          | integer  | current item stock |     100 |
| reserved_stock | integer  | reserverd stock    |       9 |
|----------------+----------+--------------------+---------|
```

## Table ReservationItem
```
|------------+----------+-------------------------------------------+--------------------------|
| field      | type     | description                               |                  example |
|------------+----------+-------------------------------------------+--------------------------|
| id         | auto int | primary key id                            |                        2 |
| user_id    | int      | foregitn key to table User                |                       10 |
| item_id    | int      | foreign key to table Item                 |                        2 |
| quantity   | int      | reserved quantity to stock                |                        2 |
| expires_at | datetime |                                           | 2026-10-08T14:01:51+0700 |
| confirmed  | boolean  | reservation has been confirmed            |                     true |
| confirm_at | datetime | time which reservation has been confirmed | 2026-10-08T14:01:51+0700 |
|------------+----------+-------------------------------------------+--------------------------|
```

## Table User
```
|-------+----------+----------------+---------|
| field | type     | description    | example |
|-------+----------+----------------+---------|
| id    | auto int | primary key id | 1       |
| name  | varchar  | name of user   | martin  |
|-------+----------+----------------+---------|
```

# Reservation flow

# Inventory service routing:

```
:HOST=http://localhost:8080/api/v1
```

## Create Reservation
**Sample request**:
```
POST :HOST/inventory/reserve
content-type: application/json

{
  "user_id": "usr_9981",
  "item_id": "1",
  "quantity": 90
}
```

**Sample JSON response**:
```
{
  "status": "success",
  "reservation_id": "3",
  "item_id": "1",
  "user_id": "usr_9981",
  "quantity": 90,
  "expires_at": "2026-10-08T14:14:00.211160844+07:00"
}
```

## Confirm Reservation
**Sample request**:
```
POST :HOST/inventory/confirm
content-type: application/json

{
	"reservation_id": "2"
}
```

**Sample response**:
```
{
  "status": "success",
  "reservation_id": "3",
  "confirmed_at": "2026-10-08T14:09:28.13025949+07:00"
}
```

## Get inventory status
**Sample request**:
```
GET :HOST/inventory/stock?item_id=1
```

**Sample response**:
```
{
  "item_id": "1",
  "total_stock": 100,
  "reserved_stock": 90,
  "available_stock": 10
}
```
