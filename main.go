package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const NMAX = 100

type Mood struct {
	Tanggal      string
	TugasTerkait string
	Skor         int
	Deskripsi    string
}

type Tugas struct {
	Tanggal   string
	Nama      string
	Prioritas int
	Durasi    int
	Selesai   bool
}

var dataMood [NMAX]Mood
var jumMood int

var dataTugas [NMAX]Tugas
var jumTugas int

var pembaca *bufio.Scanner

func bacaStr() string {
	pembaca.Scan()
	return strings.TrimSpace(pembaca.Text())
}

func bacaInt() int {
	pembaca.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(pembaca.Text()))
	return n
}

func potong(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}
	return s
}

func tampilHeaderBesar() {
	fmt.Print("\033[97m")
	fmt.Println("  █████╗ ███████╗██╗███████╗████████╗███████╗███╗   ██╗    ██╗   ██╗██╗██████╗ ████████╗██╗   ██╗ █████╗ ██╗")
	fmt.Println(" ██╔══██╗██╔════╝██║██╔════╝╚══██╔══╝██╔════╝████╗  ██║    ██║   ██║██║██╔══██╗╚══██╔══╝██║   ██║██╔══██╗██║")
	fmt.Println(" ███████║███████╗██║███████╗   ██║   █████╗  ██╔██╗ ██║    ██║   ██║██║██████╔╝   ██║   ██║   ██║███████║██║")
	fmt.Println(" ██╔══██║╚════██║██║╚════██║   ██║   ██╔══╝  ██║╚██╗██║    ╚██╗ ██╔╝██║██╔══██╗   ██║   ██║   ██║██╔══██║██║")
	fmt.Println(" ██║  ██║███████║██║███████║   ██║   ███████╗██║ ╚████║     ╚████╔╝ ██║██║  ██║   ██║   ╚██████╔╝██║  ██║███████╗")
	fmt.Println(" ╚═╝  ╚═╝╚══════╝╚═╝╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═══╝      ╚═══╝  ╚═╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝")
	fmt.Println()
	fmt.Println("                   ██╗  ██╗███████╗███████╗███████╗██╗  ██╗ █████╗ ████████╗ █████╗ ███╗   ██╗")
	fmt.Println("                   ██║ ██╔╝██╔════╝██╔════╝██╔════╝██║  ██║██╔══██╗╚══██╔══╝██╔══██╗████╗  ██║")
	fmt.Println("                   █████╔╝ █████╗  ███████╗█████╗  ███████║███████║   ██║   ███████║██╔██╗ ██║")
	fmt.Println("                   ██╔═██╗ ██╔══╝  ╚════██║██╔══╝  ██╔══██║██╔══██║   ██║   ██╔══██║██║╚██╗██║")
	fmt.Println("                   ██║  ██╗███████╗███████║███████╗██║  ██║██║  ██║   ██║   ██║  ██║██║ ╚████║")
	fmt.Println("                   ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝")
	fmt.Print("\033[0m\n")
}

func tambahMood() {
	if jumMood >= NMAX {
		fmt.Println(">> Memori catatan mood penuh.")
		return
	}
	fmt.Print("Tanggal (DD-MM-YYYY)  : ")
	tgl := bacaStr()
	fmt.Print("Tugas terkait         : ")
	tg := bacaStr()
	fmt.Print("Skor emosi (1 - 10)   : ")
	sk := bacaInt()
	fmt.Print("Deskripsi perasaan    : ")
	ds := bacaStr()

	dataMood[jumMood].Tanggal = tgl
	dataMood[jumMood].TugasTerkait = tg
	dataMood[jumMood].Skor = sk
	dataMood[jumMood].Deskripsi = ds
	jumMood = jumMood + 1

	fmt.Println(">> Catatan suasana hati berhasil ditambahkan.")
}

func ubahMood() {
	tampilMood()
	if jumMood == 0 {
		return
	}
	fmt.Print("Nomor catatan yang ingin diubah: ")
	no := bacaInt() - 1
	if no < 0 || no >= jumMood {
		fmt.Println(">> Nomor tidak valid.")
		return
	}
	fmt.Print("Tanggal baru (DD-MM-YYYY) : ")
	dataMood[no].Tanggal = bacaStr()
	fmt.Print("Tugas terkait baru        : ")
	dataMood[no].TugasTerkait = bacaStr()
	fmt.Print("Skor baru                 : ")
	dataMood[no].Skor = bacaInt()
	fmt.Print("Deskripsi baru            : ")
	dataMood[no].Deskripsi = bacaStr()
	fmt.Println(">> Catatan berhasil diubah.")
}

