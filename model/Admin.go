package model


type Admin struct {
    Username string
    Password string
}

func (a *Admin) Melihat() {
    // Method untuk melihat data
}

func (a *Admin) Mengatur() {
    // Method untuk mengatur data
}

func (a *Admin) Mengontrol() {
    // Method untuk mengontrol akses
}

func (a *Admin) Mengubah() {
    // Method untuk mengubah data
}

func (a *Admin) Menghapus() {
    // Method untuk menghapus data
}

func (a *Admin) Mengisi() {
    // Method untuk mengisi username dan password
}
