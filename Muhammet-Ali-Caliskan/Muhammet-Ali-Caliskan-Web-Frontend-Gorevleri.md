# Web Front-End Görevleri

**Proje:** LibraNet – Kütüphane Yönetim Sistemi
**Geliştirici:** Muhammet Ali Çalışkan
**Platform:** Web (Tarayıcı)

---

## Navigasyon

Web uygulaması sayfa geçişlerini `React Router` (veya eşdeğer bir istemci yönlendirici) aracılığıyla yönetir.

- **Korumalı Rotalar:** JWT token yoksa kullanıcı otomatik olarak `/login` sayfasına yönlendirilir.
- **Rol Bazlı Yönlendirme:** Admin kullanıcılar `/admin` prefix'li rotalara erişebilir; standart kullanıcılar bu rotalara girmeye çalışırsa 403 sayfasına düşer.
- **Sayfa Yapısı:**
  - `/login` → Giriş ekranı
  - `/register` → Kayıt ekranı
  - `/books` → Kitap arama & listeleme
  - `/books/:id` → Kitap detay sayfası
  - `/reservations` → Kullanıcı rezervasyon listesi
  - `/reservations/:id` → Rezervasyon detay sayfası
  - `/admin/books` → Kitap yönetimi (Admin)
  - `/admin/reports` → Raporlar (Admin)

---

## Durum Yönetimi (State Management)

Uygulama genelindeki veriler merkezi bir state store üzerinden yönetilir (Redux, Zustand veya React Context).

- **Auth State:** Kullanıcı bilgisi ve JWT token `localStorage`'a kaydedilir; sayfa yenilemede otomatik olarak okunur.
- **Book State:** Arama sonuçları ve seçili kitap detayı store'da tutulur; gereksiz API çağrısı önlenir.
- **Reservation State:** Aktif rezervasyonlar store'da önbelleğe alınır; iade işlemi sonrası otomatik güncellenir.

---

## Performans

Kullanıcı deneyimini etkileyen yavaş yüklemeleri önlemek için aşağıdaki stratejiler uygulanır:

- **Lazy Loading:** Her rota bileşeni dinamik `import()` ile yüklenerek başlangıç bundle boyutu küçültülür.
- **Sayfalandırma (Pagination):** Kitap ve rezervasyon listelerinde sayfa başına 10 sonuç gösterilir; kullanıcı kaydırdıkça sonraki sayfa yüklenir.
- **Debounce:** Kitap arama inputuna yazma sırasında API çağrısı 400ms geciktirilir; her tuş basışında istek yapılmaz.
- **Skeleton Screens:** Veri yüklenirken boş ekran yerine iskelet bileşenler gösterilir.

---

## Önbellek (Caching)

- **API Önbelleği:** Aynı kitap detay sayfası 5 dakika içinde tekrar ziyaret edildiğinde API'ye yeni istek yapılmaz; store'daki veri kullanılır.
- **Token Yenileme:** JWT süresi dolmadan 1 dakika önce arka planda `refresh token` isteği gönderilir.
- **Offline Desteği:** Service Worker ile son görüntülenen kitap listesi önbellekte tutulur; bağlantı kesildiğinde bilgi mesajı gösterilir.

---

## Hata Yönetimi (Error Handling)

- **401 Unauthorized:** Token geçersiz veya süresi dolmuşsa kullanıcı oturumu kapatılır ve `/login`'e yönlendirilir.
- **403 Forbidden:** Yetkisiz sayfa erişiminde kullanıcıya açıklayıcı hata mesajı gösterilir.
- **404 Not Found:** Bulunamayan kitap veya rezervasyon için özel 404 bileşeni render edilir.
- **Ağ Hatası:** API yanıt vermezse ekranda yeniden deneme butonu ile birlikte hata bildirimi gösterilir.
- **Form Doğrulama:** Giriş ve kayıt formları gönderilmeden önce istemci tarafında doğrulanır (boş alan, e-posta formatı, şifre uzunluğu).

---

## Erişilebilirlik (Accessibility)

- Tüm butonlar ve form alanları `aria-label` ile etiketlenmiştir.
- Klavye navigasyonu desteklenir (Tab, Enter, Escape).
- Renk kontrastı WCAG 2.1 AA standardına uygundur.
