# 📘 Voucher API Documentation

## Techstack
 
 - **Gin**
 - **Gorm**
 - **Goose**
 - **Postgressql**
  
## 📌 Introduction

Voucher API adalah layanan yang memungkinkan pengguna untuk membuat merek (brand), membuat voucher, mendapatkan informasi voucher, dan melakukan penukaran (redemption).

## Setting .env
 1. Gunakan Setting env sesuai env example
 2. lakukan perintah ```goose -dir ./db/migrations postgres "host=localhost user=username password=password dbname=database_name sslmode=disable" up ```


## 🛠️ Base URL

```
http://your-api-url.com/api
```

---

## 🔹 Endpoints

### 🏷️ Create Brand

- **URL:** `/brand`
- **Method:** `POST`
- **Description:** Membuat merek baru.
- **Request Body:**
  ```json
  {
    "name": "Brand Name"
  }
  ```
- **Response:**
  ```json
  {
    "data": {
      "id": 2,
      "name": "Alfamart"
    },
    "message": "Brand created successfully",
    "statusCode": 201
  }
  ```

---

### 🎟️ Create Voucher

- **URL:** `/voucher`
- **Method:** `POST`
- **Description:** Membuat voucher baru untuk sebuah merek.
- **Request Body:**
  ```json
  {
    "brand_id": 1,
    "name": "Voucher Indomaret 20.000",
    "cost_in_point": 20000
  }
  ```
- **Response:**
  ```json
  {
    "data": {
      "id": 3,
      "brand_id": 1,
      "name": "Voucher Indomaret 30.000",
      "cost_in_point": 30000
    },
    "message": "Voucher created successfully",
    "statusCode": 201
  }
  ```

---

### 🔍 Get Single Voucher

- **URL:** `/voucher?id={voucher_id}`
- **Method:** `GET`
- **Description:** Mengambil informasi voucher berdasarkan ID.
- **Response:**
  ```json
  {
    "data": {
      "id": 3,
      "brand_id": 1,
      "name": "Voucher Indomaret 30.000",
      "cost_in_point": 30000
    },
    "message": "Voucher found",
    "statusCode": 200
  }
  ```

---

### 🏷️ Get All Vouchers by Brand

- **URL:** `/voucher/brand?id={brand_id}`
- **Method:** `GET`
- **Description:** Mengambil semua voucher berdasarkan ID merek.
- **Response:**
  ```json
  {
    "data": [
      {
        "id": 1,
        "brand_id": 1,
        "name": "Voucher Indomaret 10.000",
        "cost_in_point": 10000
      },
      {
        "id": 2,
        "brand_id": 1,
        "name": "Voucher Indomaret 20.000",
        "cost_in_point": 20000
      },
      {
        "id": 3,
        "brand_id": 1,
        "name": "Voucher Indomaret 30.000",
        "cost_in_point": 30000
      }
    ],
    "message": "Voucher found",
    "statusCode": 200
  }
  ```

---

### 🔄 Make Redemption

- **URL:** `/transaction/redemption`
- **Method:** `POST`
- **Description:** Melakukan penukaran voucher.
- **Request Body:**
  ```json
  {
    "voucher_id": 1,
    "quantity": 2,
    "total_point": 30000
  }
  ```
- **Response:**
  ```json
  {
    "data": {
      "id": 3,
      "voucher_id": 1,
      "quantity": 2,
      "total_point": 20000,
      "voucher": { "id": 0, "brand_id": 0, "name": "", "cost_in_point": 0 }
    },
    "message": "Transaction created successfully",
    "statusCode": 201
  }
  ```

---

## 🛡️ Authentication

Saat ini, API ini belum menggunakan autentikasi. Untuk keamanan, direkomendasikan untuk menggunakan **JWT** atau **API Key** di masa depan.

## 🚀 Best Practices

- Gunakan **header `Content-Type: application/json`** pada request yang membutuhkan body.
- Setiap response mengembalikan status HTTP yang sesuai (`201 Created`, `200 OK`, `400 Bad Request`, `500 Internal Server Error`).

---
