# MarketPlace

This is a simple `REST APT` server.

## Getting started

To run this server, you need to create `.env` file in the root directory with following data:

```bash
PG_USER=postgres_user
PG_DBNAME=postgres_db_name
PG_PASSWORD=postgres_user_password
PG_PORT=5432
SECRET_KEY=secret_key_to_generate_jwt
```

Also you need to install few cli utils. To make it easier I implemented single `make` command. Simply run:

```bash
make beautiful version=0.0.1
```
where `version` is a image version of an app.

## End points

### Authorization

- `POST /register` - register with your `username` and `password`
- `POST /login` - login with your `username` and `password`

### Ads

- `POST /ads` - post a new ad. In body must contain json with fields
  - must supply with header like `Authorization: Bearer <your_access_token_here>`

```json
{
    "name": "ad name",
    "description": "ad description",
    "price": "10.05",
    "pictureUrl": "picture url"
}
```

- `GET /ads` - get a ads list. Allowed query params:
  - `page_num` - current page, dafaults to `1`
  - `results_per_page` - number of ads per page, defaults to `10`
  - `sort_by` - field to sort by, allowed values:
    - `"date"` - ad's creation date, default
    - `"price"` - ad's price
  - `order` - sorting order:
    - `"desc"` - descending, default
    - `"asc"` - ascending
  - `"max_price"` - max price, defaults to `0`
  - `"min_price"` - min price, defaults to unlimited
