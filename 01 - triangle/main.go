package main

import "fmt"

func triangle(n int) {
	/**
	if n = 5 
	output :	1
				22
				333
				4444
				55555
	Algoritma :
	baris pertama 
		for pertama i = 1, syarat i <= 5 (true),
		masuk ke for ke dua j = 1, syarat j <= i (true),
		cetak i(1), j menjadi 2 (j++), syarat menjadi false lalu keluar dari for ke 2
		keluar for ke 2 , cetak baris baru, i = 2 (i++)


	baris kedua 
		for pertama i=2, syarat i <=5 (true)
		masuk ke for ke dua j = 1, syarat j <= i (true)
		cetak i(2), j = 2 (j++), syarat true j = i 
		cetak i(2) output jadi 22, j = 3 (j++), syarat false.
		keluar for ke 2 , cetak baris baru, i = 3 (i++)

	baris ke tiga 
		for pertama i=3, syarat i <=5 (true)
		masuk ke for ke dua j = 1, syarat j <= i (true)
		cetak i(3), j = 2 (j++), syarat true j < i 
		cetak i(3) output jadi 33, j = 3 (j++), syarat true j = i .
		cetak i(3) output jadi 333, j = 4 (j++), syarat false.
		keluar for ke 2 , cetak baris baru, i = 4 (i++)

	dst...
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
	// sama dgn code di atas
	fmt.Println("=== ~~~ ====")
	for i := 1; i <= n;{
		for j := 1; j <= i; {
			fmt.Print(i)
			 j++
		}
		fmt.Println()
		 i++ 
	}

		/**
	if n = 5 
	output :	
     1
    22
   333
  4444
 55555
	*/
	fmt.Println("=== ~~~ ====")
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}

	/**
	if n = 5 
	output :	0 1 2 3 4
				1 2 3 4
				2 3 4
				3 4
				4
	*/
	fmt.Println("=== Code 2 ====")
	for i := 0; i < n; i++ {

		for j := i; j < n; j++ {
			fmt.Print(j, " ")
		}

    	fmt.Println()
	}

	/**
	if n = 5 
	output :	1
				21
				321
				4321
				54321
	*/
	fmt.Println("=== Code 3 ====")
	for i := 1; i <= n; i++ {

		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}

    	fmt.Println()
	}

	/**
	if n = 5 
	output :	1
				23
				454
				3212
				34543
	*/
	
	fmt.Println("=== Code 4 ====")
	k := 1
	isN := true
	for i := 1; i <= n; i++ {

		for j := i; j >= 1; j-- {
			fmt.Print(k)
			if isN {
				k++
			} else {
				k--
			}
			if k >= n {
				isN = false
			} else if k <= 1 {
				isN = true
			}

		}

    	fmt.Println()
	}


}

func main() {
triangle(5)
}