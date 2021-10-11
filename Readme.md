

# Filter Rectangles!!
This service will filter all input rectangle based on given rectangle then store them into database for further use.


## Run

For run this service simply run the command below:

```bash
docker-compose up -d --build
```

## API'S

This service has 2 endpoint:

- POST: http://127.0.0.1:18080/

You can send your main rectangle and input rectangles to this endpoint and what it will do is 
that filter all input rectangles which is overlap with main rectangle then it will store them to database.

on successful return response with status code 201.

- GET: http://127.0.0.1:18080/

This endpoint will return all rectangles that stored in database so far.

on successful return response with status code 200.

## Support other databases

This service can easily support other databases. What you should do is that
implement the interface below and launch it through main function. Actually this
service make access to databases through Repository design pattern.

```go
//RectangleRepo should be implemented for all databases.
type RectangleRepo interface {

	//Query api's
	GetRectByID(ctx context.Context, Id string) *rect_model.Rectangle
	GetAllRect(ctx context.Context) []rect_model.Rectangle

	//modify api's
	Migrate(ctx context.Context)
	SaveRectangle(ctx context.Context, rect *rect_model.Rectangle) error
	UpdateRectangle(ctx context.Context, rect *rect_model.Rectangle) error
	DeleteRectangleByID(ctx context.Context, Id string) error
	DeleteAllRectangle(tx context.Context)
}
```

## Examples

### Send main rectangle and input rectangles

```bash
curl -X POST http://127.0.0.1:18080/ -d '{
    "main": {"x": 3, "y": 2, "width": 5, "height": 10},
    "inputs": [
        {"x": 4, "y": 10, "width": 1, "height": 1},
        {"x": 9, "y": 10, "width": 5, "height": 4}
    ]
}'  
```

### Get filtered and valid rectangles

```bash
curl http://127.0.0.1:18080/
```

Return:

```json
[{"x":4,"y":10,"width":1,"height":1,"time":"Mon, 11 Oct 2021 08:56:20 +0000"}]
```