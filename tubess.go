package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Ini tipe data utama/global
type tablahir struct {
	bulan   []string
	tanggal []string
	tahun   []string
	tempat  []string
}

type dataDiri struct {
	nama    []string
	usia    []string
	ttl     tablahir
	kerjaan []string
	jenisK  []string
	toefl   []int
}

var kerja = [3]string{"Android Developer", "Data Scientist", "Software Engineer"}
var jk = [2]string{"Laki-laki", "Perempuan"}
var daftar dataDiri
var thn, bln, tgl, stoefl int
var jenk, i, usia int
var namalengkap, tempatlahirlengkap string
var del = 1

func main() {
	welcome()
}

func welcome() {
	var pilih int
	fmt.Println(" ____________________________________________________________")
	fmt.Println("|   SELAMAT DATANG DI APLIKASI SELEKSI CALON PEGAWAI BARU!!  |")
	fmt.Println("|------------------------------------------------------------|")
	fmt.Println("|  Silakan pilih menu di bawah ini                           |")
	fmt.Println("| (1) Login sebagai admin                                    |")
	fmt.Println("| (2) Login sebagai guest                                    |")
	fmt.Println("|  Tekan tombol sembarang untuk keluar aplikasi              |")
	fmt.Println("|____________________________________________________________|\n")
	fmt.Print("-> ")
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		login_admin()
	case 2:
		login_guest()
	default:
		fmt.Printf("\nTerima kasih telah menggunakan layanan kami!\n")
	}
}

func login_admin() {
	var pilih string
	fmt.Printf(" ____________________________________________________________\n")
	fmt.Printf("|            SELAMAT DATANG DI LAMAN LOGIN ADMIN             |\n")
	fmt.Printf("|           APLIKASI SELEKSI CALON PEGAWAI BARU!!            |\n")
	fmt.Printf("|------------------------------------------------------------|\n")
	fmt.Printf("|  Masukkan Kode OTP untuk bisa akses                        |\n")
	fmt.Printf("|------------------------------------------------------------|\n")
	fmt.Printf("|  Jika belum memiliki akses admin,                          |\n")
	fmt.Printf("|  sila hubungi developer!                                   |\n")
	fmt.Printf("|------------------------------------------------------------|\n")
	fmt.Printf("|  Untuk kembali ke menu utama,                              |\n")
	fmt.Printf("|  sila masukkan angka 1                                     |\n")
	fmt.Printf("|____________________________________________________________|\n")
	fmt.Print("\n-> ")
	fmt.Scanf("%s\n", &pilih)
	enkrip := enkripsi(pilih)
	if enkrip {
		menu_admin()
	} else {
		switch pilih {
		case "1":
			welcome()
		default:
			login_admin()
		}
	}
}

func akses_admin() {
	var pilih, pilih2 int
	fmt.Printf("|    Nomor Indeks; Nama; Usia; Jenis Kelamin; Tempat dan Tanggal Lahir; Lowongan Kerja yang Dipilih; Skor TOEFL\n")
	for j := 0; j <= len(daftar.nama)-del; j++ {
		fmt.Printf("| %d. %s; %s; %s; %s,%s-%s-%s; %s; %d\n", j+1, daftar.nama[j], daftar.usia[j], daftar.jenisK[j], daftar.ttl.tempat[j], daftar.ttl.tanggal[j], daftar.ttl.bulan[j], daftar.ttl.tahun[j], daftar.kerjaan[j], daftar.toefl[j])
	}
	fmt.Println("\n|  Silakan memilih menu di bawah ini:")
	fmt.Printf("\n|  1. Tambahkan data\n|  2. Edit data\n|  3. Hapus data\n|  4. Cari data\n|  5. Urutkan data\n|  6. Logout\n\n-> ")
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		pendaftaran()
	case 2:
		fmt.Printf("|  Masukkan nomor indeks atau nomor urut yang akan diedit berdasarkan data di atas\n-> ")
		fmt.Scanln(&pilih2)
		for (pilih2-1 > len(daftar.nama)-1) || (pilih2-1 < 0) || (daftar.nama[pilih2-1] == "") {
			fmt.Println("Mohon maaf data yang anda masukkan tidak ada")
			fmt.Scanln(&pilih2)
		}
		pilih3 := pilih2 - 1
		edit(&pilih3)
	case 3:
		fmt.Printf("|  Masukkan nomor indeks atau nomor urut yang akan dihapus berdasarkan data di atas\n-> ")
		fmt.Scanln(&pilih2)
		for (pilih2-1 > len(daftar.nama)-1) || (pilih2-1 < 0) || (daftar.nama[pilih2-1] == "") {
			fmt.Println("|  Mohon maaf data yang anda masukkan tidak ada\n -> ")
			fmt.Scanln(&pilih2)
		}
		terhapus := pilih2 - 1
		hapus(&terhapus)
	case 4:
		cari() //pencarian berdasarkan 1. nama saja, 2. usia saja,
	case 5:
		urut() //Urutkan berdasarkan 1.usia 2. Skor toefl
	case 6:
		welcome()
	default:
		akses_admin()
	}
}

