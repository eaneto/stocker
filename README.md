# stocker

![](https://github.com/eaneto/stocker/workflows/Stocker%20CI/badge.svg)
[![codecov](https://codecov.io/gh/eaneto/stocker/branch/main/graph/badge.svg)](https://codecov.io/gh/eaneto/stocker)

Bare bones.

## Endpoints

### Stocks resource

```json
GET /stocks
[
  {
    "ticker": "MST4",
    "price": 12.50
  },
  {
    "ticker": "MST3",
    "price": 8.50
  }
]
```

```json
GET /stocks/{ticker}
{
  "ticker": "MST4",
  "price": 12.50
}
```

```json
POST /stocks
{
  "ticker": "MST4",
  "price": 12.50
}
```

```json
PATCH /stocks
{
  "ticker": "MST4",
  "price": 12.50
}
```

```json
POST /customers
{
  "name": "Edison"
}
```

```json
GET /customers
[
  {
    "code": "240902c2-c4e0-4256-8430-60a21e0ab36c"
    "name": "Edison"
  }
]
```

```json
POST /order
{
  "customer": "240902c2-c4e0-4256-8430-60a21e0ab36c",
  "ticker": "PTR4",
  "amount": 10
}
```
