# DNA Sequence Matcher

**DNA Sequence Matcher** is a web application than can
be used to match a DNA sequence submitted by users with
a database of disease-identifying nucleotide sequences.
If a segment of the submitted DNA matches the database, 
the DNA sample will be flagged as positive, and vice versa. 
A history of app usage will be stored in the database.

Users' inputs are sanitised using regex. The sequence 
matcher implements the Knutt-Morris-Pratt algorithm and 
the Boyer-Moore algorithm.

## Setup
1. Pertama, server dinyalakan dengan cara masuk ke directory backend lalu ketik command “go run main.go”
2. Lalu  pindah ke directory “new-frontend” dan nyalakan frontend dengan command “npm start”
3. Secara  default, homepage menampilkan bagian cek kesehatan
4. Pada bagian sidebar terdapat tombol-tombol yang dapat ditekan untuk pindah ke page-page fitur lain

## Note
TBD
