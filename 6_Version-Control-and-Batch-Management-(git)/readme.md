# Version Control and Branch Management (Git)

## version control history
1. single user based : scss, RCS
2. centralized based : CVS, Perforce, etc
3. distributed based : git, mercurial, bazaar

## Apa itu git
git merupakan tools untuk memanagement perubahan pada file/code yang kita miliki. dibuat oleh linus torvald pada 2005. distibuted sehingga dapat digunakan secara lokal maupun untuk server

## Git repository
letak folder/file yang akan ditrack perubahan oleh git menggunakan file history yang dibuat saat inisialisasi git repo tersebut

## Commit
titik perubahan yang telah dilakukan dan akan disimpan sebagai riwayat perubahan oleh git

jika terdapat kesalahan karena perubahan di titik ini dapat di 'undo' ke titik sebelumnya

## git hosting service
dapat menggunakan github sebagai server git repository. terdapat banyak project open source yang ada di github.

## common command/feature list
- git init : inisialisasi repo
- git status : menunjukan status file2 yang ada di repo
- git add [nama-file] : mengubah stage file menjadi staged
- git log : menunjukkan histori commit
- git diff : menunjukkan perbedaan antara commit
- git branch, checkout : membuat branch/cabang dari repo, berguna untuk melakukan test feature baru, pengerjaan fitur secara paralel, dsb
- git push, pull, fetch : mengambil/mengupload commit dari remote repo ke local repo/sebaliknya
- git stash : menyimpan perubahan secara sementara ke stash

## conflict di git
- jika terdapat perubahan yang terjadi di line yang sama, akan terjadi conflict
- conflict dapat diselesaikan dengan 

## git workflow
- best practice adalah dengan membuat banyak branch tergantung fitur yang sedang dikembangkan
- jika fitur telah selesai, merge ke branch development terlebih dulu
- jangan langsung merge/commit ke branch dev atau bahkan ke production
- jika sudah yakin integrasi seluruh fitur telah berhasil, baru merge ke branch production