func hapus(indeks *int) {
	if *indeks == 0 && len(daftar.toefl)-del == 0 {
		daftar.toefl[0] = -1
		if del <= len(daftar.toefl) {
			del++
		}
	} else {
		if *indeks <= len(daftar.nama)-del {
			for i := *indeks; i <= len(daftar.nama)-del; i++ {
				if i+1 <= len(daftar.nama)-del {
					daftar.nama[i] = daftar.nama[i+1]
					daftar.jenisK[i] = daftar.jenisK[i+1]
					daftar.usia[i] = daftar.usia[i+1]
					daftar.ttl.tempat[i] = daftar.ttl.tempat[i+1]
					daftar.ttl.tanggal[i] = daftar.ttl.tanggal[i+1]
					daftar.ttl.bulan[i] = daftar.ttl.bulan[i+1]
					daftar.ttl.tahun[i] = daftar.ttl.tahun[i+1]
					daftar.toefl[i] = daftar.toefl[i+1]
					daftar.kerjaan[i] = daftar.kerjaan[i+1]
				}
			}
			del++
		} else {
			fmt.Println("|  Nomor Indeks tidak ditemukan!")
		}
	}
	fmt.Println("|  Data berhasil dihapus!")
	menu_admin()
}

func urut() {
	var pilihan int
	fmt.Println(" _______________________________________")
	fmt.Println("|    Urutkan data berdasarkan:          |")
	fmt.Println("|---------------------------------------|")
	fmt.Println("|  1. Usia                              |")
	fmt.Println("|  2. Skor TOEFL                        |")
	fmt.Println("|_______________________________________|")
	fmt.Print("-> ")
	fmt.Scanln(&pilihan)
	var pilihan2 int
	switch pilihan {
	case 1:
		fmt.Println(" _______________________________________")
		fmt.Println("| 	  Urutkan secara:                    |")
		fmt.Println("|---------------------------------------|")
		fmt.Println("|  (1) Menaik                           |")
		fmt.Println("|  (2) Menurun                          |")
		fmt.Println("|_______________________________________|\n\n->")
		fmt.Scanln(&pilihan2)
		switch pilihan2 {
		case 1:
			urutbyusiaNaik()
		case 2:
			urutbyusiaTurun()
		default:
			urut()
		}
	case 2:
		fmt.Println(" _______________________________________")
		fmt.Println("| 	  Urutkan secara:                    |")
		fmt.Println("|---------------------------------------|")
		fmt.Println("|  (1) Menaik                           |")
		fmt.Println("|  (2) Menurun                          |")
		fmt.Println("|_______________________________________|\n\n->")
		fmt.Scanln(&pilihan2)
		switch pilihan2 {
		case 1:
			urutbytoeflNaik()
		case 2:
			urutbytoeflTurun()
		default:
			urut()
		}
		menu_admin()
	default:
		akses_admin()
	}
}

func urutbytoeflNaik() {
	var tampung [8]string
	var tampungtoefl int
	var urutan dataDiri
	for i := 0; i <= len(daftar.nama)-del; i++ {
		urutan.nama = append(urutan.nama, daftar.nama[i])
		urutan.jenisK = append(urutan.jenisK, daftar.jenisK[i])
		urutan.usia = append(urutan.usia, daftar.usia[i])
		urutan.ttl.tempat = append(urutan.ttl.tempat, daftar.ttl.tempat[i])
		urutan.ttl.tanggal = append(urutan.ttl.tanggal, daftar.ttl.tanggal[i])
		urutan.ttl.bulan = append(urutan.ttl.bulan, daftar.ttl.bulan[i])
		urutan.ttl.tahun = append(urutan.ttl.tahun, daftar.ttl.tahun[i])
		urutan.toefl = append(urutan.toefl, daftar.toefl[i])
		urutan.kerjaan = append(urutan.kerjaan, daftar.kerjaan[i])
	}
	i := 1
	for i < len(daftar.nama)-del+1 {
		tampungtoefl = urutan.toefl[i]
		tampung[0] = urutan.nama[i]
		tampung[1] = urutan.jenisK[i]
		tampung[2] = urutan.usia[i]
		tampung[3] = urutan.ttl.tempat[i]
		tampung[4] = urutan.ttl.tanggal[i]
		tampung[5] = urutan.ttl.bulan[i]
		tampung[6] = urutan.ttl.tahun[i]
		tampung[7] = urutan.kerjaan[i]
		j := i - 1
		for j >= 0 && urutan.toefl[j] > tampungtoefl {
			urutan.nama[j+1] = urutan.nama[j]
			urutan.jenisK[j+1] = urutan.jenisK[j]
			urutan.usia[j+1] = urutan.usia[j]
			urutan.ttl.tempat[j+1] = urutan.ttl.tempat[j]
			urutan.ttl.tanggal[j+1] = urutan.ttl.tanggal[j]
			urutan.ttl.bulan[j+1] = urutan.ttl.bulan[j]
			urutan.ttl.tahun[j+1] = urutan.ttl.tahun[j]
			urutan.kerjaan[j+1] = urutan.kerjaan[j]
			urutan.toefl[j+1] = urutan.toefl[j]
			j--
		}
		urutan.nama[j+1] = tampung[0]
		urutan.jenisK[j+1] = tampung[1]
		urutan.usia[j+1] = tampung[2]
		urutan.ttl.tempat[j+1] = tampung[3]
		urutan.ttl.tanggal[j+1] = tampung[4]
		urutan.ttl.bulan[j+1] = tampung[5]
		urutan.ttl.tahun[j+1] = tampung[6]
		urutan.kerjaan[j+1] = tampung[7]
		urutan.toefl[j+1] = tampungtoefl
		i++
	}
	for j := 0; j <= len(urutan.toefl)-1; j++ {
		fmt.Printf("%d. %s; %s; %s; %s,%s-%s-%s; %s; %d\n", j+1, urutan.nama[j], urutan.usia[j], urutan.jenisK[j], urutan.ttl.tempat[j], urutan.ttl.tanggal[j], urutan.ttl.bulan[j], urutan.ttl.tahun[j], urutan.kerjaan[j], urutan.toefl[j])
	}
}

