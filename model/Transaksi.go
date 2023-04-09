package model


type Transaksi struct {
    NomorBilling string
    NamaUser     string
    Tarif        int
    Biaya        int
}

func (t *Transaksi) Dilihat() {
    // Method untuk melihat transaksi
}

func (t *Transaksi) Dibayar() {
    // Method untuk melakukan pembayaran transaksi
}
