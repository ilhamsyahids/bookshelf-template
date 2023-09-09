# Bookshelf API

## Get All Books

GET: `/books`

This endpoint is used to fetch all the latest books by date published. Books could be get by `query` based on book title.

**Query Parameter:**

- `query` (String) => Requested query. This is for querying books by book title
- `page` (Int) => Requested page. If not provided, the first page will be returned. Default value is `1` and at least `1`.
- `limit` (Int) => Number of books return. Default value is `10` and at least `1`.

**Example Request:**

```bash
GET /books?query=Harry%20Potter&page=2&limit=2
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": [
        {
            "id": 1,
            "isbn": "0653737570553",
            "title": "Harry Potter 1",
            "author": "J.K. Rowling",
            "published": "2002-10-30"
        },
        {
            "id": "1",
            "isbn": "4870077968367",
            "title": "Harry Potter et aliquid.",
            "author": "Jovany Schulist II",
            "published": "2003-11-15"
        }
    ]
}
```

**Error Response:**

- Page invalid:

    ```json
        {
            "ok": false,
            "message": "ERR_INVALID_PAGE"
        }
    ```

- Limit invalid:

    ```json
        {
            "ok": false,
            "message": "ERR_INVALID_LIMIT"
        }
    ```

## Get Book by ID

GET: `/books/:id`

This endpoint is used to fetch a book by its ID.

**Path Parameter:**

- `id` (Int) => Book ID

**Example Request:**

```bash
GET /books/1
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": {
        "id": 1,
        "isbn": "0653737570553",
        "title": "Harry Potter 1",
        "author": "J.K. Rowling",
        "published": "2002-10-30"
    }
}
```

**Error Response:**

- Book not found:

    ```json
        {
            "ok": false,
            "message": "ERR_BOOK_NOT_FOUND"
        }
    ```

## Create Book

POST: `/books`

This endpoint is used to create a new book.

**Request Body:**

- `isbn` (String, _REQUIRED_) => Book ISBN
- `title` (String, _REQUIRED_) => Book title
- `author` (String, _REQUIRED_) => Book author
- `published` (String, _REQUIRED_) => Book published date

**Example Request:**

```json
POST /books

{
    "isbn": "0653737570553",
    "title": "Harry Potter 1",
    "author": "J.K. Rowling",
    "published": "2002-10-30"
}
```

**Success Response:**

```json
HTTP/1.1 201 Created
Content-Type: application/json

{
    "ok": true,
    "data": {
        "id": 1,
        "isbn": "0653737570553",
        "title": "Harry Potter 1",
        "author": "J.K. Rowling",
        "published": "2002-10-30"
    }
}
```

**Error Response:**

- ISBN missing:

    ```json
        {
            "ok": false,
            "message": "ERR_ISBN_MISSING"
        }
    ```

- Title missing:

    ```json
        {
            "ok": false,
            "message": "ERR_TITLE_MISSING"
        }
    ```

- Author missing:

    ```json
        {
            "ok": false,
            "message": "ERR_AUTHOR_MISSING"
        }
    ```

- Date published missing:

    ```json
        {
            "ok": false,
            "message": "ERR_PUBLISHED_MISSING"
        }
    ```

## Update Book

PUT: `/books/:id`

This endpoint is used to update a book by its ID.

**Path Parameter:**

- `id` (Int) => Book ID

**Request Body:**

- `isbn` (String, _OPTIONAL_) => Book ISBN
- `title` (String, _OPTIONAL_) => Book title
- `author` (String, _OPTIONAL_) => Book author
- `published` (String, _OPTIONAL_) => Book published date

**Example Request:**

```json
PUT /books/1

{
    "title": "Harry Potter 3"
}
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": {
        "id": 1,
        "isbn": "0653737570553",
        "title": "Harry Potter 3",
        "author": "J.K. Rowling",
        "published": "2002-10-30"
    }
}
```

**Error Response:**

- Book not found:

    ```json
        {
            "ok": false,
            "message": "ERR_BOOK_NOT_FOUND"
        }
    ```

## Delete Book

DELETE: `/books/:id`

This endpoint is used to delete a book by its ID.

**Path Parameter:**

- `id` (Int) => Book ID

**Example Request:**

```bash
DELETE /books/1
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true
}
```

**Error Response:**

- Book not found:

    ```json
        {
            "ok": false,
            "message": "ERR_BOOK_NOT_FOUND"
        }
    ```