func urutbytoeflTurun() {
	var tampung [8]string
	var tampungtoefl int
	var urutan dataDiri
	for i := 0; i <= len(daftar.nama)-del; i++ {
		urutan.nama = append(urutan.nama, daftar.nama[i])
		urutan.jenisK = append(urutan.jenisK, daftar.jenisK[i])
		urutan.usia = append(urutan.usia, daftar.usia[i])
		urutan.ttl.tempat = append(urutan.ttl.tempat, daftar.ttl.tempat[i])
		urutan.ttl.tanggal = append(urutan.ttl.tanggal, daftar.ttl.tanggal[i])
		urutan.ttl.bulan = append(urutan.ttl.bulan, daftar.ttl.bulan[i])
		urutan.ttl.tahun = append(urutan.ttl.tahun, daftar.ttl.tahun[i])
		urutan.toefl = append(urutan.toefl, daftar.toefl[i])
		urutan.kerjaan = append(urutan.kerjaan, daftar.kerjaan[i])
	}
	i := 1
	for i < len(daftar.nama)-del+1 {
		tampungtoefl = urutan.toefl[i]
		tampung[0] = urutan.nama[i]
		tampung[1] = urutan.jenisK[i]
		tampung[2] = urutan.usia[i]
		tampung[3] = urutan.ttl.tempat[i]
		tampung[4] = urutan.ttl.tanggal[i]
		tampung[5] = urutan.ttl.bulan[i]
		tampung[6] = urutan.ttl.tahun[i]
		tampung[7] = urutan.kerjaan[i]
		j := i - 1
		for j >= 0 && urutan.toefl[j] < tampungtoefl {
			urutan.nama[j+1] = urutan.nama[j]
			urutan.jenisK[j+1] = urutan.jenisK[j]
			urutan.usia[j+1] = urutan.usia[j]
			urutan.ttl.tempat[j+1] = urutan.ttl.tempat[j]
			urutan.ttl.tanggal[j+1] = urutan.ttl.tanggal[j]
			urutan.ttl.bulan[j+1] = urutan.ttl.bulan[j]
			urutan.ttl.tahun[j+1] = urutan.ttl.tahun[j]
			urutan.kerjaan[j+1] = urutan.kerjaan[j]
			urutan.toefl[j+1] = urutan.toefl[j]
			j--
		}
		urutan.nama[j+1] = tampung[0]
		urutan.jenisK[j+1] = tampung[1]
		urutan.usia[j+1] = tampung[2]
		urutan.ttl.tempat[j+1] = tampung[3]
		urutan.ttl.tanggal[j+1] = tampung[4]
		urutan.ttl.bulan[j+1] = tampung[5]
		urutan.ttl.tahun[j+1] = tampung[6]
		urutan.kerjaan[j+1] = tampung[7]
		urutan.toefl[j+1] = tampungtoefl
		i++
	}
	for j := 0; j <= len(urutan.toefl)-1; j++ {
		fmt.Printf("%d. %s; %s; %s; %s,%s-%s-%s; %s; %d\n", j+1, urutan.nama[j], urutan.usia[j], urutan.jenisK[j], urutan.ttl.tempat[j], urutan.ttl.tanggal[j], urutan.ttl.bulan[j], urutan.ttl.tahun[j], urutan.kerjaan[j], urutan.toefl[j])
	}
}

func urutbyusiaNaik() {
	var tampungtoefl int
	var urutusia dataDiri
	var nampungurut [9]string
	for i := 0; i <= len(daftar.nama)-del; i++ {
		urutusia.nama = append(urutusia.nama, daftar.nama[i])
		urutusia.jenisK = append(urutusia.jenisK, daftar.jenisK[i])
		urutusia.usia = append(urutusia.usia, daftar.usia[i])
		urutusia.ttl.tempat = append(urutusia.ttl.tempat, daftar.ttl.tempat[i])
		urutusia.ttl.tanggal = append(urutusia.ttl.tanggal, daftar.ttl.tanggal[i])
		urutusia.ttl.bulan = append(urutusia.ttl.bulan, daftar.ttl.bulan[i])
		urutusia.ttl.tahun = append(urutusia.ttl.tahun, daftar.ttl.tahun[i])
		urutusia.toefl = append(urutusia.toefl, daftar.toefl[i])
		urutusia.kerjaan = append(urutusia.kerjaan, daftar.kerjaan[i])
	}
	imin := 0
	min, _ := strconv.Atoi(urutusia.usia[0])
	for i := 0; i <= len(urutusia.usia)-1; i++ {
		for j := i; j <= len(urutusia.usia)-1; j++ {
			temp, _ := strconv.Atoi(urutusia.usia[j])
			if temp <= min {
				min, _ = strconv.Atoi(urutusia.usia[j])
				imin = j
			}
		}
		if imin != 0 {
			nampungurut[0] = urutusia.nama[imin]
			nampungurut[1] = urutusia.jenisK[imin]
			nampungurut[2] = urutusia.usia[imin]
			nampungurut[3] = urutusia.ttl.tempat[imin]
			nampungurut[4] = urutusia.ttl.tanggal[imin]
			nampungurut[5] = urutusia.ttl.bulan[imin]
			nampungurut[6] = urutusia.ttl.tahun[imin]
			tampungtoefl = urutusia.toefl[imin]
			nampungurut[7] = urutusia.kerjaan[imin]
			urutusia.nama[imin] = urutusia.nama[i]
			urutusia.jenisK[imin] = urutusia.jenisK[i]
			urutusia.usia[imin] = urutusia.usia[i]
			urutusia.ttl.tempat[imin] = urutusia.ttl.tempat[i]
			urutusia.ttl.tanggal[imin] = urutusia.ttl.tanggal[i]
			urutusia.ttl.bulan[imin] = urutusia.ttl.bulan[i]
			urutusia.ttl.tahun[imin] = urutusia.ttl.tahun[i]
			urutusia.toefl[imin] = urutusia.toefl[i]
			urutusia.kerjaan[imin] = urutusia.kerjaan[i]
			urutusia.nama[i] = nampungurut[0]
			urutusia.jenisK[i] = nampungurut[1]
			urutusia.usia[i] = nampungurut[2]
			urutusia.ttl.tempat[i] = nampungurut[3]
			urutusia.ttl.tanggal[i] = nampungurut[4]
			urutusia.ttl.bulan[i] = nampungurut[5]
			urutusia.ttl.tahun[i] = nampungurut[6]
			urutusia.toefl[i] = tampungtoefl
			urutusia.kerjaan[i] = nampungurut[7]
		}
	}
	for j := 0; j <= len(urutusia.nama)-1; j++ {
		fmt.Printf("%d. %s; %s; %s; %s,%s-%s-%s; %s; %d\n", j+1, urutusia.nama[j], urutusia.usia[j], urutusia.jenisK[j], urutusia.ttl.tempat[j], urutusia.ttl.tanggal[j], urutusia.ttl.bulan[j], urutusia.ttl.tahun[j], urutusia.kerjaan[j], urutusia.toefl[j])
	}
	menu_admin()
}

