# Mobil Backend Görevleri

**Proje:** LibraNet – Kütüphane Yönetim Sistemi
**Geliştirici:** Muhammet Ali Çalışkan
**Rol:** Mobil Backend (REST API altyapısı, veri katmanı)

---

## Mimari Genel Bakış

LibraNet backend'i **Node.js + Express** (veya eşdeğer) üzerine kuruludur. Katmanlı mimari ile sorumluluklar ayrıştırılmıştır:

```
Katman         Sorumluluk
─────────────────────────────────────────
Router         HTTP rotaları ve parametre doğrulama
Controller     İş mantığı ve yanıt üretimi
Service        Veri işleme ve kurallar
Repository     Veritabanı sorguları
Model          Şema tanımları (Mongoose / ORM)
Middleware     Auth, hata yakalama, loglama
```

---

## Kimlik Doğrulama ve Yetkilendirme

- **JWT:** Kullanıcı girişinde imzalı token üretilir; geçerlilik süresi 24 saattir.
- **Refresh Token:** 30 günlük refresh token ile sessiz oturum yenileme sağlanır.
- **`authenticate` Middleware:** Her korunan rotada `Authorization` başlığındaki token doğrulanır.
- **`authorize(role)` Middleware:** Admin gerektiren rotalarda rol kontrolü yapılır; yetersiz yetki durumunda 403 döner.

---

## Veri Katmanı (Database)

- **Veritabanı:** MongoDB (NoSQL) — Kitap ve rezervasyon dökümanları için esnek şema.
- **Koleksiyonlar:**

| Koleksiyon | Açıklama |
|-----------|---------|
| `users` | Kullanıcı kayıtları (ad, e-posta, hash'lenmiş şifre, rol) |
| `books` | Kitap kataloğu (başlık, yazar, ISBN, kopya sayısı) |
| `reservations` | Rezervasyon kayıtları (userId, bookId, tarihler, durum) |

- **İndeksler:** `email` alanına tekil indeks; `bookId` ve `status` alanlarına bileşik indeks ile sorgu performansı artırılır.

---

## Performans

- **Sayfalandırma:** Tüm liste endpointleri `page` ve `limit` parametreleri ile sayfalandırılmış yanıt döner; `skip/limit` yerine imleç (cursor) tabanlı sayfalandırma büyük veri setlerinde tercih edilir.
- **Projeksiyon:** Veritabanı sorgularında yalnızca ihtiyaç duyulan alanlar seçilir; gereksiz veri transferi önlenir.
- **Bağlantı Havuzu (Connection Pooling):** Mongoose bağlantı havuzu yapılandırılarak her istek için yeni bağlantı açılması engellenir.
- **Async/Await:** Tüm G/Ç işlemleri asenkron olarak gerçekleştirilir; bloklanma önlenir.

---

## Önbellek (Caching)

- **In-Memory Cache (Redis):** Rapor verileri ve sık okunan kitap listeleri Redis'te 10 dakika önbelleklenir.
- **Cache Invalidation:** Kitap eklendiğinde veya rezervasyon değiştiğinde ilgili önbellek anahtarları temizlenir.
- **ETag Desteği:** GET yanıtlarına `ETag` başlığı eklenir; değişmeyen içerik için istemci 304 ile cevap alır.

---

## Hata Yönetimi (Error Handling)

- **Global Error Handler:** Express hata middleware'i tüm route hatalarını merkezi olarak yakalar ve standart `{ message }` formatında JSON döner.
- **Doğrulama Hataları:** Request body, `express-validator` veya `Joi` ile doğrulanır; eksik alan durumunda açıklayıcı 400 mesajı döner.
- **Async Hata Sarmalayıcı:** Her controller fonksiyonu `asyncHandler` wrapper'ı ile sarmalanır; try-catch tekrarı ortadan kalkar.
- **Loglama:** `morgan` HTTP istek loglama; `winston` uygulama düzeyinde yapılandırılmış loglama.

---

## Güvenlik

- **Şifre Hashleme:** Kullanıcı şifreleri `bcrypt` (10 salt round) ile saklanır; düz metin asla veritabanına yazılmaz.
- **Rate Limiting:** `express-rate-limit` ile `/api/auth` endpointlerine dakikada maksimum 10 istek sınırı uygulanır.
- **Helmet:** HTTP güvenlik başlıkları (`X-Content-Type-Options`, `X-Frame-Options` vb.) otomatik eklenir.
- **CORS:** Yalnızca tanımlı origin adreslerinden gelen isteklere izin verilir.
- **Input Sanitization:** Tüm kullanıcı girdileri veritabanına yazılmadan önce temizlenir; injection saldırıları önlenir.

---

## API Dokümantasyonu

Tüm endpointlerin detaylı tanımları için → **[REST API Tasarımı](../REST-API-Tasarimi.md)** ve **[libranet.yaml](../API-Tasarımı.md)**
