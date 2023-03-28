# Echo
salah satu framework yang dapat digunakan untuk membuat http server.

## Echo advantage
- optimized router
- data rendering
- data binding
- middleware support
- easily scalable

(note: tidak menggunakan db driver/orm secara default)

## Instalasi
```bash
go get github.com/labstack/echo/v4
```

## Routing dan Controller
```go
// inisialisasi echo app
func main() {
    e := echo.New()

    // define route
    // method juga tersedia untuk POST, PUT, DELETE
    e.GET("/route/url", getController)
}

// controller yang menghandle req
func getController(e echo.Context) error {
    ...
}
```

## Data rendering
```go
func getController(e echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}
```

## Retrieve dan binding data
```go
func main() {
    c.Param("param") // mengambil parameter dari url
    c.QueryParam("query-param") // mengambil value dari query param di url(setelah ?)
    c.FormValue("form-param") // mengambil value dari form param
    c.Bind(&objToBind) // menempelkan value yang dikirimkan ke obj yang dimasukkan
}
```