func urutbyusiaTurun() {
	var urutusia dataDiri
	var tampungtoefl int
	var nampungurut [9]string
	for i := 0; i <= len(daftar.nama)-del; i++ {
		urutusia.nama = append(urutusia.nama, daftar.nama[i])
		urutusia.jenisK = append(urutusia.jenisK, daftar.jenisK[i])
		urutusia.usia = append(urutusia.usia, daftar.usia[i])
		urutusia.ttl.tempat = append(urutusia.ttl.tempat, daftar.ttl.tempat[i])
		urutusia.ttl.tanggal = append(urutusia.ttl.tanggal, daftar.ttl.tanggal[i])
		urutusia.ttl.bulan = append(urutusia.ttl.bulan, daftar.ttl.bulan[i])
		urutusia.ttl.tahun = append(urutusia.ttl.tahun, daftar.ttl.tahun[i])
		urutusia.toefl = append(urutusia.toefl, daftar.toefl[i])
		urutusia.kerjaan = append(urutusia.kerjaan, daftar.kerjaan[i])
	}
	imin := 0
	max, _ := strconv.Atoi(urutusia.usia[0])
	for i := 0; i <= len(urutusia.usia)-1; i++ {
		for j := i; j <= len(urutusia.usia)-1; j++ {
			temp, _ := strconv.Atoi(urutusia.usia[j])
			if temp >= max {
				max, _ = strconv.Atoi(urutusia.usia[j])
				imin = j
			}
		}
		if imin != 0 {
			nampungurut[0] = urutusia.nama[imin]
			nampungurut[1] = urutusia.jenisK[imin]
			nampungurut[2] = urutusia.usia[imin]
			nampungurut[3] = urutusia.ttl.tempat[imin]
			nampungurut[4] = urutusia.ttl.tanggal[imin]
			nampungurut[5] = urutusia.ttl.bulan[imin]
			nampungurut[6] = urutusia.ttl.tahun[imin]
			tampungtoefl = urutusia.toefl[imin]
			nampungurut[7] = urutusia.kerjaan[imin]
			urutusia.nama[imin] = urutusia.nama[i]
			urutusia.jenisK[imin] = urutusia.jenisK[i]
			urutusia.usia[imin] = urutusia.usia[i]
			urutusia.ttl.tempat[imin] = urutusia.ttl.tempat[i]
			urutusia.ttl.tanggal[imin] = urutusia.ttl.tanggal[i]
			urutusia.ttl.bulan[imin] = urutusia.ttl.bulan[i]
			urutusia.ttl.tahun[imin] = urutusia.ttl.tahun[i]
			urutusia.toefl[imin] = urutusia.toefl[i]
			urutusia.kerjaan[imin] = urutusia.kerjaan[i]
			urutusia.nama[i] = nampungurut[0]
			urutusia.jenisK[i] = nampungurut[1]
			urutusia.usia[i] = nampungurut[2]
			urutusia.ttl.tempat[i] = nampungurut[3]
			urutusia.ttl.tanggal[i] = nampungurut[4]
			urutusia.ttl.bulan[i] = nampungurut[5]
			urutusia.ttl.tahun[i] = nampungurut[6]
			urutusia.toefl[i] = tampungtoefl
			urutusia.kerjaan[i] = nampungurut[7]
		}
	}
	for j := 0; j <= len(urutusia.nama)-1; j++ {
		fmt.Printf("%d. %s; %s; %s; %s,%s-%s-%s; %s; %d\n", j+1, urutusia.nama[j], urutusia.usia[j], urutusia.jenisK[j], urutusia.ttl.tempat[j], urutusia.ttl.tanggal[j], urutusia.ttl.bulan[j], urutusia.ttl.tahun[j], urutusia.kerjaan[j], urutusia.toefl[j])
	}
	menu_admin()
}
func edit(index *int) {
	type tabnama struct {
		lengkap []string
	}
	type tabtempatlahir struct {
		nama []string
	}
	var nama tabnama
	var tempatlahir tabtempatlahir
	var pilihan int
	fmt.Printf("\n ____________________________________________________________\n")
	fmt.Printf("|  Berikut ini tampilan data dari indeks yang anda masukkan, |\n")
	fmt.Println("|------------------------------------------------------------|")
	fmt.Printf("|    1. Nama : %s", daftar.nama[*index])
	fmt.Printf("\n|    2. Usia : %s", daftar.usia[*index])
	fmt.Printf("\n|    3. Jenis Kelamin : %s", daftar.jenisK[*index])
	fmt.Printf("\n|    4. Tempat lahir : %s", daftar.ttl.tempat[*index])
	fmt.Printf("\n|    5.  Tanggal lahir : %s-%s-%s", daftar.ttl.tanggal[*index], daftar.ttl.bulan[*index], daftar.ttl.tahun[*index])
	fmt.Printf("\n|    6. Skor TOEFL : %d", daftar.toefl[*index])
	fmt.Printf("\n|    7. Pekerjaan yang dipilih : %s", daftar.kerjaan[*index])
	fmt.Println("\n|____________________________________________________________|\n")
	fmt.Printf("\n\nMasukkan angka bagian yang akan diedit:\n")
	fmt.Scanln(&pilihan)
	var a, b, c, d, e string
	switch pilihan {
	case 1:
		fmt.Println("|  Nama :")
		fmt.Scanln(&a, &b, &c, &d, &e)
		if a != "" {
			nama.lengkap = append(nama.lengkap, a)
		}
		if b != "" {
			nama.lengkap = append(nama.lengkap, b)
		}
		if c != "" {
			nama.lengkap = append(nama.lengkap, c)
		}
		if d != "" {
			nama.lengkap = append(nama.lengkap, d)
		}
		if e != "" {
			nama.lengkap = append(nama.lengkap, e)
		}
		namafixed := strings.Join(nama.lengkap, " ")
		daftar.nama[*index] = namafixed
		akses_admin()
	case 2:
		fmt.Println("|  Usia: \n->")
		fmt.Scanln(&usia)
		daftar.usia[*index] = strconv.Itoa(usia)
		akses_admin()
	case 3:
		fmt.Printf("\n|  Jenis Kelamin\n|  1. %s\n|  2. %s\n", jk[0], jk[1])
		fmt.Scanln(&jenk)
		for jenk < 1 || jenk > 2 {
			fmt.Printf("|  Jenis Kelamin : %s", jenk)
			fmt.Printf("\n|  Mohon maaf, jenis kelamin yang anda masukkan tidak valid.")
			fmt.Printf("\n|  Jenis Kelamin\n|  1. %s\n|  2. %s\n-> ", jk[0], jk[1])
			fmt.Scanln(&jenk)
		}
		daftar.jenisK[*index] = jk[jenk-1]
		akses_admin()
	case 4:
		fmt.Print("|  Tempat lahir : ")
		var f, g, h, i, j string
		fmt.Scanln(&f, &g, &h, &i, &j)
		if f != "" {
			tempatlahir.nama = append(tempatlahir.nama, f)
		}
		if g != "" {
			tempatlahir.nama = append(tempatlahir.nama, g)
		}
		if h != "" {
			tempatlahir.nama = append(tempatlahir.nama, h)
		}
		if i != "" {
			tempatlahir.nama = append(tempatlahir.nama, i)
		}
		if j != "" {
			tempatlahir.nama = append(tempatlahir.nama, j)
		}
		tempatfixed := strings.Join(tempatlahir.nama, " ")
		daftar.ttl.tempat[*index] = tempatfixed
		akses_admin()
	case 5:
		fmt.Print("|  Tahun lahir (angka): ")
		fmt.Scanln(&thn)
		fmt.Print("|  Bulan lahir (angka): ")
		fmt.Scanln(&bln)
		fmt.Print("|  Tanggal lahir (angka): ")
		fmt.Scanln(&tgl)
		cek()
		daftar.ttl.tanggal[*index] = strconv.Itoa(tgl)
		daftar.ttl.tahun[*index] = strconv.Itoa(thn)
		daftar.ttl.bulan[*index] = strconv.Itoa(bln)
		akses_admin()
	case 6:
		fmt.Print("|  Skor TOEFL: ")
		fmt.Scanln(&stoefl)
		for stoefl < 0 || stoefl > 1000 {
			fmt.Println("|  WARNING!!!! Harap masukkan angka yang valid!")
			fmt.Scanln(&stoefl)
		}
		daftar.toefl[*index] = stoefl
		akses_admin()
	case 7:
		fmt.Printf("\n|  Masukkan angka sesuai dengan pekerjaan yang anda inginkan\n|  1. %s\n|  2. %s\n|  3. %s\n\n->", kerja[0], kerja[1], kerja[2])
		fmt.Scanln(&i)
		for i < 1 || i > 3 {
			fmt.Println("|  Harap masukkan angka yang valid!")
			fmt.Printf("\n|  Masukkan angka sesuai dengan pekerjaan yang anda inginkan\n|  1. %s\n|  2. %s\n|  3. %s\n-> ", kerja[0], kerja[1], kerja[2])
			fmt.Scanln(&i)
		}
		daftar.kerjaan[*index] = kerja[i-1]
		akses_admin()
	default:
		akses_admin()
	}
}

