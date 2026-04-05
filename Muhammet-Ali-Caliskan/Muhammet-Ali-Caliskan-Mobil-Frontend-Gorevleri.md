# Mobil Front-End Görevleri

**Proje:** LibraNet – Kütüphane Yönetim Sistemi
**Geliştirici:** Muhammet Ali Çalışkan
**Platform:** Mobil (React Native / Flutter)

---

## Navigasyon

Mobil uygulama sayfa geçişlerini Stack ve Tab Navigator kombinasyonu ile yönetir.

- **Stack Navigator:** Giriş → Kayıt → Ana Sayfa akışını yönetir; geri tuşu davranışı platforma özgü şekilde yapılandırılır.
- **Tab Navigator:** Ana ekranda `Kitaplar`, `Rezervasyonlar` ve `Profil` sekmeleri sabit olarak görünür.
- **Modal Ekranlar:** Kitap detayı ve rezervasyon oluşturma, mevcut sayfa üzerine modal olarak açılır.
- **Deep Link Desteği:** `libranet://books/:id` formatında derin bağlantılar desteklenerek bildirim tıklamalarında doğru sayfaya yönlendirme yapılır.
- **Korumalı Rotalar:** Token yoksa uygulama `Auth Stack`'e yönlendirilir; token geçerliyse `App Stack` otomatik açılır.

---

## Performans

Mobil cihazların sınırlı kaynakları gözetilerek aşağıdaki optimizasyonlar uygulanır:

- **FlatList & Windowing:** Kitap listeleri `FlatList` bileşeni ile render edilir; yalnızca ekranda görünen öğeler bellekte tutulur.
- **Image Lazy Loading:** Kitap kapak görselleri kaydırma sırasında tembel yüklenir; düşük çözünürlüklü placeholder önce gösterilir.
- **Sayfalandırma:** `onEndReached` olayı ile sonraki sayfa otomatik çekilir (infinite scroll).
- **Bundle Optimizasyonu:** Kullanılmayan kütüphaneler tree-shaking ile çıkarılır; JS bundle boyutu minimize edilir.
- **Native Driver:** Animasyonlar `useNativeDriver: true` ile GPU'ya devredilir.

---

## Önbellek (Caching)

- **AsyncStorage / SecureStore:** JWT token ve kullanıcı bilgisi güvenli depolama alanında saklanır; uygulama yeniden açıldığında otomatik okunur.
- **API Response Cache:** Kitap listeleri ve kullanıcı rezervasyonları 5 dakika süreyle bellekte tutulur; tekrar istek yapılmaz.
- **Görsel Önbelleği:** İndirilen kitap kapakları disk önbelleğine alınır (react-native-fast-image veya eşdeğeri).
- **Offline Modu:** Bağlantı kesildiğinde son önbelleğe alınmış veriler gösterilir; yenileme yapılamayacağına dair bildirim çıkar.

---

## Hata Yönetimi (Error Handling)

- **Ağ Hatası:** API'ye ulaşılamadığında kullanıcıya toast bildirimi gösterilir ve retry butonu sunulur.
- **401 Unauthorized:** Token süresi dolmuşsa oturum kapatılır, kullanıcı Giriş ekranına yönlendirilir.
- **403 Forbidden:** Yetersiz yetki durumunda kullanıcıya açıklayıcı hata mesajı içeren alert gösterilir.
- **Form Doğrulama:** Giriş ve kayıt formları göndermeden önce istemci tarafında doğrulanır.
- **Çökme Raporlama:** Global `ErrorBoundary` ile yakalanmayan hatalar loglama servisine gönderilir.

---

## Push Bildirimleri

- Rezervasyon iade tarihi yaklaştığında (3 gün önce ve gün içinde) push bildirimi gönderilir.
- Bildirime tıklandığında uygulama ilgili rezervasyon detay sayfasına deep link ile açılır.

---

## Erişilebilirlik (Accessibility)

- Tüm dokunma alanları minimum 44×44 pt boyutundadır.
- `accessibilityLabel` ve `accessibilityHint` prop'ları ekran okuyucu desteği için tanımlanmıştır.
- Dinamik metin boyutu (Dynamic Type) desteklenir.
