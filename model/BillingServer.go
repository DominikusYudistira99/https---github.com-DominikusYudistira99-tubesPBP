package model

import "time"

type BillingServer struct {
    NomorBilling string
    NamaUser     string
    Tarif        float64
    WaktuMulai   time.Time
    Waktu        time.Duration
    TotalHarga   float64
}

// func (bs *BillingServer) BillingKlien() {
//     // melakukan proses billing pada klien berdasarkan data billing pada instance bs
// }

// func (bs *BillingServer) MenampilkanData() {
//     // menampilkan data billing pada instance bs
// }

// func (bs *BillingServer) MenampilkanLaporan() {
//     // menampilkan laporan billing pada instance bs
// }

// func (bs *BillingServer) MembacaDataBilling(nomorBilling string) error {
//     // membaca data billing dari sistem berdasarkan nomor billing
//     // mengisi nilai atribut pada instance bs sesuai dengan data billing yang dibaca
//     // mengembalikan error jika terjadi kesalahan dalam membaca data billing
// }
