package main

import (
	"fmt"
	"time"
)

// Define constants
const maxBarang = 100
const maxTransaksi = 100

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

var barangList [maxBarang]Barang
var transaksiList [maxTransaksi]Transaksi

var jumlahBarang, jumlahTransaksi int

func main() {

	for {
		fmt.Println("\nMenu Kasir:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Lihat Daftar Barang")
		fmt.Println("5. Catat Transaksi")
		fmt.Println("6. Lihat Daftar Transaksi")
		fmt.Println("7. Lihat Omzet Harian")
		fmt.Println("8. Keluar")

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
			catatTransaksi()
		case 6:
			lihatDaftarTransaksi()
		case 7:
			lihatOmzetHarian()
		case 8:
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

func catatTransaksi() {
	if jumlahTransaksi >= maxTransaksi {
		fmt.Println("Tidak bisa mencatat transaksi lagi, kapasitas penuh!")
		return
	}

	var transaksi Transaksi

	fmt.Print("Masukkan ID Barang yang dibeli: ")
	fmt.Scanln(&transaksi.BarangID)

	index := findBarangByID(transaksi.BarangID)
	if index == -1 {
		fmt.Println("Barang dengan ID tersebut tidak ditemukan")
		return
	}

	fmt.Print("Masukkan Jumlah Barang: ")
	fmt.Scanln(&transaksi.Jumlah)

	transaksi.ID = jumlahTransaksi + 1
	transaksi.TotalHarga = float64(transaksi.Jumlah) * barangList[index].Harga
	transaksi.Tanggal = time.Now()

	transaksiList[jumlahTransaksi] = transaksi
	jumlahTransaksi++
	fmt.Println("Transaksi berhasil dicatat")
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

func lihatOmzetHarian() {
	today := time.Now().Truncate(24 * time.Hour)
	totalOmzet := 0.0
	fmt.Printf("Omzet Harian %s:\n", today.Format("02-01-2006"))

	for i := 0; i < jumlahTransaksi; i++ {
		if transaksiList[i].Tanggal.Truncate(24 * time.Hour).Equal(today) {
			totalOmzet += transaksiList[i].TotalHarga
			fmt.Printf("ID Transaksi: %d, BarangID: %d, Jumlah: %d, Total Harga: Rp %.2f\n",
				transaksiList[i].ID, transaksiList[i].BarangID, transaksiList[i].Jumlah, transaksiList[i].TotalHarga)
		}
	}
	fmt.Printf("Total Omzet: Rp %.2f\n", totalOmzet)
}

func findBarangByID(id int) int {
	for i := 0; i < jumlahBarang; i++ {
		if barangList[i].ID == id {
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
		if minIndex != i {
			barangList[i], barangList[minIndex] = barangList[minIndex], barangList[i]
		}
	}
}
