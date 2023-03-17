## API Reference
#### Create todo

```http
PUT /items/{id}
```

| Parameter | Type  | Description                |
|:----------|:------| :------------------------- |
| `Todo ID` | `int` | **Required** |

#### Get todo by ID

```http
GET /items/{id}
```

| Parameter | Type  | Description                |
|:----------|:------| :------------------------- |
| `Todo ID` | `int` | **Required** |

#### Get all todos

```http
GET /items
```

#### Update todo in database

```http
PATCH /items/{id}
```

| Parameter | Type  | Description                |
| :-------- |:------| :------------------------- |
| `StudentId` | `int` | **Required** |

#### Delete todo (Mark Status as "Delete")

```http
DELETE /items/{id}
```

| Parameter | Type  | Description                |
| :-------- |:------| :------------------------- |
| `StudentId` | `int` | **Required** |

#### Paging

```http
GET /items?page=&limit=
```

| Parameter | Type  | Description                |
|:----------|:------| :------------------------- |
| `Page`    | `int` | **Required** |
| `Limit`   | `int` | **Required** |

#### Filter

```http
GET /items?status=
```

| Parameter | Type     | Description                | 
|:----------|:---------| :------------------------- |
| `Status`  | `string` | **Required** |

`Value: "Done", "Doing", "Deleted"`