func menu_admin() {
	var pilih string
	fmt.Printf("\n ____________________________________________________________\n")
	fmt.Printf("|               SELAMAT DATANG DI LAMAN ADMIN                |\n")
	fmt.Printf("|           APLIKASI SELEKSI CALON PEGAWAI BARU!!            |\n")
	fmt.Println("|------------------------------------------------------------|")
	fmt.Println("|------------------------------------------------------------|")
	fmt.Println("|  Tekan tombol (1) untuk menambahkan data calon pegawai     |")
	fmt.Println("|  Tekan tombol (2) untuk melihat data calon pegawai         |")
	fmt.Printf("|------------------------------------------------------------|\n")
	fmt.Printf("|------------------------------------------------------------|\n")
	fmt.Printf("|  Tekan tombol sembarang untuk logout                       |\n")
	fmt.Printf("|____________________________________________________________|\n")
	fmt.Print("\n-> ")
	fmt.Scanf("%s\n", &pilih)
	switch pilih {
	case "1":
		pendaftaran()
	case "2":
		akses_admin()
	default:
		welcome()
	}
}

func login_guest() {
	var pilih string
	fmt.Printf("\n ____________________________________________________________\n")
	fmt.Printf("|             SELAMAT DATANG DI LAMAN LOGIN GUEST            |\n")
	fmt.Printf("|            APLIKASI SELEKSI CALON PEGAWAI BARU!!           |\n")
	fmt.Println("|------------------------------------------------------------|")
	fmt.Println("|  Tekan tombol 1 untuk mendaftar                            |")
	fmt.Println("|------------------------------------------------------------|")
	fmt.Printf("|  Tekan tombol 2 untuk kembali                              |\n")
	fmt.Println("|____________________________________________________________|\n")
	fmt.Print("-> ")
	fmt.Scanf("%s\n", &pilih)
	switch pilih {
	case "1":
		pendaftaran()
	case "2":
		welcome()
	default:
		welcome()
	}
}

