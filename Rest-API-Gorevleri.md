# REST API Metotları

**Proje:** LibraNet – Kütüphane Yönetim Sistemi
**Geliştirici:** Muhammet Ali Çalışkan
**Rol:** Backend REST API

---

## 1. Kullanıcı Kaydı

- **Endpoint:** `POST /api/auth/register`
- **Request Body:**
  ```json
  {
    "name": "Muhammet Ali Çalışkan",
    "email": "ali@example.com",
    "password": "gizli1234"
  }
  ```
- **Response:** `201 Created` – Kullanıcı başarıyla oluşturuldu
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "_id": "usr123",
      "name": "Muhammet Ali Çalışkan",
      "email": "ali@example.com",
      "role": "user"
    }
  }
  ```

---

## 2. Kullanıcı Girişi

- **Endpoint:** `POST /api/auth/login`
- **Request Body:**
  ```json
  {
    "email": "ali@example.com",
    "password": "gizli1234"
  }
  ```
- **Response:** `200 OK` – Giriş başarılı, JWT token döndürüldü
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "_id": "usr123",
      "name": "Muhammet Ali Çalışkan",
      "email": "ali@example.com",
      "role": "user"
    }
  }
  ```

---

## 3. Kitap Ekleme (Admin)

- **Endpoint:** `POST /api/books`
- **Authentication:** Bearer Token gerekli — **Admin rolü zorunlu**
- **Request Body:**
  ```json
  {
    "title": "Dune",
    "author": "Frank Herbert",
    "isbn": "978-0-441-01359-7",
    "genre": "Bilim Kurgu",
    "publishedYear": 1965,
    "totalCopies": 3
  }
  ```
- **Response:** `201 Created` – Kitap başarıyla eklendi
  ```json
  {
    "_id": "book789",
    "title": "Dune",
    "author": "Frank Herbert",
    "isbn": "978-0-441-01359-7",
    "genre": "Bilim Kurgu",
    "publishedYear": 1965,
    "available": true,
    "totalCopies": 3,
    "availableCopies": 3
  }
  ```

---

## 4. Kitap Güncelleme (Admin)

- **Endpoint:** `PUT /api/books/{bookid}`
- **Path Parameters:**
  - `bookid` (string, required) – Güncellenecek kitabın ID'si
- **Authentication:** Bearer Token gerekli — **Admin rolü zorunlu**
- **Request Body:**
  ```json
  {
    "title": "Dune (Yeni Baskı)",
    "totalCopies": 5
  }
  ```
- **Response:** `200 OK` – Kitap başarıyla güncellendi

---

## 5. Kitap Silme (Admin)

- **Endpoint:** `DELETE /api/books/{bookid}`
- **Path Parameters:**
  - `bookid` (string, required) – Silinecek kitabın ID'si
- **Authentication:** Bearer Token gerekli — **Admin rolü zorunlu**
- **Response:** `204 No Content` – Kitap başarıyla silindi

---

## 6. Kitap Detayı Görüntüleme

- **Endpoint:** `GET /api/books/{bookid}`
- **Path Parameters:**
  - `bookid` (string, required) – Görüntülenecek kitabın ID'si
- **Authentication:** Bearer Token gerekli
- **Response:** `200 OK` – Kitap detayları başarıyla getirildi
  ```json
  {
    "_id": "book789",
    "title": "Dune",
    "author": "Frank Herbert",
    "isbn": "978-0-441-01359-7",
    "genre": "Bilim Kurgu",
    "publishedYear": 1965,
    "available": true,
    "totalCopies": 3,
    "availableCopies": 2
  }
  ```

---

## 7. Kitap Arama (Kütüphaneler Arası)

- **Endpoint:** `GET /api/books`
- **Authentication:** Bearer Token gerekli
- **Query Parameters:**
  - `query` (string, optional) – Kitap adı, yazar veya ISBN
  - `genre` (string, optional) – Tür filtresi
  - `available` (boolean, optional) – Yalnızca müsait kitapları getir
  - `page` (integer, optional) – Sayfa numarası (varsayılan: 1)
  - `limit` (integer, optional) – Sayfa başına sonuç (varsayılan: 10)
- **Response:** `200 OK` – Kitap listesi başarıyla getirildi
  ```json
  [
    {
      "_id": "book789",
      "title": "Dune",
      "author": "Frank Herbert",
      "available": true
    }
  ]
  ```

---

## 8. Rezervasyon Oluşturma

- **Endpoint:** `POST /api/reservations`
- **Authentication:** Bearer Token gerekli
- **Request Body:**
  ```json
  {
    "bookId": "book789",
    "dueDate": "2026-03-31"
  }
  ```
- **Response:** `201 Created` – Rezervasyon başarıyla oluşturuldu
  ```json
  {
    "_id": "rsv001",
    "userId": "usr123",
    "bookId": "book789",
    "status": "active",
    "reservedAt": "2026-03-09T10:00:00Z",
    "dueDate": "2026-03-31T10:00:00Z",
    "returnedAt": null
  }
  ```

---

## 9. Kitap İadesi

- **Endpoint:** `DELETE /api/reservations/{reservationid}`
- **Path Parameters:**
  - `reservationid` (string, required) – İade edilecek rezervasyonun ID'si
- **Authentication:** Bearer Token gerekli
- **Response:** `204 No Content` – Kitap başarıyla iade edildi, rezervasyon kapatıldı

---

## 10. Rezervasyon Detayı Görüntüleme

- **Endpoint:** `GET /api/reservations/{reservationid}`
- **Path Parameters:**
  - `reservationid` (string, required) – Görüntülenecek rezervasyonun ID'si
- **Authentication:** Bearer Token gerekli
- **Response:** `200 OK` – Rezervasyon detayları başarıyla getirildi
  ```json
  {
    "_id": "rsv001",
    "userId": "usr123",
    "bookId": "book789",
    "status": "active",
    "reservedAt": "2026-03-09T10:00:00Z",
    "dueDate": "2026-03-31T10:00:00Z",
    "returnedAt": null
  }
  ```

---

## 11. Kullanıcı Rezervasyon Listesi

- **Endpoint:** `GET /api/users/{userid}/reservations`
- **Path Parameters:**
  - `userid` (string, required) – Rezervasyonları listelenecek kullanıcının ID'si
- **Authentication:** Bearer Token gerekli
- **Query Parameters:**
  - `status` (string, optional) – `active`, `returned`, `all` (varsayılan: `all`)
  - `page` (integer, optional) – Sayfa numarası (varsayılan: 1)
  - `limit` (integer, optional) – Sayfa başına sonuç (varsayılan: 10)
- **Response:** `200 OK` – Rezervasyon listesi başarıyla getirildi

---

## 12. Raporlama & İstatistik (Admin)

- **Endpoint:** `GET /api/reports`
- **Authentication:** Bearer Token gerekli — **Admin rolü zorunlu**
- **Query Parameters:**
  - `type` (string, optional) – `books`, `reservations`, `users`
  - `from` (date, optional) – Başlangıç tarihi (örn. `2026-01-01`)
  - `to` (date, optional) – Bitiş tarihi (örn. `2026-03-09`)
- **Response:** `200 OK` – Rapor başarıyla getirildi
  ```json
  {
    "generatedAt": "2026-03-09T12:00:00Z",
    "type": "reservations",
    "summary": {
      "total": 450,
      "active": 120,
      "returned": 330
    },
    "topBooks": [
      { "bookId": "book789", "title": "Dune", "reservationCount": 42 }
    ]
  }
  ```
