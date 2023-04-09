package model


type Billing struct {
    tanggal      string
    Waktumulai   string
    Waktuselesai string
}

// func createBilling(c echo.Context) error {
//     b := new(Billing)
//     if err := c.Bind(b); err != nil {
//         return err
//     }

//     // lakukan validasi pada inputan yang dimasukkan

//     // simpan data ke database di BillingServer

//     return c.JSON(http.StatusCreated, b)
// }

// func getBilling(c echo.Context) error {
//     // ambil data billing dari database di BillingServer

//     return c.JSON(http.StatusOK, billing)
// }

// func updateBilling(c echo.Context) error {
//     // lakukan validasi pada inputan yang dimasukkan

//     // update data di database di BillingServer

//     return c.NoContent(http.StatusNoContent)
// }
