package main

import "fmt"


func login() {
	var nama, peran string

	fmt.Println("===================================")
	fmt.Println("1. Pengguna login sebagai")
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Sebagai (dokter/pasien): ")
	fmt.Scanln(&peran)

	if peran == "dokter" {
		fmt.Printf("Selamat datang, Dr. %s. Anda masuk sebagai Dokter.\n", nama)
		fiturDokter(nama)
	} else if peran == "pasien" {
		fmt.Printf("Selamat datang, %s. Anda masuk sebagai Pasien.\n", nama)
	} else {
		fmt.Println("Peran tidak valid. Silakan coba lagi.")
	}
}

func fiturDokter(namaDokter string) {
	var pilihan string

	for {
		fmt.Println("===================================")
		fmt.Println("Fitur Dokter:")
		fmt.Println("1. Memberi Tanggapan Postingan Pertanyaan")
		fmt.Println("2. Mencari Pertanyaan Berdasarkan Tag")
		fmt.Println("3. Melihat Forum Postingan Pertanyaan")
		fmt.Println("4. Keluar")
		fmt.Println("===================================")
		fmt.Print("Pilih fitur (1-4): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			var nomorPertanyaan, tanggapan string
			fmt.Print("Jawab Pertanyaan Nomor Berapa: ")
			fmt.Scanln(&nomorPertanyaan)
			fmt.Print("Jawaban: ")
			fmt.Scanln(&tanggapan)
			fmt.Printf("Dr. %s telah menjawab pertanyaan nomor %s dengan jawaban: %s\n", namaDokter, nomorPertanyaan, tanggapan)
		case "2":
			var tag string
			fmt.Print("Cari pertanyaan berdasarkan tag: ")
			fmt.Scanln(&tag)
			fmt.Printf("Menampilkan pertanyaan dengan tag '%s'\n", tag)
		case "3":
			fmt.Println("Menampilkan forum postingan pertanyaan")
		case "4":
			fmt.Println("Keluar dari fitur Dokter.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}


func main() {
	var menuChoice string

	fmt.Println("===================================")
	fmt.Println("Tampilan Menu Utama :")
	fmt.Println("1. Login")
	fmt.Println("2. Forum Pertanyaan")
	fmt.Println("3. Keluar")
	fmt.Println("===================================")
	fmt.Print("Pilih menu (1-3): ")
	fmt.Scanln(&menuChoice)

	switch menuChoice {
	case "1":
		login()
	case "2":
		fmt.Println("Fitur Forum Pertanyaan belum tersedia.")
	case "3":
		fmt.Println("Terima kasih, program selesai.")
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
	}
}

