# API Reference and Usage Guide

## Common Query Parameters

The following query parameters are supported by all endpoints:

- `page`: Page number for pagination.
- `page_size`: Number of items per page. Defaults to all items if not specified.
- `search`: Keyword to search within Thai or English names of address components.
- `sort_order`: Sort order. Use `0` for ascending (default) or `1` for descending.

**Note:** All endpoints support `sort_by`, but valid values depend on the specific endpoint.

## Response Structure

All responses are returned in JSON format:

```json
{
  "meta": {
    "total_records": 100,
    "display_records": 10,
    "current_page": 1,
    "records_per_page": 10,
    "total_pages": 10
  },
  "data": {},
  "message": "Normal or error message"
}
```

- `data`: Contains the result set. Structure varies by endpoint.
- `meta`: Pagination details. Included only if both `page` and `page_size` are provided.
- `message`: Status message.
- **Error responses**: Contain only the `message` field; `meta` and `data` are omitted.

## Endpoints

### `/api/v1/provinces`

- **Description:** Retrieves a list of provinces in Thailand.
- **Additional Query Parameters:**

  - `sort_by`: `0` (province_id), `1` (name_th), `2` (name_en)

- **Example Request:**

  ```http
  GET /api/v1/provinces?page=1&page_size=2&sort_by=1
  ```

- **Example Response:**

  ```json
  {
    "meta": {
      "total_records": 77,
      "display_records": 10,
      "current_page": 1,
      "records_per_page": 10,
      "total_pages": 8
    },
    "data": [
      {
        "province_id": 10,
        "name_th": "กรุงเทพมหานคร",
        "name_en": "Bangkok"
      },
      ...
    ],
  }
  ```

### `/api/v1/districts`

- **Description:** Retrieves a list of districts in Thailand.
- **Additional Query Parameters:**

  - `province_id` _(optional)_: Filters districts by province.
  - `sort_by`: `0` (district_id), `1` (name_th), `2` (name_en)

- **Example Request:**

  ```http
  GET /api/v1/districts?province_id=1&page=1&page_size=10&sort_by=1
  ```

- **Example Response:**

  ```json
  {
    "meta": {
      "total_records": 928,
      "display_records": 10,
      "current_page": 1,
      "records_per_page": 10,
      "total_pages": 93
    },
    "data": [
      {
        "district_id": 1001,
        "name_th": "พระนคร",
        "name_en": "Phra Nakhon",
      },
      ...
    ],
  }
  ```

### `/api/v1/subdistricts`

- **Description:** Retrieves a list of subdistricts in Thailand.
- **Additional Query Parameters:**

  - `district_id` _(optional)_: Filters subdistricts by district.
  - `sort_by`: `0` (subdistrict_id), `1` (name_th), `2` (name_en), `3` (postal_code)

- **Example Request:**

  ```http
  GET /api/v1/subdistricts?district_id=1&page=1&page_size=10&sort_by=1
  ```

- **Example Response:**

  ```json
  {
    "meta": {
      "total_records": 7436,
      "display_records": 10,
      "current_page": 1,
      "records_per_page": 10,
      "total_pages": 744
    },
    "data": [
      {
        "subdistrict_id": 100101,
        "name_th": "พระบรมมหาราชวัง",
        "name_en": "Phra Borom Maha Ratchawang",
        "postal_code": "10200"
      },
      ...
    ],
  }
  ```

### `/api/v1/subdistricts/:postal_code`

- **Description:** Retrieves subdistricts, districts, and provinces for a given postal code.
- **Additional Query Parameters:**

  - `sort_by`: `0` (subdistrict_id), `1` (name_th), `2` (name_en), `3` (postal_code)

- **Example Request:**

  ```http
  GET /api/v1/subdistricts/10200?page=1&page_size=10&sort_by=1
  ```

- **Example Response:**

  ```json
  {
    "meta": {
      "total_records": 7436,
      "displayed_records": 10,
      "current_page": 1,
      "records_per_page": 10,
      "total_pages": 744
    },
    "data": [
      {
        "subdistrict_id": 100109,
        "thai_name": "ชนะสงคราม",
        "english_name": "Chana Songkhram",
        "postal_code": 10200,
        "district": {
          "district_id": 1001,
          "thai_name": "พระนคร",
          "english_name": "Phra Nakhon"
        },
        "province": {
          "province_id": 10,
          "thai_name": "กรุงเทพมหานคร",
          "english_name": "Bangkok"
        }
      },
      ...
    ],
  }
  ```

## Health Check

- **Endpoint:** `/`
- **Response:**

  ```json
  { "message": "Service is running" }
  ```
