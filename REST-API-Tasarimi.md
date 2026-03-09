# REST API Tasarımı

**Proje:** LibraNet – Kütüphane Yönetim Sistemi
**Standart:** OpenAPI 3.0.3
**Kimlik Doğrulama:** JWT (Bearer Token)

---

## API Spesifikasyonu

LibraNet REST API'sinin tam tasarımı OpenAPI 3.0.3 standardında hazırlanmıştır.

> **[libranet.yaml dosyasını görüntüle](API-Tasarımı.md)** — Tüm endpoint tanımları, şemalar ve örnek yanıtlar bu dosyada yer almaktadır.

---

## Temel Tasarım Kararları

### Kimlik Doğrulama
- JWT tabanlı `Bearer Token` kullanılmaktadır.
- Her korunan isteğe `Authorization: Bearer <token>` başlığı eklenmeli.
- `/api/auth/register` ve `/api/auth/login` endpointleri herkese açıktır (auth gerektirmez).

### Rol Tabanlı Yetkilendirme
| Rol | Yetkiler |
|-----|---------|
| **user** | Kayıt, giriş, kitap arama/görüntüleme, rezervasyon oluşturma/iade |
| **admin** | User yetkileri + kitap ekleme/güncelleme/silme + raporlama |

### HTTP Metotları
| Metot | Kullanım |
|-------|---------|
| `GET` | Veri okuma |
| `POST` | Yeni kayıt oluşturma |
| `PUT` | Mevcut kaydı güncelleme |
| `DELETE` | Kayıt silme / işlem kapatma |

### HTTP Durum Kodları
| Kod | Anlam |
|-----|-------|
| `200` | Başarılı |
| `201` | Kayıt oluşturuldu |
| `204` | Başarılı, içerik yok (silme/iade) |
| `400` | Geçersiz istek |
| `401` | Kimlik doğrulama hatası |
| `403` | Yetersiz yetki |
| `404` | Kayıt bulunamadı |
| `409` | Çakışma (örn. aynı e-posta) |

---

## Endpoint Özeti

### Kimlik Doğrulama
| Metot | Endpoint | Açıklama | Yetki |
|-------|----------|----------|-------|
| `POST` | `/api/auth/register` | Kullanıcı kaydı | Herkese açık |
| `POST` | `/api/auth/login` | Kullanıcı girişi | Herkese açık |

### Kitaplar
| Metot | Endpoint | Açıklama | Yetki |
|-------|----------|----------|-------|
| `POST` | `/api/books` | Kitap ekle | Admin |
| `GET` | `/api/books` | Kitap ara / listele | User, Admin |
| `GET` | `/api/books/{bookid}` | Kitap detayını görüntüle | User, Admin |
| `PUT` | `/api/books/{bookid}` | Kitabı güncelle | Admin |
| `DELETE` | `/api/books/{bookid}` | Kitabı sil | Admin |

### Rezervasyonlar
| Metot | Endpoint | Açıklama | Yetki |
|-------|----------|----------|-------|
| `POST` | `/api/reservations` | Rezervasyon oluştur | User, Admin |
| `GET` | `/api/reservations/{reservationid}` | Rezervasyon detayı | User (kendi), Admin |
| `DELETE` | `/api/reservations/{reservationid}` | Kitap iade et | User (kendi), Admin |
| `GET` | `/api/users/{userid}/reservations` | Kullanıcı rezervasyon listesi | User (kendi), Admin |

### Raporlar
| Metot | Endpoint | Açıklama | Yetki |
|-------|----------|----------|-------|
| `GET` | `/api/reports` | İstatistik raporu | Admin |

---

## Sunucu Adresleri

| Ortam | URL |
|-------|-----|
| Production | `https://api.libranet.com` |
| Staging | `https://staging-api.libranet.com` |
| Development | `https://localhost:3000` |