func enkripsi(x string) bool {
	if x == "1301191171" || x == "1301190312" {
		return true
	} else {
		return false
	}
}

func pendaftaran() {
	type tabnama struct {
		lengkap []string
	}
	type tabtempatlahir struct {
		nama []string
	}
	var nama tabnama
	var tempatlahir tabtempatlahir
	fmt.Printf("\n ____________________________________________________________\n")
	fmt.Printf("|            SELAMAT DATANG DI LAMAN PENDAFTARAN             |\n")
	fmt.Printf("|           APLIKASI SELEKSI CALON PEGAWAI BARU!!            |\n")
	fmt.Println("|------------------------------------------------------------|")
	fmt.Println("| (*) Persyaratan usia untuk mendaftar: 18-61 tahun          |")
	fmt.Println("|------------------------------------------------------------|")
	var a, b, c, d, e string
	fmt.Print("| Nama : \n|  -> ")
	fmt.Scanln(&a, &b, &c, &d, &e)
	if a != "" {
		nama.lengkap = append(nama.lengkap, a)
	}
	if b != "" {
		nama.lengkap = append(nama.lengkap, b)
	}
	if c != "" {
		nama.lengkap = append(nama.lengkap, c)
	}
	if d != "" {
		nama.lengkap = append(nama.lengkap, d)
	}
	if e != "" {
		nama.lengkap = append(nama.lengkap, e)
	}
	namafixed := strings.Join(nama.lengkap, " ")
	namalengkap = namafixed
	fmt.Printf("\n| Jenis Kelamin\n| 1. %s\n| 2. %s\n|  -> ", jk[0], jk[1])
	fmt.Scanln(&jenk)
	for jenk < 1 || jenk > 2 {
		fmt.Printf("| Jenis Kelamin : %s", jenk)
		fmt.Printf("\n| Mohon maaf, jenis kelamin yang anda masukkan tidak valid.")
		fmt.Printf("\n| Jenis Kelamin\n1. %s\n2. %s\n", jk[0], jk[1])
		fmt.Scanln(&jenk)
	}
	fmt.Println("\n| TTL ")
	fmt.Print("| Tempat lahir : ")
	var f, g, h, k, l string
	fmt.Scanln(&f, &g, &h, &k, &l)
	if f != "" {
		tempatlahir.nama = append(tempatlahir.nama, f)
	}
	if g != "" {
		tempatlahir.nama = append(tempatlahir.nama, g)
	}
	if h != "" {
		tempatlahir.nama = append(tempatlahir.nama, h)
	}
	if k != "" {
		tempatlahir.nama = append(tempatlahir.nama, k)
	}
	if l != "" {
		tempatlahir.nama = append(tempatlahir.nama, l)
	}
	tempatfixed := strings.Join(tempatlahir.nama, " ")
	tempatlahirlengkap = tempatfixed
	fmt.Print("| Tahun lahir (angka): ")
	fmt.Scanln(&thn)
	fmt.Print("| Bulan lahir (angka): ")
	fmt.Scanln(&bln)
	fmt.Print("| Tanggal lahir (angka): ")
	fmt.Scanln(&tgl)
	cek()
	fmt.Print("\n| Skor TOEFL: \n|  -> ")
	fmt.Scanln(&stoefl)
	for stoefl < 0 || stoefl > 1000 {
		fmt.Println("| Harap masukkan angka yang valid!")
		fmt.Scanln(&stoefl)
	}
	fmt.Printf("\n| Masukkan angka sesuai dengan pekerjaan yang anda inginkan\n| 1. %s\n| 2. %s\n| 3. %s\n|  -> ", kerja[0], kerja[1], kerja[2])
	fmt.Scanf("%d", &i)
	for i < 1 || i > 3 {
		fmt.Println("| Harap masukkan angka yang valid!")
		fmt.Printf("\n| Masukkan angka sesuai dengan pekerjaan yang anda inginkan\n| 1. %s\n| 2. %s\n| 3. %s\n|  -> ", kerja[0], kerja[1], kerja[2])
		fmt.Scanln(&i)
	}
	fmt.Printf("|____________________________________________________________|\n")
	view()
}

