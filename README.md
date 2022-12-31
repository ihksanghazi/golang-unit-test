# golang-unit-test

belajar unit test di golang

Aturan File Test

- untuk membuat file unit test, harus menggunakan akhiran \_test
  Contoh : hello_world_test.go

Aturan Function Unit Test

- nama function untuk unit test harus diawali dengan nama Test
- harus memiliki parameter (t \*testing.T) dan tidak mengembalikan return value

Menjalankan unit Test

- go test (run semua testing)
- go test -v (run testing dengan menampilkan nama function nya)
- go test -v -run namafunction (untuk menjalankan testing tertentu)
- go test ./... (run semua test walaupun di luar package)
