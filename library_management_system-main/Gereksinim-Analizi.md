# Gereksinim Analizi

**Proje:** LibraNet – Kütüphane Yönetim Sistemi
**Geliştirici:** Muhammet Ali Çalışkan
**Grup:** CaliskanAli

---

## Sistem Aktörleri

| Aktör | Açıklama |
|-------|----------|
| **Kullanıcı (User)** | Sisteme kayıt olmuş, kitap arayabilen ve rezervasyon yapabilen standart kullanıcı |
| **Admin** | Kitap kataloğunu yönetme (ekleme, güncelleme, silme) ve raporlara erişme yetkisine sahip yönetici |

---

## Gereksinim Tablosu

| # | Gereksinim | İşlem Türü | Aktör | Öncelik |
|---|-----------|-----------|-------|---------|
| 1 | Kullanıcı Girişi | READ | Kullanıcı, Admin | Yüksek |
| 2 | Kullanıcı Kaydı | CREATE | Kullanıcı | Yüksek |
| 3 | Kitap Ekleme | CREATE | Admin | Yüksek |
| 4 | Kitap Güncelleme | UPDATE | Admin | Orta |
| 5 | Kitap Silme | DELETE | Admin | Orta |
| 6 | Kitap Detayı Görüntüleme | READ | Kullanıcı, Admin | Yüksek |
| 7 | Kütüphaneler Arası Kitap Arama | READ | Kullanıcı, Admin | Yüksek |
| 8 | Kitap Rezervasyonu Oluşturma | CREATE | Kullanıcı | Yüksek |
| 9 | Kitap İadesi | DELETE | Kullanıcı | Yüksek |
| 10 | Rezervasyon Detayı Görüntüleme | READ | Kullanıcı, Admin | Orta |
| 11 | Kullanıcıya Ait Rezervasyonları Listeleme | READ | Kullanıcı, Admin | Orta |
| 12 | Raporlama & İstatistik | READ | Admin | Düşük |

---

## Detaylı Gereksinim Açıklamaları

### 1. Kullanıcı Girişi (READ)
Sisteme kayıtlı kullanıcıların e-posta ve şifre bilgisiyle kimlik doğrulaması yaparak JWT token almasını sağlar. Başarılı girişin ardından token ile korunan tüm endpointlere erişim mümkün hale gelir.

**Girdi:** `email`, `password`
**Çıktı:** JWT token, kullanıcı bilgileri
**Kısıt:** Hatalı bilgi girişinde 401 hatası döner; 5 başarısız denemede hesap geçici olarak kilitlenir.

---

### 2. Kullanıcı Kaydı (CREATE)
Yeni kullanıcıların sisteme kayıt olmasını sağlar. Kayıt sırasında ad-soyad, e-posta ve şifre zorunludur. Aynı e-posta ile birden fazla kayıt oluşturulamaz.

**Girdi:** `name`, `email`, `password`
**Çıktı:** Oluşturulan kullanıcı bilgileri, JWT token
**Kısıt:** E-posta benzersiz olmalı; şifre en az 8 karakter içermeli.

---

### 3. Kitap Ekleme (CREATE) — Admin
Admin yetkisine sahip kullanıcıların kitap kataloğuna yeni kitap eklemesini sağlar.

**Girdi:** `title`, `author`, `isbn`, `genre`, `publishedYear`, `totalCopies`
**Çıktı:** Oluşturulan kitap kaydı
**Kısıt:** Yalnızca Admin rolü erişebilir; ISBN tekil olmalı.

---

### 4. Kitap Güncelleme (UPDATE) — Admin
Mevcut bir kitabın bilgilerinin (başlık, yazar, kopya sayısı vb.) güncellenmesini sağlar.

**Girdi:** `bookid` (path), güncellenecek alanlar
**Çıktı:** Güncellenmiş kitap kaydı
**Kısıt:** Yalnızca Admin rolü erişebilir; kitap mevcut olmalı.

---

### 5. Kitap Silme (DELETE) — Admin
Katalogdan bir kitabın kaldırılmasını sağlar.

**Girdi:** `bookid` (path)
**Çıktı:** 204 No Content
**Kısıt:** Yalnızca Admin rolü erişebilir; aktif rezervasyonu olan kitap silinemez.

---

### 6. Kitap Detayı Görüntüleme (READ)
Belirli bir kitabın tüm bilgilerini (başlık, yazar, ISBN, müsaitlik durumu vb.) görüntülemeyi sağlar.

**Girdi:** `bookid` (path)
**Çıktı:** Kitap detay bilgileri
**Kısıt:** Kimlik doğrulaması gerekir.

---

### 7. Kütüphaneler Arası Kitap Arama (READ)
Kitap adı, yazar veya ISBN ile tüm katalogda arama yapılmasını sağlar. Tür ve müsaitlik durumuna göre filtreleme desteklenir.

**Girdi:** `query`, `genre`, `available`, `page`, `limit` (query params)
**Çıktı:** Eşleşen kitap listesi (sayfalandırılmış)
**Kısıt:** Kimlik doğrulaması gerekir.

---

### 8. Kitap Rezervasyonu Oluşturma (CREATE)
Müsait bir kitap için ödünç rezervasyonu oluşturulmasını sağlar.

**Girdi:** `bookId`, `dueDate`
**Çıktı:** Oluşturulan rezervasyon kaydı
**Kısıt:** Kitap müsait olmalı; kullanıcı başına maksimum aktif rezervasyon sayısı uygulanabilir.

---

### 9. Kitap İadesi (DELETE)
Ödünç alınan bir kitabın iade edilmesini ve rezervasyonun kapatılmasını sağlar.

**Girdi:** `reservationid` (path)
**Çıktı:** 204 No Content
**Kısıt:** Yalnızca rezervasyonu oluşturan kullanıcı veya Admin iade işlemi yapabilir.

---

### 10. Rezervasyon Detayı Görüntüleme (READ)
Belirli bir rezervasyonun tüm bilgilerini (kitap, tarihler, durum) görüntülemeyi sağlar.

**Girdi:** `reservationid` (path)
**Çıktı:** Rezervasyon detay bilgileri
**Kısıt:** Yalnızca ilgili kullanıcı veya Admin erişebilir.

---

### 11. Kullanıcıya Ait Rezervasyonları Listeleme (READ)
Bir kullanıcının tüm aktif ve geçmiş rezervasyonlarının listelenmesini sağlar.

**Girdi:** `userid` (path), `status`, `page`, `limit` (query params)
**Çıktı:** Rezervasyon listesi (sayfalandırılmış)
**Kısıt:** Kullanıcı yalnızca kendi listesini görebilir; Admin herkesi görebilir.

---

### 12. Raporlama & İstatistik (READ) — Admin
Kütüphane operasyonlarına ait istatistiksel verilerin (en çok okunan kitaplar, aktif rezervasyon sayısı vb.) raporlanmasını sağlar.

**Girdi:** `type`, `from`, `to` (query params)
**Çıktı:** Özet istatistik raporu
**Kısıt:** Yalnızca Admin rolü erişebilir.