func hapusMood() {
	tampilMood()
	if jumMood == 0 {
		return
	}
	fmt.Print("Nomor catatan yang ingin dihapus: ")
	no := bacaInt() - 1
	if no < 0 || no >= jumMood {
		fmt.Println(">> Nomor tidak valid.")
		return
	}
	for i := no; i < jumMood-1; i++ {
		dataMood[i] = dataMood[i+1]
	}
	jumMood = jumMood - 1
	fmt.Println(">> Catatan berhasil dihapus.")
}

func tambahTugas() {
	if jumTugas >= NMAX {
		fmt.Println(">> Memori daftar tugas penuh.")
		return
	}
	fmt.Print("Tanggal (DD-MM-YYYY)        : ")
	tg := bacaStr()
	fmt.Print("Nama tugas                  : ")
	nm := bacaStr()
	fmt.Print("Prioritas (1 - 5)           : ")
	pr := bacaInt()
	fmt.Print("Durasi pengerjaan (menit)   : ")
	dr := bacaInt()
	fmt.Print("Status (1=selesai, 0=belum) : ")
	st := bacaInt()

	dataTugas[jumTugas].Tanggal = tg
	dataTugas[jumTugas].Nama = nm
	dataTugas[jumTugas].Prioritas = pr
	dataTugas[jumTugas].Durasi = dr
	if st == 1 {
		dataTugas[jumTugas].Selesai = true
	} else {
		dataTugas[jumTugas].Selesai = false
	}
	jumTugas = jumTugas + 1

	fmt.Println(">> Tugas berhasil ditambahkan.")
}

func ubahTugas() {
	tampilTugas()
	if jumTugas == 0 {
		return
	}
	fmt.Print("Nomor tugas yang ingin diubah: ")
	no := bacaInt() - 1
	if no < 0 || no >= jumTugas {
		fmt.Println(">> Nomor tidak valid.")
		return
	}
	fmt.Print("Tanggal baru (DD-MM-YYYY) : ")
	dataTugas[no].Tanggal = bacaStr()
	fmt.Print("Nama tugas baru           : ")
	dataTugas[no].Nama = bacaStr()
	fmt.Print("Prioritas                 : ")
	dataTugas[no].Prioritas = bacaInt()
	fmt.Print("Durasi (menit)            : ")
	dataTugas[no].Durasi = bacaInt()
	fmt.Print("Status (1/0)              : ")
	st := bacaInt()
	if st == 1 {
		dataTugas[no].Selesai = true
	} else {
		dataTugas[no].Selesai = false
	}
	fmt.Println(">> Tugas berhasil diubah.")
}

func hapusTugas() {
	tampilTugas()
	if jumTugas == 0 {
		return
	}
	fmt.Print("Nomor tugas yang ingin dihapus: ")
	no := bacaInt() - 1
	if no < 0 || no >= jumTugas {
		fmt.Println(">> Nomor tidak valid.")
		return
	}
	for i := no; i < jumTugas-1; i++ {
		dataTugas[i] = dataTugas[i+1]
	}
	jumTugas = jumTugas - 1
	fmt.Println(">> Tugas berhasil dihapus.")
}

