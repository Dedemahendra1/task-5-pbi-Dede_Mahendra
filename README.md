# task-5-pbi-Dede_Mahendra

<h2>HINT TASK 5</h2


<p>Pada task akhir VIX Full Stack Developer ini kalian diarahkan untuk membentuk API berdasarkan kasus yang telah diberikan. Pada kasus ini, kalian diinstruksikan untuk membuat API untuk mengupload dan menghapus gambar. API yang kalian bentuk adalah POST, GET, PUT, dan DELETE.</p>



<span>Ketentuan API :</span>

<span>Pada bagian User Endpoint :</span>

POST : /users/register, dan gunakan atribut berikut ini
ID (primary key, required)
Username (required)
Email (unique & required) 
Password (required & minlength 6)
Relasi dengan model Photo (Gunakan constraint cascade)
Created At (timestamp)
Updated At (timestamp)
GET: /users/login
Using email & password (required)
PUT : /users/:userId (Update User)
DELETE : /users/:userId (Delete User)
Photos Endpoint

POST : /photos 
ID
Title
Caption
PhotoUrl
UserID
Relasi dengan model User
GET : /photos
PUT : /photoId
DELETE : /:photoId


Requirement :
Authorization dapat menggunakan tool Go JWT (https://github.com/dgrijalva/jwt-go) 
Pastikan hanya user yang membuat foto yang dapat menghapus / mengubah foto


Environment
Struktur dokumen / environment dari GoLang yang akan dibentuk kurang lebih sebagai berikut :

app :Menampung pembuatan struct dalam kasus ini menggunakan struct user untuk keperluan data dan authentication
controllers : Berisi antara logic database yaitu models dan query
database: Berisi konfigurasi database serta digunakan untuk menjalankan koneksi database dan migration
helpers : Berisi fungsi-fungsi yang dapat digunakan di setiap tempat dalam hal ini jwt, bcrypt, headerValue
middlewares :Berisi fungsi yang digunakan untuk proses otentikasi jwt yang digunakan untuk proteksi api
models : Berisi models yang digunakan untuk relasi database 
router : Berisi konfigurasi routing / endpoint yang akan digunakan untuk mengakses api
go mod : Yang digunakan untuk manajemen package / dependency berupa library


Tools :
Tools yang dapat kalian gunakan : 

Gin Gonic Framework : https://github.com/gin-gonic/gin 
Gorm : https://gorm.io/index.html 
JWT Go : https://github.com/dgrijalva/jwt-go 
Go Validator : http://github.com/asaskevich/govalidator 


Untuk database, gunakanlah server SQL open source seperti MySQL, PostgreSQL, atau Microsoft SQL Server.



Format Submit :
Upload hasil code kalian dalam bentuk link repository Github.
Buka website Github
login atau sign up Github menggunakan email yang kalian gunakan
Klik tombol profil di pojok kanan atas dan pilih ‘My repositories’ dan klik New
Namakan repository tersebut dengan format : ‘task-5-vix-btpns-[NAMA]’
Pastikan repository tersebut adalah Public lalu klik Create Repository
Upload seluruh folder environment pada repository tersebut.
Submit link kalian didalam dokumen notepad dengan format nama ‘vix-btpns-fd-[NAMA]’ yang berisikan link repository github yang telah kalian kerjakan!
