# Web Boilerplate

Proyek boilerplate untuk aplikasi web modern menggunakan Go dan Fiber framework, dengan templating Pongo2 dan UI Tabler.

## Deskripsi

Boilerplate ini menyediakan struktur dasar untuk membangun aplikasi web dengan:
- Backend Go menggunakan Fiber
- Templating dengan Pongo2
- UI admin dengan Tabler CSS framework
- Hot reload untuk development
- Struktur modular dengan handlers, repositories, dan utils

## Fitur

- ✅ Server web dengan Fiber
- ✅ Templating engine Pongo2
- ✅ UI admin responsif dengan Tabler
- ✅ Sidebar collapsible dengan persistent state
- ✅ Hot reload untuk development
- ✅ Struktur proyek yang terorganisir
- ✅ Static file serving
- ✅ Middleware logger

## Prerequisites

- Go 1.19 atau lebih baru
- entr (untuk hot reload): `brew install entr` atau `sudo apt install entr`

## Instalasi

1. Clone repository:
```bash
git clone https://github.com/i-sub135/web-boilerplate.git
cd web-boilerplate
```

2. Install dependencies:
```bash
go mod download
```

## Menjalankan

### Development (dengan hot reload)
```bash
make dev
```
Server akan berjalan di `http://localhost:3003` dan reload otomatis saat file Go atau HTML berubah.

### Production
```bash
make build
make run
```

### Atau langsung:
```bash
go run .
```

## Struktur Proyek

```
web-boilerplate/
├── main.go                 # Entry point aplikasi
├── go.mod                  # Go modules
├── Makefile                # Build scripts
├── common/                 # Kode bersama
│   ├── models/            # Model data
│   ├── repositories/      # Data access layer
│   └── utils/
│       └── webutil/       # Utility untuk web
├── internal/               # Kode internal aplikasi
│   ├── api/               # API handlers
│   └── app/               # HTTP server setup
├── templates/              # Template HTML
│   ├── base.html          # Layout utama
│   └── navigation.html    # Sidebar navigation
├── public/                 # Static files
│   ├── lib/               # Libraries (Tabler, Bootstrap Icons)
│   └── js/                # JavaScript files
└── internal/web/           # Web features
    └── exampletocopy/     # Contoh fitur
```

## Teknologi

- **Backend**: Go, Fiber v2
- **Templating**: Pongo2 v7
- **UI**: Tabler CSS, Bootstrap Icons
- **Build**: Makefile, entr untuk hot reload

## Development

- Gunakan `make dev` untuk development dengan hot reload
- Template disimpan di `templates/`
- Static files di `public/`
- Fitur web di `internal/web/`

## Kontribusi

1. Fork repository
2. Buat branch fitur baru
3. Commit perubahan
4. Push ke branch
5. Buat Pull Request

## Lisensi

MIT License