func tampilMood() {
	tampilHeaderBesar()
	fmt.Println("\n=== Daftar Catatan Suasana Hati ===")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf(" %-4s | %-12s | %-30s | %-4s | %-80s\n", "No", "Tanggal", "Tugas Terkait", "Skor", "Deskripsi")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")
	if jumMood == 0 {
		fmt.Println(" (belum ada data)")
	} else {
		for i := 0; i < jumMood; i++ {
			fmt.Printf(" %-4d | %-12s | %-30s | %-4d | %-80s\n",
				i+1, dataMood[i].Tanggal, potong(dataMood[i].TugasTerkait, 30), dataMood[i].Skor, potong(dataMood[i].Deskripsi, 80))
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------")
}

func tampilTugas() {
	tampilHeaderBesar()
	fmt.Println("\n=== Daftar Tugas Harian ===")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf(" %-4s | %-12s | %-65s | %-9s | %-6s | %-8s\n", "No", "Tanggal", "Nama Tugas", "Prioritas", "Durasi", "Status")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
	if jumTugas == 0 {
		fmt.Println(" (belum ada data)")
	} else {
		for i := 0; i < jumTugas; i++ {
			stat := "Belum"
			if dataTugas[i].Selesai {
				stat = "Selesai"
			}
			fmt.Printf(" %-4d | %-12s | %-65s | %-9d | %-6d | %-8s\n",
				i+1, dataTugas[i].Tanggal, potong(dataTugas[i].Nama, 65), dataTugas[i].Prioritas, dataTugas[i].Durasi, stat)
		}
	}
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------")
}

func cariTugasSequential(kata string) {
	fmt.Println("\n===========================================")
	fmt.Println("| HASIL PENCARIAN TUGAS (Sequential Search)|")
	fmt.Println("===========================================")
	ketemu := 0
	kataKecil := strings.ToLower(kata)
	for i := 0; i < jumTugas; i++ {
		cocok := false
		if dataTugas[i].Tanggal == kata {
			cocok = true
		}
		if strings.Contains(strings.ToLower(dataTugas[i].Nama), kataKecil) {
			cocok = true
		}
		if cocok {
			fmt.Printf(" - [%d] %s | %s | Prioritas %d | %d menit\n",
				i+1, dataTugas[i].Tanggal, dataTugas[i].Nama,
				dataTugas[i].Prioritas, dataTugas[i].Durasi)
			ketemu = ketemu + 1
		}
	}
	if ketemu == 0 {
		fmt.Println("(tidak ada hasil yang cocok)")
	} else {
		fmt.Printf("Total ditemukan: %d data\n", ketemu)
	}
}

func cariMoodSequential(kata string) {
	fmt.Println("\n==========================================")
	fmt.Println("| HASIL PENCARIAN MOOD (Sequential Search)|")
	fmt.Println("==========================================")
	ketemu := 0
	kataKecil := strings.ToLower(kata)
	for i := 0; i < jumMood; i++ {
		cocok := false
		if dataMood[i].Tanggal == kata {
			cocok = true
		}
		if strings.Contains(strings.ToLower(dataMood[i].Deskripsi), kataKecil) || strings.Contains(strings.ToLower(dataMood[i].TugasTerkait), kataKecil) {
			cocok = true
		}
		if cocok {
			fmt.Printf(" - [%d] %s | Tugas: %s | Skor %d | %s\n",
				i+1, dataMood[i].Tanggal, dataMood[i].TugasTerkait,
				dataMood[i].Skor, dataMood[i].Deskripsi)
			ketemu = ketemu + 1
		}
	}
	if ketemu == 0 {
		fmt.Println("(tidak ada hasil yang cocok)")
	} else {
		fmt.Printf("Total ditemukan: %d data\n", ketemu)
	}
}

func urutTugasTanggal() {
	for pass := 1; pass < jumTugas; pass++ {
		idx := pass - 1
		for i := pass; i < jumTugas; i++ {
			if dataTugas[i].Tanggal < dataTugas[idx].Tanggal {
				idx = i
			}
		}
		dataTugas[pass-1], dataTugas[idx] = dataTugas[idx], dataTugas[pass-1]
	}
}

func urutMoodTanggal() {
	for pass := 1; pass < jumMood; pass++ {
		idx := pass - 1
		for i := pass; i < jumMood; i++ {
			if dataMood[i].Tanggal < dataMood[idx].Tanggal {
				idx = i
			}
		}
		dataMood[pass-1], dataMood[idx] = dataMood[idx], dataMood[pass-1]
	}
}

func cariTugasBinary(tgl string) {
	urutTugasTanggal()
	low := 0
	high := jumTugas - 1
	idx := -1
	for low <= high {
		mid := low + (high-low)/2
		if dataTugas[mid].Tanggal == tgl {
			idx = mid
			break
		} else if dataTugas[mid].Tanggal < tgl {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("\n=======================================")
	fmt.Println("| HASIL PENCARIAN TUGAS (Binary Search)|")
	fmt.Println("=======================================")
	if idx == -1 {
		fmt.Println("Tugas dengan tanggal", tgl, "tidak ditemukan.")
		return
	}
	stat := "Belum"
	if dataTugas[idx].Selesai {
		stat = "Selesai"
	}
	fmt.Printf(" Tanggal   : %s\n", dataTugas[idx].Tanggal)
	fmt.Printf(" Nama      : %s\n", dataTugas[idx].Nama)
	fmt.Printf(" Prioritas : %d\n", dataTugas[idx].Prioritas)
	fmt.Printf(" Durasi    : %d menit\n", dataTugas[idx].Durasi)
	fmt.Printf(" Status    : %s\n", stat)
}

func cariMoodBinary(tgl string) {
	urutMoodTanggal()
	low := 0
	high := jumMood - 1
	idx := -1
	for low <= high {
		mid := low + (high-low)/2
		if dataMood[mid].Tanggal == tgl {
			idx = mid
			break
		} else if dataMood[mid].Tanggal < tgl {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("\n======================================")
	fmt.Println("| HASIL PENCARIAN MOOD (Binary Search)|")
	fmt.Println("======================================")
	if idx == -1 {
		fmt.Println("Catatan mood dengan tanggal", tgl, "tidak ditemukan.")
		return
	}
	fmt.Printf(" Tanggal       : %s\n", dataMood[idx].Tanggal)
	fmt.Printf(" Tugas Terkait : %s\n", dataMood[idx].TugasTerkait)
	fmt.Printf(" Skor          : %d\n", dataMood[idx].Skor)
	fmt.Printf(" Deskripsi     : %s\n", dataMood[idx].Deskripsi)
}

func selectionSortPrioritas() {
	for pass := 1; pass < jumTugas; pass++ {
		idx := pass - 1
		for i := pass; i < jumTugas; i++ {
			if dataTugas[i].Prioritas > dataTugas[idx].Prioritas {
				idx = i
			}
		}
		dataTugas[pass-1], dataTugas[idx] = dataTugas[idx], dataTugas[pass-1]
	}
	fmt.Println("\n===================================================================")
	fmt.Println("| Tugas terurut berdasarkan PRIORITAS (Selection Sort - Descending)|")
	fmt.Println("===================================================================")
	tampilTugas()
}

func insertionSortDurasi() {
	for i := 1; i < jumTugas; i++ {
		temp := dataTugas[i]
		j := i - 1
		for j >= 0 && dataTugas[j].Durasi > temp.Durasi {
			dataTugas[j+1] = dataTugas[j]
			j = j - 1
		}
		dataTugas[j+1] = temp
	}
	fmt.Println("\n===============================================================")
	fmt.Println("| Tugas terurut berdasarkan DURASI (Insertion Sort - Ascending)|")
	fmt.Println("===============================================================")
	tampilTugas()
}

func statistikMood() {
	fmt.Println("\n============================================")
	fmt.Println("|   STATISTIK TREN SUASANA HATI MINGGUAN   |")
	fmt.Println("============================================")
	if jumMood == 0 {
		fmt.Println("(belum ada data mood untuk dihitung)")
		return
	}

	total := 0
	tertinggi := dataMood[0].Skor
	terendah := dataMood[0].Skor
	for i := 0; i < jumMood; i++ {
		total = total + dataMood[i].Skor
		if dataMood[i].Skor > tertinggi {
			tertinggi = dataMood[i].Skor
		}
		if dataMood[i].Skor < terendah {
			terendah = dataMood[i].Skor
		}
	}
	rata := float64(total) / float64(jumMood)

	fmt.Println("-------------------------------------")
	fmt.Printf(" Total catatan        | %-8d\n", jumMood)
	fmt.Printf(" Rata-rata skor       | %-8.2f\n", rata)
	fmt.Printf(" Skor tertinggi       | %-8d\n", tertinggi)
	fmt.Printf(" Skor terendah        | %-8d\n", terendah)
	fmt.Println("-------------------------------------")
	fmt.Println()

	urutMoodTanggal()
	fmt.Println("=== Diagram batang harian (urut tanggal) ===")
	fmt.Println("--------------------------------------------------")
	fmt.Printf(" %-12s | %-20s | %-4s\n", "Tanggal", "Bar", "Skor")
	fmt.Println("--------------------------------------------------")
	for i := 0; i < jumMood; i++ {
		bar := ""
		for k := 0; k < dataMood[i].Skor; k++ {
			bar = bar + "#"
		}
		for len(bar) < 19 {
			bar = bar + " "
		}
		fmt.Printf(" %-12s | %-20s | %-4d\n",
			dataMood[i].Tanggal, bar, dataMood[i].Skor)
	}
	fmt.Println("--------------------------------------------------")
}

func statistikTugas() {
	fmt.Println("\n===============================================")
	fmt.Println("| STATISTIK TINGKAT PENYELESAIAN TUGAS HARIAN |")
	fmt.Println("===============================================")
	if jumTugas == 0 {
		fmt.Println("(belum ada data tugas untuk dihitung)")
		return
	}
	selesai := 0
	totalDurasi := 0
	for i := 0; i < jumTugas; i++ {
		if dataTugas[i].Selesai {
			selesai = selesai + 1
		}
		totalDurasi = totalDurasi + dataTugas[i].Durasi
	}
	belum := jumTugas - selesai
	persen := float64(selesai) * 100.0 / float64(jumTugas)
	rataDurasi := float64(totalDurasi) / float64(jumTugas)

	fmt.Println("-------------------------------------------")
	fmt.Printf(" Total tugas            | %-8d\n", jumTugas)
	fmt.Printf(" Sudah selesai          | %-8d\n", selesai)
	fmt.Printf(" Belum selesai          | %-8d\n", belum)
	fmt.Printf(" Tingkat penyelesaian   | %-7.1f%%\n", persen)
	fmt.Printf(" Rata-rata durasi       | %-5.1f mnt\n", rataDurasi)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	fmt.Print("Progress: [")
	bar := int(persen) / 5
	for k := 0; k < 20; k++ {
		if k < bar {
			fmt.Print("#")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Printf("] %.1f%%\n", persen)
}

func tampilMenuUtama() {
	fmt.Println()
	tampilHeaderBesar()
	fmt.Println("=======================================================================================================")
	fmt.Println("| Halo bagaimana kabarmu saya siap membantu kamu dalam kesehatan mental dan produktivitas!            |")
	fmt.Println("=======================================================================================================")
	fmt.Println("| 1. Kelola Catatan Suasana Hati                                                                      |")
	fmt.Println("| 2. Kelola Daftar Tugas                                                                              |")
	fmt.Println("| 3. Pencarian Data                                                                                   |")
	fmt.Println("| 4. Pengurutan Tugas                                                                                 |")
	fmt.Println("| 5. Statistik                                                                                        |")
	fmt.Println("| 0. Keluar                                                                                           |")
	fmt.Println("=======================================================================================================")
	fmt.Print(">> Pilih menu: ")
}

func tampilMenuMood() {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("|  MENU CATATAN SUASANA HATI   |")
	fmt.Println("================================")
	fmt.Println("| 1. Tambah catatan            |")
	fmt.Println("| 2. Ubah catatan              |")
	fmt.Println("| 3. Hapus catatan             |")
	fmt.Println("| 4. Lihat semua catatan       |")
	fmt.Println("| 0. Kembali                   |")
	fmt.Println("================================")
	fmt.Print(">> Pilih: ")
}

func tampilMenuTugas() {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("|      MENU DAFTAR TUGAS       |")
	fmt.Println("================================")
	fmt.Println("| 1. Tambah tugas              |")
	fmt.Println("| 2. Ubah tugas                |")
	fmt.Println("| 3. Hapus tugas               |")
	fmt.Println("| 4. Lihat semua tugas         |")
	fmt.Println("| 0. Kembali                   |")
	fmt.Println("================================")
	fmt.Print(">> Pilih: ")
}

func tampilMenuCari() {
	fmt.Println()
	fmt.Println("=================================================")
	fmt.Println("|              MENU PENCARIAN DATA              |")
	fmt.Println("=================================================")
	fmt.Println("| 1. Cari tugas (Sequential - kata/tanggal)     |")
	fmt.Println("| 2. Cari tugas (Binary - tanggal)              |")
	fmt.Println("| 3. Cari mood  (Sequential - kata/tanggal)     |")
	fmt.Println("| 4. Cari mood  (Binary - tanggal)              |")
	fmt.Println("| 0. Kembali                                    |")
	fmt.Println("=================================================")
	fmt.Print(">> Pilih: ")
}

func tampilMenuUrut() {
	fmt.Println()
	fmt.Println("=================================================")
	fmt.Println("|             MENU PENGURUTAN TUGAS             |")
	fmt.Println("=================================================")
	fmt.Println("| 1. Urut berdasarkan Prioritas (Selection Sort)|")
	fmt.Println("| 2. Urut berdasarkan Durasi    (Insertion Sort)|")
	fmt.Println("| 0. Kembali                                    |")
	fmt.Println("=================================================")
	fmt.Print(">> Pilih: ")
}

func tampilMenuStat() {
	fmt.Println()
	fmt.Println("==================================")
	fmt.Println("|         MENU STATISTIK         |")
	fmt.Println("==================================")
	fmt.Println("| 1. Tren suasana hati mingguan  |")
	fmt.Println("| 2. Tingkat penyelesaian tugas  |")
	fmt.Println("| 0. Kembali                     |")
	fmt.Println("==================================")
	fmt.Print(">> Pilih: ")
}

func subMenuMood() {
	for {
		tampilMenuMood()
		p := bacaInt()
		if p == 0 {
			return
		}
		switch p {
		case 1:
			tambahMood()
		case 2:
			ubahMood()
		case 3:
			hapusMood()
		case 4:
			tampilMood()
		default:
			fmt.Println(">> Pilihan tidak tersedia.")
		}
	}
}

func subMenuTugas() {
	for {
		tampilMenuTugas()
		p := bacaInt()
		if p == 0 {
			return
		}
		switch p {
		case 1:
			tambahTugas()
		case 2:
			ubahTugas()
		case 3:
			hapusTugas()
		case 4:
			tampilTugas()
		default:
			fmt.Println(">> Pilihan tidak tersedia.")
		}
	}
}

func subMenuCari() {
	for {
		tampilMenuCari()
		p := bacaInt()
		if p == 0 {
			return
		}
		switch p {
		case 1:
			fmt.Print("Masukkan kata kunci / tanggal (DD-MM-YYYY): ")
			k := bacaStr()
			cariTugasSequential(k)
		case 2:
			fmt.Print("Masukkan tanggal (DD-MM-YYYY): ")
			k := bacaStr()
			cariTugasBinary(k)
		case 3:
			fmt.Print("Masukkan kata kunci / tanggal (DD-MM-YYYY): ")
			k := bacaStr()
			cariMoodSequential(k)
		case 4:
			fmt.Print("Masukkan tanggal (DD-MM-YYYY): ")
			k := bacaStr()
			cariMoodBinary(k)
		default:
			fmt.Println(">> Pilihan tidak tersedia.")
		}
	}
}

func subMenuUrut() {
	for {
		tampilMenuUrut()
		p := bacaInt()
		if p == 0 {
			return
		}
		switch p {
		case 1:
			selectionSortPrioritas()
		case 2:
			insertionSortDurasi()
		default:
			fmt.Println(">> Pilihan tidak tersedia.")
		}
	}
}

func subMenuStat() {
	for {
		tampilMenuStat()
		p := bacaInt()
		if p == 0 {
			return
		}
		switch p {
		case 1:
			statistikMood()
		case 2:
			statistikTugas()
		default:
			fmt.Println(">> Pilihan tidak tersedia.")
		}
	}
}

func main() {
	kesalahan := godotenv.Load()
	if kesalahan != nil {
		fmt.Println("File .env tidak ditemukan atau gagal dimuat.")
		return
	}

	pembaca = bufio.NewScanner(os.Stdin)
	pembaca.Buffer(make([]byte, 1024*1024), 1024*1024)

	jumMood = 0
	jumTugas = 0

	for {
		tampilMenuUtama()
		pilih := bacaInt()
		switch pilih {
		case 1:
			subMenuMood()
		case 2:
			subMenuTugas()
		case 3:
			subMenuCari()
		case 4:
			subMenuUrut()
		case 5:
			subMenuStat()
		case 0:
			fmt.Println("\nTerima kasih!")
			return
		default:
			fmt.Println(">> Pilihan tidak tersedia.")
		}
	}
}
