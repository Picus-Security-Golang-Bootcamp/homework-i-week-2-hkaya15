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
	for _, v := range bookList {
		fmt.Println(v)
	}
}
