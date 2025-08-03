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
- `meta`: Number of records and pagination details. The last three are included only if both `page` and `page_size` are provided.
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
      "displayed_records": 1,
      "current_page": 1,
      "records_per_page": 1,
      "total_pages": 77
    },
    "data": [
      {
        "province_id": 81,
        "thai_name": "กระบี่",
        "english_name": "Krabi"
      }
    ]
  }
  ```

### `/api/v1/districts`

- **Description:** Retrieves a list of districts in Thailand.
- **Additional Query Parameters:**

  - `province_id` _(optional)_: Filters districts by province.
  - `sort_by`: `0` (district_id), `1` (name_th), `2` (name_en)

- **Example Request:**

  ```http
  GET /api/v1/districts?province_id=81&page=1&page_size=2&sort_by=1
  ```

- **Example Response:**

  ```json
  {
    "meta": {
      "total_records": 928,
      "displayed_records": 2,
      "current_page": 1,
      "records_per_page": 2,
      "total_pages": 464
    },
    "data": [
      {
        "district_id": 8104,
        "thai_name": "คลองท่อม",
        "english_name": "Khlong Thom"
      },
      {
        "district_id": 8106,
        "thai_name": "ปลายพระยา",
        "english_name": "Plai Phraya"
      }
    ]
  }
  ```

### `/api/v1/subdistricts`

- **Description:** Retrieves a list of subdistricts in Thailand.
- **Additional Query Parameters:**

  - `district_id` _(optional)_: Filters subdistricts by district.
  - `sort_by`: `0` (subdistrict_id), `1` (name_th), `2` (name_en), `3` (postal_code)

- **Example Request:**

  ```http
  GET api/v1/subdistricts?district_id=8104&page=1&page_size=1&sort_by=1
  ```

- **Example Response:**

  ```json
  {
    "meta": {
      "total_records": 7436,
      "displayed_records": 1,
      "current_page": 1,
      "records_per_page": 1,
      "total_pages": 7436
    },
    "data": [
      {
        "subdistrict_id": 810402,
        "thai_name": "คลองท่อมเหนือ",
        "english_name": "Khlong Thom Nuea",
        "postal_code": 81120
      }
    ]
  }
  ```

### `/api/v1/subdistricts/:postal_code`

- **Description:** Retrieves subdistricts, districts, and provinces for a given postal code.
- **Additional Query Parameters:**

  - `sort_by`: `0` (subdistrict_id), `1` (name_th), `2` (name_en), `3` (postal_code)

- **Example Request:**

  ```http
  GET /api/v1/subdistricts/81120?page=1&page_size=2
  ```

- **Example Response:**

  ```json
  {
    "meta": {
      "total_records": 7436,
      "displayed_records": 2,
      "current_page": 1,
      "records_per_page": 2,
      "total_pages": 3718
    },
    "data": [
      {
        "subdistrict_id": 810303,
        "thai_name": "เกาะกลาง",
        "english_name": "Ko Klang",
        "postal_code": 81120,
        "district": {
          "district_id": 8103,
          "thai_name": "เกาะลันตา",
          "english_name": "Ko Lanta"
        },
        "province": {
          "province_id": 81,
          "thai_name": "กระบี่",
          "english_name": "Krabi"
        }
      },
      {
        "subdistrict_id": 810304,
        "thai_name": "คลองยาง",
        "english_name": "Khlong Yang",
        "postal_code": 81120,
        "district": {
          "district_id": 8103,
          "thai_name": "เกาะลันตา",
          "english_name": "Ko Lanta"
        },
        "province": {
          "province_id": 81,
          "thai_name": "กระบี่",
          "english_name": "Krabi"
        }
      }
    ]
  }
  ```

## Health Check

- **Endpoint:** `/`
- **Response:**

  ```json
  { "message": "Service is running" }
  ```
