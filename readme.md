## <span style="color:cyan">Go Routine</span> 🔥
1. main function itu go routine. Berarti se simple hello world aja itu udah ada 1 go routine.
2. go routines adalah thread yang berukuran kecil yang di kelola oleh go runtime. Ukuran 1 go routine sekitar 2 kB.
3. Apa itu thread, thread adalah bagian dari sistem operasi yang bertanggung jawab untuk menjalankan aplikasi.
4. Apa itu go runtime? go runtime itu library yang digunakan oleh program golang saat sedang dijalankan.
5. Pada baris 10 di file main.go, go routine ```printSomething``` akan berjalan secara concurrency dengan function ```main```. Namun, ketika dijalankan tulisan `One!` tidak akan muncul yang muncul hanya `Two!`, karena tulisan `One!` sudah dijalankan di belakang layar oleh go runtime.