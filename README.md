
# CLI Expense Tracker

Aplikasi pencatatan pengeluaran berbasis Command Line Interface (CLI) yang dibuat menggunakan Go. Project ini bertujuan untuk mempelajari Go Standard Library melalui project sederhana.

## Features

- Menambahkan pengeluaran
- Melihat daftar pengeluaran
- Menghitung ringkasan pengeluaran
- Mengubah data pengeluaran
- Menghapus data pengeluaran
- Mencari pengeluaran berdasarkan deskripsi
- Memfilter pengeluaran berdasarkan kategori
- Mengurutkan pengeluaran berdasarkan nominal
- Mengekspor data ke CSV
- Menyimpan data menggunakan JSON
- Menjalankan export melalui CLI argument
- Unit testing

## Tech Stack

Project ini menggunakan Go Standard Library:

- `fmt` — input dan output
- `bufio` — membaca input terminal
- `strconv` — konversi tipe data
- `strings` — manipulasi string
- `time` — pengelolaan tanggal
- `os` — operasi file
- `encoding/json` — penyimpanan JSON
- `slices` — manipulasi slice
- `errors` — penanganan error
- `regexp` — pencarian menggunakan regular expression
- `sort` — pengurutan data
- `math` — perhitungan matematika
- `encoding/csv` — ekspor CSV
- `flag` — CLI arguments
- `testing` — unit testing

## Project Structure

```text
expense-tracker/
├── internal/
│   └── expense/
│       ├── expense.go
│       ├── storage.go
│       ├── operations.go
│       ├── query.go
│       ├── summary.go
│       ├── export.go
│       └── summary_test.go
├── main.go
├── go.mod
├── expenses.json
├── expenses.csv
└── README.md
```

## Requirements

- Go 1.21 atau lebih baru

## Installation

Clone repository:

```bash
git clone <repository-url>
```

Masuk ke direktori project:

```bash
cd expense-tracker
```

Jalankan aplikasi:

```bash
go run .
```

## Usage

Jalankan aplikasi:

```bash
go run .
```

Pilih menu yang tersedia untuk mengelola pengeluaran.

### Export CSV

Untuk langsung mengekspor data ke CSV tanpa membuka menu:

```bash
go run . -export
```

File hasil ekspor akan disimpan sebagai `expenses.csv`.

## Testing

Jalankan unit test:

```bash
go test ./...
```

## Data Storage

Data pengeluaran disimpan dalam file `expenses.json` menggunakan format JSON.

Data dapat diekspor ke file `expenses.csv` untuk digunakan pada aplikasi spreadsheet seperti Microsoft Excel atau Google Sheets.

## Learning Objectives

Melalui project ini, saya mempelajari:

- Dasar bahasa pemrograman Go
- Struct dan slice
- Function dan package
- Error handling
- File handling
- JSON dan CSV
- Sorting dan filtering
- Regular expression
- CLI arguments
- Unit testing
- Penggunaan Go Standard Library

## License

This project is open-source and available for learning purposes.
