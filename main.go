package main

import "fmt"

// Create Book type
type Book struct {
	Name        string
	Author      string
	Publisher   string
	PublishDate int
	Price       float64
}

// Create Slice that contains Book type
var bookList = []Book{
	{Name: "Harry Potter ve Felsefe Taşı", Author: "J.K. Rowling", Publisher: "Yapı Kredi Yayınları", PublishDate: 2018, Price: 32.00},
	{Name: "Yüzüklerin Efendisi Yüzük Kardeşliği", Author: "J.R.R. Tolkien", Publisher: "Metis Yayınları", PublishDate: 2020, Price: 62.54},
	{Name: "The Return of the King: The Lord of the Rings", Author: "J.R.R. Tolkien", Publisher: "Sarya Kitabevi", PublishDate: 1987, Price: 139.00},
	{Name: "Açlık", Author: "Knut Hamsun", Publisher: "Varlık Yayınları", PublishDate: 2016, Price: 16.50},
	{Name: "Satranç", Author: "Stefan Zweig", Publisher: "İş Bankası Kültür Yayınları", PublishDate: 2021, Price: 6.30},
}

func main() {
	showBookList()
}

// It returns the book list
func showBookList(){
	turkishLira:=string(rune(8378))
	for _, v:= range bookList{
		book:=fmt.Sprintf("\tAdı: %s\n \tYazarı: %s\n \tYayıncı: %s\n \tYayın Tarihi: %v\n \tFiyat: %.2f %v\n",v.Name,v.Author,v.Publisher,v.PublishDate,v.Price,turkishLira)
		fmt.Println()
		fmt.Println(book)
	}
}