func carinama() {
	type tabnama struct {
		lengkap [5]string
	}
	var nama tabnama
	fmt.Println("\n|  Tulis nama yang akan anda cari")
	fmt.Scanln(&nama.lengkap[0], &nama.lengkap[1], &nama.lengkap[2], &nama.lengkap[3], &nama.lengkap[4])
	fmt.Println("\n|  Berikut data yang anda cari:")
	fmt.Printf("\n|  Nomor Indeks; Nama; Usia; Jenis Kelamin; Tempat dan Tanggal Lahir; Lowongan Kerja yang Dipilih; Skor TOEFL\n")
	cekada := 0
	for i := 0; i <= 4; i++ {
		if nama.lengkap[i] != "" {
			nama.lengkap[i] = strings.ToLower(nama.lengkap[i])
			for j := 0; j <= len(daftar.nama)-del; j++ {
				hurufkecilsemua := strings.ToLower(daftar.nama[j])
				carinama := strings.Split(hurufkecilsemua, " ")
				for k := 0; k <= len(carinama)-del; k++ {
					if nama.lengkap[i] == carinama[k] {
						fmt.Printf("%d. %s; %s; %s; %s,%s-%s-%s; %s; %d\n", j+1, daftar.nama[j], daftar.usia[j], daftar.jenisK[j], daftar.ttl.tempat[j], daftar.ttl.tanggal[j], daftar.ttl.bulan[j], daftar.ttl.tahun[j], daftar.kerjaan[j], daftar.toefl[j])
						k = len(carinama) + 9999
						cekada++
					}

				}
			}
		} else {
			i = 9999
		}
	}
	if cekada <= 0 {
		fmt.Println("|  Data tidak ditemukan!")
	}
}

func cariusia() {
	fmt.Println("|  Tulis usia yang akan anda cari")
	fmt.Scanln(&usia)
	usia_cari := strconv.Itoa(usia)
	countingusia := 0
	fmt.Println("|  Berikut data yang anda cari:")
	fmt.Printf("|  Nomor Indeks; Nama; Usia; Jenis Kelamin; Tempat dan Tanggal Lahir; Lowongan Kerja yang Dipilih; Skor TOEFL\n")
	for j := 0; j <= len(daftar.usia)-del; j++ {
		if usia_cari == daftar.usia[j] {
			fmt.Printf("|  %d. %s; %s; %s; %s,%s-%s-%s; %s; %d\n", j+1, daftar.nama[j], daftar.usia[j], daftar.jenisK[j], daftar.ttl.tempat[j], daftar.ttl.tanggal[j], daftar.ttl.bulan[j], daftar.ttl.tahun[j], daftar.kerjaan[j], daftar.toefl[j])
			countingusia++
		}
	}
	if countingusia <= 0 {
		fmt.Printf("|  Tidak ada data yang berusia %s\n", usia_cari)
	}
}

func cari() {
	var pencarian int
	fmt.Println()
	fmt.Println("|  Apa yang ingin anda cari?")
	fmt.Println("|  1. Nama")
	fmt.Println("|  2. Usia")
	fmt.Scanln(&pencarian)
	switch pencarian {
	case 1:
		carinama()
	case 2:
		cariusia()
	default:
		akses_admin()
	}
	menu_admin()
}

func cek() {
	c := 0
	usia = 2019 - thn
	if (usia < 18) || (usia > 60) {
		c++
		fmt.Printf("\n|  Usia : %d", usia)
		fmt.Printf("\n|  Mohon maaf, usia yang anda masukkan belum memenuhi kriteria kami")
	}
	if (thn < 1958) || (thn > 2001) {
		c++
		fmt.Printf("\n|  Tanggal lahir : %d-%d-%d", tgl, bln, thn)
		fmt.Printf("\n|  Mohon maaf, tahun lahir yang anda masukkan belum memenuhi kriteria kami")
	}
	if (bln < 1) || (bln > 12) {
		c++
		fmt.Printf("\n|  Tanggal lahir : %d-%d-%d", tgl, bln, thn)
		fmt.Printf("\n|  Mohon maaf, bulan lahir yang anda masukkan tidak valid, pastikan ditik dalam rantang 1 s.d. 12")
	}
	if bln == 2 {
		if thn%4 == 0 {
			if tgl < 1 || tgl > 29 {
				c++
				fmt.Printf("\n|  Tanggal lahir : %d-%d-%d", tgl, bln, thn)
				fmt.Printf("\n|  Mohon maaf, tanggal lahir yang anda masukkan tidak valid")
			}
		} else {
			if tgl < 1 || tgl > 28 {
				c++
				fmt.Printf("\n|  Tanggal lahir : %d-%d-%d", tgl, bln, thn)
				fmt.Printf("\n|  Mohon maaf, tanggal lahir yang anda masukkan tidak valid")
			}
		}

	} else {
		if tgl < 1 || tgl > 31 {
			c++
			fmt.Printf("\n|  Tanggal lahir : %d-%d-%d", tgl, bln, thn)
			fmt.Printf("\n|  Mohon maaf, tanggal lahir yang anda masukkan tidak valid")
		}
	}
	for c > 0 {
		c = 0
		fmt.Printf("\n\n|  Harap masukkan kembali data yang valid dan memenuhi kriteria usia\n")
		fmt.Print("|  Tahun lahir (angka): ")
		fmt.Scanln(&thn)
		fmt.Print("|  Bulan lahir (angka): ")
		fmt.Scanln(&bln)
		fmt.Print("|  Tanggal lahir (angka): ")
		fmt.Scanln(&tgl)
		cek()
	}
}

