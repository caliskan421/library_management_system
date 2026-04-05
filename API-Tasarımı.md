# API Tasarımı — LibraNet OpenAPI 3.0.3 Spesifikasyonu

```yaml
openapi: 3.0.3
info:
  title: LibraNet - Kütüphane Yönetim Sistemi API'si
  version: 1.0.0
  description: >
    LibraNet, kütüphane işlemlerini dijitalleştirmek amacıyla tasarlanmış
    REST API tabanlı bir kütüphane yönetim sistemidir. Kullanıcı kimlik doğrulama,
    kitap katalog yönetimi ve rezervasyon takibi gibi temel işlemleri destekler.
    API, JWT tabanlı kimlik doğrulama ve rol tabanlı yetkilendirme (Kullanıcı / Admin)
    sistemi ile korunmaktadır.
  contact:
    name: Muhammet Ali Çalışkan
    email: caliskanmuhammetali82@gmail.com

servers:
  - url: https://api.libranet.com
    description: Üretim sunucusu (Production)
  - url: https://staging-api.libranet.com
    description: Test sunucusu (Staging)
  - url: https://localhost:3000
    description: Yerel geliştirme sunucusu (Development)

tags:
  - name: Kimlik Doğrulama
    description: Kullanıcı kayıt ve giriş işlemleri
  - name: Kitaplar
    description: Kitap ekleme, listeleme, güncelleme, silme ve arama işlemleri
  - name: Rezervasyonlar
    description: Kitap rezervasyonu oluşturma, görüntüleme ve iade işlemleri
  - name: Raporlar
    description: İstatistik ve raporlama işlemleri (yalnızca Admin)

security:
  - BearerAuth: []

paths:
  # ─────────────────────────────────────────────
  # GEREKSİNİM 2 – Kullanıcı Kaydı (CREATE)
  # ─────────────────────────────────────────────
  /api/auth/register:
    post:
      tags:
        - Kimlik Doğrulama
      summary: Kullanıcı Kaydı
      operationId: registerUser
      security: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/RegisterInput'
      responses:
        "201":
          description: Kullanıcı başarıyla oluşturuldu
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/AuthResponse'
        "400":
          description: Geçersiz istek verisi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "409":
          description: Bu e-posta adresi zaten kayıtlı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  # ─────────────────────────────────────────────
  # GEREKSİNİM 1 – Kullanıcı Girişi (READ)
  # ─────────────────────────────────────────────
  /api/auth/login:
    post:
      tags:
        - Kimlik Doğrulama
      summary: Kullanıcı Girişi
      operationId: loginUser
      security: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/LoginInput'
      responses:
        "200":
          description: Giriş başarılı, JWT token döndürüldü
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/AuthResponse'
        "400":
          description: Geçersiz istek verisi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "401":
          description: E-posta veya şifre hatalı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  # ─────────────────────────────────────────────
  # GEREKSİNİM 3 – Kitap Ekleme (CREATE, Admin)
  # GEREKSİNİM 7 – Kütüphaneler Arası Kitap Arama (READ)
  # ─────────────────────────────────────────────
  /api/books:
    post:
      tags:
        - Kitaplar
      summary: Kitap Ekle (Admin)
      operationId: addBook
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/BookInput'
      responses:
        "201":
          description: Kitap başarıyla eklendi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Book'
        "400":
          description: Geçersiz istek verisi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "403":
          description: Bu işlem için Admin yetkisi gereklidir
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
    get:
      tags:
        - Kitaplar
      summary: Kitap Ara (Kütüphaneler Arası)
      operationId: searchBooks
      parameters:
        - name: query
          in: query
          required: false
          description: Kitap adı, yazar veya ISBN ile arama terimi
          schema:
            type: string
          example: "Dune"
        - name: genre
          in: query
          required: false
          description: Kitap türüne göre filtre
          schema:
            type: string
          example: "Bilim Kurgu"
        - name: available
          in: query
          required: false
          description: Yalnızca müsait kitapları getir
          schema:
            type: boolean
          example: true
        - name: page
          in: query
          required: false
          description: Sayfa numarası (varsayılan 1)
          schema:
            type: integer
            minimum: 1
            default: 1
          example: 1
        - name: limit
          in: query
          required: false
          description: Sayfa başına sonuç sayısı (varsayılan 10, maksimum 50)
          schema:
            type: integer
            minimum: 1
            maximum: 50
            default: 10
          example: 10
      responses:
        "200":
          description: Kitaplar başarıyla listelendi
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Book'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  # ─────────────────────────────────────────────
  # GEREKSİNİM 4  – Kitap Güncelleme (UPDATE, Admin)
  # GEREKSİNİM 5  – Kitap Silme (DELETE, Admin)
  # GEREKSİNİM 6  – Kitap Detayı Görüntüleme (READ)
  # ─────────────────────────────────────────────
  /api/books/{bookid}:
    parameters:
      - name: bookid
        in: path
        required: true
        description: Kitabın benzersiz kimlik numarası
        schema:
          type: string
        example: "book789"

    get:
      tags:
        - Kitaplar
      summary: Kitap Detayını Görüntüle
      operationId: getBook
      responses:
        "200":
          description: Kitap başarıyla getirildi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Book'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "404":
          description: Kitap bulunamadı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

    put:
      tags:
        - Kitaplar
      summary: Kitap Güncelle (Admin)
      operationId: updateBook
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/BookInput'
      responses:
        "200":
          description: Kitap başarıyla güncellendi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Book'
        "400":
          description: Geçersiz istek verisi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "403":
          description: Bu işlem için Admin yetkisi gereklidir
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "404":
          description: Kitap bulunamadı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

    delete:
      tags:
        - Kitaplar
      summary: Kitap Sil (Admin)
      operationId: deleteBook
      responses:
        "204":
          description: Kitap başarıyla silindi
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "403":
          description: Bu işlem için Admin yetkisi gereklidir
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "404":
          description: Kitap bulunamadı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  # ─────────────────────────────────────────────
  # GEREKSİNİM 8  – Kitap Rezervasyonu Oluşturma (CREATE)
  # ─────────────────────────────────────────────
  /api/reservations:
    post:
      tags:
        - Rezervasyonlar
      summary: Rezervasyon Oluştur
      operationId: createReservation
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/ReservationInput'
      responses:
        "201":
          description: Rezervasyon başarıyla oluşturuldu
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Reservation'
        "400":
          description: Geçersiz istek verisi veya kitap müsait değil
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "404":
          description: Kitap bulunamadı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  # ─────────────────────────────────────────────
  # GEREKSİNİM 9  – Kitap İadesi (DELETE)
  # GEREKSİNİM 10 – Rezervasyon Detayı Görüntüleme (READ)
  # ─────────────────────────────────────────────
  /api/reservations/{reservationid}:
    parameters:
      - name: reservationid
        in: path
        required: true
        description: Rezervasyonun benzersiz kimlik numarası
        schema:
          type: string
        example: "rsv001"

    get:
      tags:
        - Rezervasyonlar
      summary: Rezervasyon Detayını Görüntüle
      operationId: getReservation
      responses:
        "200":
          description: Rezervasyon başarıyla getirildi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Reservation'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "403":
          description: Bu rezervasyona erişim yetkiniz bulunmuyor
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "404":
          description: Rezervasyon bulunamadı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

    delete:
      tags:
        - Rezervasyonlar
      summary: Kitap İade Et (Rezervasyonu Kapat)
      operationId: returnBook
      responses:
        "204":
          description: Kitap başarıyla iade edildi, rezervasyon kapatıldı
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "403":
          description: Bu rezervasyona erişim yetkiniz bulunmuyor
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "404":
          description: Rezervasyon bulunamadı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  # ─────────────────────────────────────────────
  # GEREKSİNİM 11 – Kullanıcı Rezervasyonlarını Listeleme (READ)
  # ─────────────────────────────────────────────
  /api/users/{userid}/reservations:
    parameters:
      - name: userid
        in: path
        required: true
        description: Kullanıcının benzersiz kimlik numarası
        schema:
          type: string
        example: "usr123"

    get:
      tags:
        - Rezervasyonlar
      summary: Kullanıcının Rezervasyonlarını Listele
      operationId: listUserReservations
      parameters:
        - name: status
          in: query
          required: false
          description: Rezervasyon durumuna göre filtre (active, returned, all)
          schema:
            type: string
            enum: [active, returned, all]
            default: all
          example: "active"
        - name: page
          in: query
          required: false
          description: Sayfa numarası (varsayılan 1)
          schema:
            type: integer
            minimum: 1
            default: 1
          example: 1
        - name: limit
          in: query
          required: false
          description: Sayfa başına sonuç sayısı (varsayılan 10, maksimum 50)
          schema:
            type: integer
            minimum: 1
            maximum: 50
            default: 10
          example: 10
      responses:
        "200":
          description: Rezervasyonlar başarıyla listelendi
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Reservation'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "403":
          description: Bu kullanıcının rezervasyonlarına erişim yetkiniz bulunmuyor
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "404":
          description: Kullanıcı bulunamadı
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  # ─────────────────────────────────────────────
  # GEREKSİNİM 12 – Raporlama ve İstatistikler (READ, Admin)
  # ─────────────────────────────────────────────
  /api/reports:
    get:
      tags:
        - Raporlar
      summary: İstatistik ve Raporları Getir (Admin)
      operationId: getReports
      parameters:
        - name: type
          in: query
          required: false
          description: Rapor türü (books, reservations, users)
          schema:
            type: string
            enum: [books, reservations, users]
          example: "reservations"
        - name: from
          in: query
          required: false
          description: Başlangıç tarihi (ISO 8601)
          schema:
            type: string
            format: date
          example: "2026-01-01"
        - name: to
          in: query
          required: false
          description: Bitiş tarihi (ISO 8601)
          schema:
            type: string
            format: date
          example: "2026-03-09"
      responses:
        "200":
          description: Rapor başarıyla getirildi
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Report'
        "401":
          description: Kimlik doğrulama başarısız (token eksik veya geçersiz)
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
        "403":
          description: Bu işlem için Admin yetkisi gereklidir
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

components:
  securitySchemes:
    BearerAuth:
      type: apiKey
      in: header
      name: Authorization
      description: 'JWT tabanlı kimlik doğrulama. İstek başlığına "Authorization: Bearer <token>" eklenmeli.'

  schemas:
    # ── Kimlik Doğrulama Şemaları ──────────────
    RegisterInput:
      type: object
      description: Kullanıcı kayıt isteği için gönderilecek veri
      properties:
        name:
          type: string
          description: Kullanıcının adı soyadı
          minLength: 2
          maxLength: 100
          example: "Muhammet Ali Çalışkan"
        email:
          type: string
          format: email
          description: Kullanıcının e-posta adresi
          example: "ali@example.com"
        password:
          type: string
          format: password
          description: Kullanıcının şifresi (en az 8 karakter)
          minLength: 8
          example: "gizli1234"
      required:
        - name
        - email
        - password

    LoginInput:
      type: object
      description: Kullanıcı giriş isteği için gönderilecek veri
      properties:
        email:
          type: string
          format: email
          description: Kayıtlı e-posta adresi
          example: "ali@example.com"
        password:
          type: string
          format: password
          description: Kullanıcı şifresi
          example: "gizli1234"
      required:
        - email
        - password

    AuthResponse:
      type: object
      description: Başarılı kimlik doğrulama sonrası dönen yanıt
      properties:
        token:
          type: string
          description: JWT erişim token'ı
          example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
        user:
          $ref: '#/components/schemas/User'
      required:
        - token
        - user

    User:
      type: object
      description: Kullanıcı bilgilerini temsil eden model
      properties:
        _id:
          type: string
          description: Kullanıcının benzersiz kimlik numarası (otomatik atanır)
          example: "usr123"
        name:
          type: string
          description: Kullanıcının adı soyadı
          example: "Muhammet Ali Çalışkan"
        email:
          type: string
          format: email
          description: Kullanıcının e-posta adresi
          example: "ali@example.com"
        role:
          type: string
          description: Kullanıcı rolü
          enum: [user, admin]
          example: "user"
        createdAt:
          type: string
          format: date-time
          description: Hesabın oluşturulma tarihi (ISO 8601 formatında)
          example: "2026-01-15T09:00:00Z"
      required:
        - _id
        - name
        - email
        - role

    # ── Kitap Şemaları ──────────────────────────
    Book:
      type: object
      description: Kitap bilgilerini temsil eden model
      properties:
        _id:
          type: string
          description: Kitabın benzersiz kimlik numarası (otomatik atanır)
          example: "book789"
        title:
          type: string
          description: Kitabın adı
          example: "Dune"
        author:
          type: string
          description: Kitabın yazarı
          example: "Frank Herbert"
        isbn:
          type: string
          description: Kitabın ISBN numarası
          example: "978-0-441-01359-7"
        genre:
          type: string
          description: Kitabın türü
          example: "Bilim Kurgu"
        publishedYear:
          type: integer
          description: Yayın yılı
          example: 1965
        available:
          type: boolean
          description: Kitabın şu an müsait olup olmadığı
          example: true
        totalCopies:
          type: integer
          description: Kütüphanedeki toplam kopya sayısı
          minimum: 1
          example: 3
        availableCopies:
          type: integer
          description: Müsait kopya sayısı
          minimum: 0
          example: 2
      required:
        - _id
        - title
        - author
        - isbn
        - available

    BookInput:
      type: object
      description: Kitap oluşturma veya güncelleme isteği için gönderilecek veri
      properties:
        title:
          type: string
          description: Kitabın adı
          minLength: 1
          maxLength: 200
          example: "Dune"
        author:
          type: string
          description: Kitabın yazarı
          minLength: 2
          maxLength: 100
          example: "Frank Herbert"
        isbn:
          type: string
          description: Kitabın ISBN numarası
          example: "978-0-441-01359-7"
        genre:
          type: string
          description: Kitabın türü
          example: "Bilim Kurgu"
        publishedYear:
          type: integer
          description: Yayın yılı
          example: 1965
        totalCopies:
          type: integer
          description: Kütüphanedeki toplam kopya sayısı
          minimum: 1
          example: 3
      required:
        - title
        - author
        - isbn
        - totalCopies

    # ── Rezervasyon Şemaları ────────────────────
    Reservation:
      type: object
      description: Rezervasyon bilgilerini temsil eden model
      properties:
        _id:
          type: string
          description: Rezervasyonun benzersiz kimlik numarası (otomatik atanır)
          example: "rsv001"
        userId:
          type: string
          description: Rezervasyonu yapan kullanıcının kimlik numarası
          example: "usr123"
        bookId:
          type: string
          description: Rezerve edilen kitabın kimlik numarası
          example: "book789"
        status:
          type: string
          description: Rezervasyonun durumu
          enum: [active, returned]
          example: "active"
        reservedAt:
          type: string
          format: date-time
          description: Rezervasyonun oluşturulma tarihi (ISO 8601 formatında)
          example: "2026-03-01T10:00:00Z"
        dueDate:
          type: string
          format: date-time
          description: Kitabın iade edilmesi gereken son tarih (ISO 8601 formatında)
          example: "2026-03-15T10:00:00Z"
        returnedAt:
          type: string
          format: date-time
          description: Kitabın iade edildiği tarih (ISO 8601 formatında, iade edilmediyse null)
          example: null
          nullable: true
      required:
        - _id
        - userId
        - bookId
        - status
        - reservedAt
        - dueDate

    ReservationInput:
      type: object
      description: Rezervasyon oluşturma isteği için gönderilecek veri
      properties:
        bookId:
          type: string
          description: Rezerve edilmek istenen kitabın kimlik numarası
          example: "book789"
        dueDate:
          type: string
          format: date
          description: İstenen iade tarihi (ISO 8601 formatında)
          example: "2026-03-15"
      required:
        - bookId
        - dueDate

    # ── Rapor Şeması ────────────────────────────
    Report:
      type: object
      description: İstatistik ve raporlama verilerini temsil eden model
      properties:
        generatedAt:
          type: string
          format: date-time
          description: Raporun oluşturulma tarihi (ISO 8601 formatında)
          example: "2026-03-09T12:00:00Z"
        type:
          type: string
          description: Rapor türü
          enum: [books, reservations, users]
          example: "reservations"
        summary:
          type: object
          description: Genel özet istatistikleri
          properties:
            total:
              type: integer
              description: Toplam kayıt sayısı
              example: 450
            active:
              type: integer
              description: Aktif kayıt sayısı (rezervasyonlar için aktif ödünç)
              example: 120
            returned:
              type: integer
              description: Tamamlanan kayıt sayısı (iade edilen kitaplar)
              example: 330
        topBooks:
          type: array
          description: En çok rezerve edilen kitaplar (books/reservations raporunda)
          items:
            type: object
            properties:
              bookId:
                type: string
                example: "book789"
              title:
                type: string
                example: "Dune"
              reservationCount:
                type: integer
                example: 42
      required:
        - generatedAt
        - type
        - summary

    # ── Hata Şeması ─────────────────────────────
    Error:
      type: object
      description: Hata durumlarında döndürülen standart hata yanıtı
      properties:
        message:
          type: string
          description: Hatayı açıklayan mesaj
          example: "Kitap bulunamadı"
      required:
        - message
```