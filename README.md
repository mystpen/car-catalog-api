# Car Catalog API
### Test task for junior golang developer position

- **Listing cars data with filtering and pagination:**
    ```http
    GET /cars
    ```
     ```http
    GET /cars?mark=toyota&page=2
    ```
    - queries for filtering:
        - regNum
        - mark
        - model
        - year
    - queries for pagination:
        - page
        - page_size
    - sample output:
    ```json

    ```
- **Adding new car data**
    ```http
    POST /cars
    ```
    - input body:
    ```json
    {
    "regNums": ["X123XX150"] // array
    }
    ```
- **Update car info:**
    - required parameter: `id`
     ```http
    PATCH /cars/:id
    ```
    - input body:
    ```json
    {
        "regNum": "X123XX156",
        "mark": "Ferrari",
        "model": "Tributo",
        "year": 2020,
        "owner": {
            "name": "John",
            "surname": "Adams",
            "patronymic": "Michael"
        }
    }
    ```
- **Delete car info:**
    - required parameter: `id`
     ```http
    DELETE /cars/:id
    ```
---
**- To run project:**
```
go run ./cmd/api 
```