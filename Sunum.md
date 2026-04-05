# Video Sunum

## Sunum Videosu

---

> **Video Linki:** [Sunum Videosu](https://www.youtube.com/watch?v=9Rlf6BbmDtk)

---

## Sunum Yapisi

### Konusmaci: Muhammet Ali Caliskan

---

## 1. Proje Tanitimi

- **Proje Adi:** LibraNet - Kutuphane Yonetim Sistemi
- **Amac:** Kutuphane islemlerini dijitallestirmek; kitap katalog yonetimi, kullanici kimlik dogrulama ve rezervasyon takibi
- **Canli Adres:** https://libranet-web.onrender.com
- **API Adresi:** https://libranet-api.onrender.com

### Kullanilan Teknolojiler

| Katman | Teknoloji |
|--------|-----------|
| Backend | Go 1.22+, Fiber v2, Bun ORM, PostgreSQL |
| Frontend | React 19, TypeScript, Vite, TailwindCSS v4 |
| State | Zustand |
| Auth | JWT (Bearer Token) |
| Deploy | Render (Web Service + Static Site + PostgreSQL) |

---

## 2. Frontend - Kullanici Arayuzu

### 2.1 Genel Kullanici Ozellikleri
- Kayit olma ve giris yapma
- Kitap arama ve listeleme (baslik, yazar, tur, mevcut durum)
- Kitap detay sayfasi
- Kitap rezervasyonu olusturma
- Kendi rezervasyonlarini goruntuleyebilme
- Kitap iade etme

---

## 3. Postman - API Test Gosterimi

**Base URL:** `https://libranet-api.onrender.com`

- Admin olusturma: `POST /api/auth/seed-admin`
- Kullanici kayit: `POST /api/auth/register`
- Kullanici giris: `POST /api/auth/login`
- Kitap ekleme: `POST /api/books` (Admin)
- Kitap guncelleme: `PUT /api/books/{bookid}` (Admin)
- Kitap silme: `DELETE /api/books/{bookid}` (Admin)
- Kitap listeleme/arama: `GET /api/books`
- Kitap detayi: `GET /api/books/{bookid}`
- Rezervasyon olusturma: `POST /api/reservations`
- Rezervasyon detayi: `GET /api/reservations/{reservationid}`
- Kullanici rezervasyonlari: `GET /api/users/{userid}/reservations`
- Kitap iade etme: `DELETE /api/reservations/{reservationid}`
- Raporlama: `GET /api/reports` (Admin)

---

## 4. Hata Yonetimi

Tum endpointler hata durumunda asagidaki formatta yanit doner:

```json
{
  "message": "Hata aciklamasi buraya gelir"
}
```

| Kod | Anlam | Ornek |
|-----|-------|-------|
| `400` | Gecersiz istek | Eksik alan, gecersiz tarih |
| `401` | Kimlik dogrulama hatasi | Token eksik veya suresi dolmus |
| `403` | Yetersiz yetki | User, admin islemine erismeye calisiyor |
| `404` | Kayit bulunamadi | Kitap veya rezervasyon yok |
| `409` | Cakisma | Ayni e-posta veya ayni kitaba tekrar rezervasyon |

---

## 5. Guvenlik Ozellikleri

- **JWT Tabanli Kimlik Dogrulama:** Her korunan istek `Authorization: Bearer <token>` gerektirir
- **Rol Tabanli Yetkilendirme:** `user` ve `admin` rolleri ile erisim kontrolu
- **Sifre Hashleme:** bcrypt ile sifre guvenligi
- **Hesap Kilitleme:** 5 basarisiz giris denemesinde 15 dakika kilitleme
- **CORS Korumasi:** Istemci tarafli erisim kontrolu
