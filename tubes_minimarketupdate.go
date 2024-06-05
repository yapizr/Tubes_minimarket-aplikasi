package main

import (
	"fmt"
	"time"
)

// Define constants
const maxBarang = 100
const maxTransaksi = 100
const maxPembayaran = 100

// Tipe Alias
type BarangList [maxBarang]Barang
type TransaksiList [maxTransaksi]Transaksi
type PembayaranList [maxPembayaran]Pembayaran
type KeranjangList [maxBarang]KeranjangItem

// Struct definitions
type Barang struct {
	ID    int
	Nama  string
	Harga float64
}

type Transaksi struct {
	ID         int
	BarangID   int
	Jumlah     int
	TotalHarga float64
	Tanggal    time.Time
}

type Pembayaran struct {
	ID         int
	TotalBayar float64
	Tanggal    time.Time
}

type KeranjangItem struct {
	BarangID int
	Jumlah   int
}

var barangList BarangList
var transaksiList TransaksiList
var pembayaranList PembayaranList
var keranjangList KeranjangList

var jumlahBarang, jumlahTransaksi, jumlahPembayaran, jumlahKeranjang int

func main() {
	for {
		fmt.Println("\nMenu Kasir:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Lihat Daftar Barang")
		fmt.Println("5. Tambah ke Keranjang")
		fmt.Println("6. Lihat Keranjang")
		fmt.Println("7. Bayar")
		fmt.Println("8. Lihat Daftar Transaksi")
		fmt.Println("9. Lihat Daftar Pembayaran")
		fmt.Println("10. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahBarang()
		case 2:
			ubahBarang()
		case 3:
			hapusBarang()
		case 4:
			lihatDaftarBarang()
		case 5:
			tambahKeKeranjang()
		case 6:
			lihatKeranjang()
		case 7:
			bayar()
		case 8:
			lihatDaftarTransaksi()
		case 9:
			lihatDaftarPembayaran()
		case 10:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahBarang() {
	if jumlahBarang >= maxBarang {
		fmt.Println("Tidak bisa menambah barang lagi, kapasitas penuh!")
		return
	}

	var id int
	var nama string
	var harga float64

	fmt.Print("Masukkan ID Barang: ")
	fmt.Scanln(&id)
	fmt.Print("Masukkan Nama Barang: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Harga Barang: ")
	fmt.Scanln(&harga)

	if findBarangByID(id) != -1 {
		fmt.Println("Barang dengan ID tersebut sudah ada")
		return
	}

	barangList[jumlahBarang] = Barang{ID: id, Nama: nama, Harga: harga}
	jumlahBarang++
	fmt.Println("Barang berhasil ditambahkan")

	selectionSortBarang()
}

func ubahBarang() {
	var id int
	fmt.Print("Masukkan ID Barang yang ingin diubah: ")
	fmt.Scanln(&id)

	index := findBarangByID(id)
	if index == -1 {
		fmt.Println("Barang dengan ID tersebut tidak ditemukan")
		return
	}

	var nama string
	var harga float64
	fmt.Print("Masukkan Nama Barang Baru: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Harga Barang Baru: ")
	fmt.Scanln(&harga)

	barangList[index].Nama = nama
	barangList[index].Harga = harga
	fmt.Println("Barang berhasil diubah")

	selectionSortBarang()
}

func hapusBarang() {
	var id int
	fmt.Print("Masukkan ID Barang yang ingin dihapus: ")
	fmt.Scanln(&id)

	index := findBarangByID(id)
	if index == -1 {
		fmt.Println("Barang dengan ID tersebut tidak ditemukan")
		return
	}

	for i := index; i < jumlahBarang-1; i++ {
		barangList[i] = barangList[i+1]
	}
	jumlahBarang--
	fmt.Println("Barang berhasil dihapus")
}

func lihatDaftarBarang() {
	if jumlahBarang == 0 {
		fmt.Println("Tidak ada barang yang tersedia")
		return
	}

	selectionSortBarang()

	fmt.Println("\nDaftar Barang:")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("%d. ID: %d, Nama: %s, Harga: Rp %.2f\n", i+1, barangList[i].ID, barangList[i].Nama, barangList[i].Harga)
	}
}

func tambahKeKeranjang() {
	if jumlahKeranjang >= maxBarang {
		fmt.Println("Keranjang penuh!")
		return
	}

	var id int
	var jumlah int

	fmt.Print("Masukkan ID Barang yang ingin ditambah ke keranjang: ")
	fmt.Scanln(&id)
	fmt.Print("Masukkan Jumlah Barang: ")
	fmt.Scanln(&jumlah)

	index := findBarangByID(id)
	if index == -1 {
		fmt.Println("Barang dengan ID tersebut tidak ditemukan")
		return
	}

	keranjangList[jumlahKeranjang] = KeranjangItem{BarangID: id, Jumlah: jumlah}
	jumlahKeranjang++
	fmt.Println("Barang berhasil ditambah ke keranjang")
}

func lihatKeranjang() {
	if jumlahKeranjang == 0 {
		fmt.Println("Keranjang kosong")
		return
	}

	fmt.Println("\nKeranjang Belanja:")
	for i := 0; i < jumlahKeranjang; i++ {
		barang := barangList[findBarangByID(keranjangList[i].BarangID)]
		fmt.Printf("%d. ID Barang: %d, Nama: %s, Jumlah: %d, Harga: Rp %.2f\n",
			i+1, keranjangList[i].BarangID, barang.Nama, keranjangList[i].Jumlah, float64(keranjangList[i].Jumlah)*barang.Harga)
	}
}

func bayar() {
	if jumlahKeranjang == 0 {
		fmt.Println("Keranjang kosong, tidak ada yang bisa dibayar")
		return
	}

	var total float64
	for i := 0; i < jumlahKeranjang; i++ {
		index := findBarangByID(keranjangList[i].BarangID)
		total += float64(keranjangList[i].Jumlah) * barangList[index].Harga
	}

	fmt.Printf("Total yang harus dibayar: Rp %.2f\n", total)

	var uangDibayar float64
	fmt.Print("Masukkan jumlah uang yang dibayar: ")
	fmt.Scanln(&uangDibayar)

	if uangDibayar < total {
		fmt.Println("Uang yang dibayar kurang, pembayaran gagal")
		return
	}

	kembalian := uangDibayar - total
	fmt.Printf("Pembayaran berhasil, kembalian: Rp %.2f\n", kembalian)

	for i := 0; i < jumlahKeranjang; i++ {
		transaksiList[jumlahTransaksi] = Transaksi{
			ID:         jumlahTransaksi + 1,
			BarangID:   keranjangList[i].BarangID,
			Jumlah:     keranjangList[i].Jumlah,
			TotalHarga: float64(keranjangList[i].Jumlah) * barangList[findBarangByID(keranjangList[i].BarangID)].Harga,
			Tanggal:    time.Now(),
		}
		jumlahTransaksi++
	}

	pembayaranList[jumlahPembayaran] = Pembayaran{
		ID:         jumlahPembayaran + 1,
		TotalBayar: total,
		Tanggal:    time.Now(),
	}
	jumlahPembayaran++

	jumlahKeranjang = 0 // Kosongkan keranjang setelah pembayaran
}

func lihatDaftarTransaksi() {
	if jumlahTransaksi == 0 {
		fmt.Println("Tidak ada transaksi yang tercatat")
		return
	}

	fmt.Println("\nDaftar Transaksi:")
	for i := 0; i < jumlahTransaksi; i++ {
		barang := barangList[findBarangByID(transaksiList[i].BarangID)]
		fmt.Printf("%d. ID Transaksi: %d, Barang: %s, Jumlah: %d, Total Harga: Rp %.2f, Tanggal: %s\n",
			i+1, transaksiList[i].ID, barang.Nama, transaksiList[i].Jumlah, transaksiList[i].TotalHarga, transaksiList[i].Tanggal.Format("02-01-2006 15:04:05"))
	}
}

func lihatDaftarPembayaran() {
	if jumlahPembayaran == 0 {
		fmt.Println("Tidak ada pembayaran yang tercatat")
		return
	}

	fmt.Println("\nDaftar Pembayaran:")
	for i := 0; i < jumlahPembayaran; i++ {
		fmt.Printf("%d. ID Pembayaran: %d, Total Bayar: Rp %.2f, Tanggal: %s\n",
			i+1, pembayaranList[i].ID, pembayaranList[i].TotalBayar, pembayaranList[i].Tanggal.Format("02-01-2006 15:04:05"))
	}
}

func findBarangByID(id int) int {
	for i := 0; i < jumlahBarang; i++ {
		if barangList[i].ID == id {
			return i
		}
	}
	return -1
}

func findTransaksiByID(id int) int {
	for i := 0; i < jumlahTransaksi; i++ {
		if transaksiList[i].ID == id {
			return i
		}
	}
	return -1
}

func selectionSortBarang() {
	for i := 0; i < jumlahBarang-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahBarang; j++ {
			if barangList[j].ID < barangList[minIndex].ID {
				minIndex = j
			}
		}
		barangList[i], barangList[minIndex] = barangList[minIndex], barangList[i]
	}
}
