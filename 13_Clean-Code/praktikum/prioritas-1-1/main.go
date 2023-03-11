package main

/*
1. bukan termasuk kesalahan clean code, namun agar struct/method dapat digunakan di luar file, huruf awal dari struct/method/func harus menggunakan huruf kapital
*/
type user struct {
	id       int
	username int
	password int
}
/*
1. penamaan yang tidak sesuai prinsip clean code, dapat diganti menggunakan penamaan secara snake atau camel case (user_service / userService)

2. nama field yang ambigu/kurang jelas dalam menyampaikan fungsi (t), dapat diubah ke penamaan yang lebih bermakna
*/
type userservice struct {
	t []user
}

/*
1. penamaan method yang tidak sesuai prinsip clean code, dapat diganti menggunakan penamaan secara snake atau camel case (get_all_user / getAllUser)
*/
func (u userservice) getallusers() []user {
	return u.t
}

/*
1. penamaan method
2. penamaan variable kurang sesuai dengan prinsip clean code (r), dapat diganti menjadi user/userObj
3. perilaku/behaviour method saat tidak menemukan id dapat diperbaiki, seperti menampilkan pesan error / mengirimkan 2 return type, user dan error (jika ada) (khusus bahasa golang)
*/
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
