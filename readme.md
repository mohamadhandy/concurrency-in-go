# <span style="color:cyan">Go Routine</span> 🔥
1. main function itu go routine. Berarti se simple hello world aja itu udah ada 1 go routine.
2. go routines adalah thread yang berukuran kecil yang di kelola oleh go runtime. Ukuran 1 go routine sekitar 2 kB.
3. Apa itu thread, thread adalah bagian dari sistem operasi yang bertanggung jawab untuk menjalankan aplikasi.
4. Apa itu go runtime? go runtime itu library yang digunakan oleh program golang saat sedang dijalankan.
5. Pada baris 10 di file main.go di branch `01-goroutines`, go routine ```printSomething``` akan berjalan secara concurrency dengan function ```main```. Namun, ketika dijalankan tulisan `One!` tidak akan muncul yang muncul hanya `Two!`, karena tulisan `One!` sudah dijalankan di belakang layar oleh go runtime.
## <span style="color:cyan">Wait Groups</span> 🔥
1. Ada 3 cara untuk menangani concurrency pada kode golang. Yang pertama `wait group`.
2. Pertama harus di set terlebih dahulu waitGroupnya dengan integer menggunakan keyword `wg.Add()` kemudian isi angka `9` menggunakan `len(words)`. Jadi dia harus nunggu sebanyak 9 hal sampai selesai (sampai 0). Baris 30 `wg.Wait()` itu dia nunggu sampai jadi 0 (sampai selesai). 
3. Cara dia sampai 0 menggunakan keyword `defer wg.Done()` yang artinya dia ngurangin 1 sampai ke 0, kemudian blok fungsi `printSomething` di akhiri. 
4. aturan dari golang, `wait group` harus menggunakan pointer karena ketika kita sudah membuat `wait group` ya jangan di copy, di passing aja menggunakan pointer. 
5. Wait group itu tidak boleh dibawah dari 0, kalau dia dibawah dari 0 maka muncul error `negative waitGroup Counter`.