// Case 3: Book Store

package main

import "fmt"

// Sebuah transaksi di toko buku akan disimpan ke dalam database.
// Untuk itu, perlu dibuat model dari transaksi menggunakan struct.
// Setiap transaksi memiliki beberapa field, yaitu:
// - Total: total harga yang perlu dibayar
// - Date: tanggal terjadinya transaksi berupa string
// - Books: buku-buku yang dibeli oleh pembeli

// Setiap buku juga memiliki beberapa field, yaitu:
// - Title: judul buku
// - Writer: nama penulis buku
// - Price: harga buku
// - Count: banyak buku yang dibeli oleh pembeli

// Pada tanggal 23-01-2021, terjadi transaksi dengan total harga sebesar
// 32000. Pembeli membeli 2 buah buku. Buku pertama berjudul "Laut Bercerita"
// karangan Leila S. Chudori seharga 108000 sebanyak 1 buah. Buku berikutnya
// berjudul "Bumi" karangan Tere Liye seharga 97000 sebanyak 1 buah.

type Books struct{
	Title string;
	Writer string;
	Price float32;
	Count int;
}
type Transaction struct{
	Total float32;
	Date string;
	Books []Books;
}

func main() {
	// Modifikasi pendefinisian variable tx1 sehingga sesuai dengan kasus
	// yang diceritakan.
	tx1 := Transaction{
		Total: 32000,
		Date: "23-01-2021",
		Books: []Books{{
			Title: "Laut Bercerita", Writer: "Writer", Price: 108000, Count: 1}, {Title: "Bumi", Writer: "Tere Liye", Price: 97000, Count: 1}},
	}

	fmt.Println(tx1)
}