func simpan() {
	var pilih int
	var cek2 bool
	cek := false

	if len(daftar.nama) > 0 {
		for a := 0; a <= len(daftar.nama)-del; a++ {
			cek = (stoefl == daftar.toefl[a]) && (namalengkap == daftar.nama[a]) && (strconv.Itoa(usia) == daftar.usia[a]) && (jk[jenk-1] == daftar.jenisK[a]) && (tempatlahirlengkap == daftar.ttl.tempat[a]) && (strconv.Itoa(tgl) == daftar.ttl.tanggal[a]) && (strconv.Itoa(bln) == daftar.ttl.bulan[a]) && (strconv.Itoa(thn) == daftar.ttl.tahun[a]) && (kerja[i-1] == daftar.kerjaan[a])
			if cek == true {
				cek2 = true
			}
		}
	}
	if cek2 {
		fmt.Println("|  Mohon maaf data yang anda masukkan sudah terdaftar sebelumnya!")
	} else {
		if len(daftar.toefl) <= 0 {
			daftar.nama = append(daftar.nama, namalengkap)
			usia = 2019 - thn
			daftar.usia = append(daftar.usia, strconv.Itoa(usia))
			daftar.ttl.tempat = append(daftar.ttl.tempat, tempatlahirlengkap)
			daftar.ttl.tanggal = append(daftar.ttl.tanggal, strconv.Itoa(tgl))
			daftar.ttl.tahun = append(daftar.ttl.tahun, strconv.Itoa(thn))
			daftar.ttl.bulan = append(daftar.ttl.bulan, strconv.Itoa(bln))
			daftar.kerjaan = append(daftar.kerjaan, kerja[i-1])
			daftar.jenisK = append(daftar.jenisK, jk[jenk-1])
			daftar.toefl = append(daftar.toefl, stoefl)
		} else {
			if len(daftar.toefl) <= del && daftar.toefl[0] == -1 {
				daftar.nama[0] = namalengkap
				usia = 2019 - thn
				daftar.usia[0] = strconv.Itoa(usia)
				daftar.ttl.tempat[0] = tempatlahirlengkap
				daftar.ttl.tanggal[0] = strconv.Itoa(tgl)
				daftar.ttl.tahun[0] = strconv.Itoa(thn)
				daftar.ttl.bulan[0] = strconv.Itoa(bln)
				daftar.kerjaan[0] = kerja[i-1]
				daftar.jenisK[0] = jk[jenk-1]
				daftar.toefl[0] = stoefl
				if del > len(daftar.toefl) {
					del = del - 1
				}
			} else {
				daftar.nama = append(daftar.nama, namalengkap)
				usia = 2019 - thn
				daftar.usia = append(daftar.usia, strconv.Itoa(usia))
				daftar.ttl.tempat = append(daftar.ttl.tempat, tempatlahirlengkap)
				daftar.ttl.tanggal = append(daftar.ttl.tanggal, strconv.Itoa(tgl))
				daftar.ttl.tahun = append(daftar.ttl.tahun, strconv.Itoa(thn))
				daftar.ttl.bulan = append(daftar.ttl.bulan, strconv.Itoa(bln))
				daftar.kerjaan = append(daftar.kerjaan, kerja[i-1])
				daftar.jenisK = append(daftar.jenisK, jk[jenk-1])
				daftar.toefl = append(daftar.toefl, stoefl)
			}
		}

		fmt.Printf("\n\n _____________________________________________________________________")
		fmt.Printf("\n|   WARNING!!! Selamat, data yang anda masukkan berhasil terdaftar!   |")
		fmt.Printf("\n|---------------------------------------------------------------------|")
	}
	fmt.Printf("\n|  Apa yang akan anda lakukan selanjutnya?                            |\n")
	fmt.Printf("|     (1) Ke menu admin                                               |\n")
	fmt.Printf("|     (2) Daftar lagi                                                 |\n")
	fmt.Printf("|     (3) Kembali ke menu utama                                       |\n")
	fmt.Printf("|_____________________________________________________________________|\n-> ")
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		login_admin()
	case 2:
		pendaftaran()
	case 3:
		welcome()
	default:
		view()
	}
}

func view() {
	var eof int
	fmt.Printf("\n\n\n ________________________________________________________________\n")
	fmt.Printf("|               SELAMAT!!! DATA BERHASIL DIREKAM                 |\n")
	fmt.Printf("|       Berikut ini ulasan dari data yang anda masukkan,         |\n")
	fmt.Println("|----------------------------------------------------------------|")
	fmt.Printf("|  Nama : %s", namalengkap)
	usia = 2019 - thn
	fmt.Printf("\n|  Usia : %d\n", usia)
	fmt.Printf("|  Jenis Kelamin : %s", jk[jenk-1])
	fmt.Printf("\n|  Tempat lahir : %s", tempatlahirlengkap)
	fmt.Printf("\n|  Tanggal lahir : %d-%d-%d", tgl, bln, thn)
	fmt.Printf("\n|  Skor TOEFL : %d", stoefl)
	fmt.Printf("\n|  Pekerjaan yang dipilih : %s\n", kerja[i-1])
	fmt.Println("|----------------------------------------------------------------|")
	fmt.Printf("| Apa anda sudah yakin dengan data di atas?                      |\n")
	fmt.Print("| (1) Ya, daftarkan!                                             |\n")
	fmt.Print("| (2) Ulangi pendaftaran                                         |\n")
	fmt.Print("| (3) Kembali ke menu utama                                      |\n")
	fmt.Printf("|________________________________________________________________|\n\n -> ")
	fmt.Scanln(&eof)
	switch eof {
	case 1:
		simpan()
	case 2:
		pendaftaran()
	case 3:
		welcome()
	default:
		view()
	}
}
