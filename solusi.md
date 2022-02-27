# Solusi Project Challenge TokoGambar

## Masalah & Solusi yang Ditawarkan
1. Pada pengecekan kesamaan gambar, menggunakan persamaan hasil hash gambar (menggunakan = secara langsung). Menurut saya, hal tersebut kurang efektif. Karena ada kemungkinan perbedaan hasil hash gambar antara 2 gambar yang memiliki persamaan. Oleh karena itu, saya mengimplementasikan levenshtein distance untuk pengecekan kesamaan hasil hash gambar. Dengan catatan bahwa gambar dinyatakan serupa (similar) jika levenshtein distance antara kedua hasil hash gambar <=5 (berdasarkan [Perceptual Hash](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)).

2. Hash gambar pada project ini sebelumnya menggunakan SHA256. Berdasarkan observasi, algoritma SHA256 kurang efektif dalam menemukan gambar-gambar yang serupa. Sebagai contoh pada kasus gambar `input_3.jpg` seharusnya mengeluarkan hasil yang sama dengan `input_1.jpg` karena keduanya identik secara visual. Oleh karena itu, saya mengimplementasikan perceptual hash sebagai pengganti SHA256.

3. Project ini pada mulanya menggunakan flat architecture, sehingga business logic, model, api, dsb terletak pada root directory project. Selanjutnya saya implementasikan hexagonal architecture supaya project lebih terstruktur dan memudahkan pengerjaan bila terdapat pengembangan lebih lanjut.

4. Project ini telah di-deploy menggunakan heroku container-based pada [https://sheltered-eyrie-92338.herokuapp.com/](https://sheltered-eyrie-92338.herokuapp.com/). Daftar API yang digunakan dapat dilihat pada [http_api.md](./http_api.md)