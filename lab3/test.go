package main

import (
  "fmt"
 
  "lab3/filesystem"
)

func main(){
	fmt.Println("Virtual File System Test -- testing file system operations")
	fmt.Println("------------------------------------------------------------")

	vfs := filesystem.NewVirtualFileSystem()

	// Dodanie katalogu Documents (dir)
	docsDir := filesystem.NewDirectory("Documents", "/Documents")
	if err := vfs.AddItem(docsDir); err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Added Documents directory")
	}

	// Dodanie pliku hello.txt do katalogu Documents (file)
	file1 := filesystem.NewFile("hello.txt", "/Documents/hello.txt")
	if err := vfs.AddItem(file1); err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Added hello.txt to Documents directory")
	}

	// Powtórzenie dodania pliku hello.txt do katalogu Documents (file)
	if err := vfs.AddItem(file1); err != nil {
		fmt.Println("\n Adding file again error:", err)
	} else {
		fmt.Println("\n ❌ Added hello.txt to Documents directory again")
	}

	// Dodanie pliku czesc.txt do katalogu Dokumenty (file) 
	file2 := filesystem.NewFile("czesc.txt", "/Dokumenty/czesc.txt")
	if err := vfs.AddItem(file2); err != nil {
		fmt.Println("\n Adding file to non-existing directory error:", err)
	} else {
		fmt.Println("\n ❌Added czesc.txt to Dokumenty  directory")
	}

	// Dodanie pliku readme.txt do katalogu Documents (readonly)
	readOnlyFile := filesystem.NewReadOnlyFile("readme.txt", "/Documents/readme.txt")
	if err := vfs.AddItem(readOnlyFile); err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Added readme.txt to Documents directory")
	}

	// Dodanie pliku hello-link.txt do katalogu root (symlink)
	targetItem, err := vfs.FindItem("/Documents/hello.txt")
	if err != nil {
		fmt.Println("\n ❌ Could not find target for symlink:", err)
	} else {
		link := filesystem.NewSymLink("hello-link.txt", "/hello-link.txt", targetItem)
		if err := vfs.AddItem(link); err != nil {
			fmt.Println("\n ❌", err)
		} else {
			fmt.Println("\n Added hello-link.txt symlink to Documents/hello.txt")
		}
	}

	// Zapis do pliku hello.txt
	_, err = file1.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Wrote to hello.txt")
	}

	// Zapis do pliku tylko do odczytu
	_, err = readOnlyFile.Write([]byte("Should fail"))
	if err != nil {
		fmt.Println("\n Writing to read-only file error: ", err)
	} else {
		fmt.Println("\n ❌ Wrote to read-only file")
	}

	// Odczyt z pliku hello.txt
	buffer := make([]byte, 100)
	n, err := file1.Read(buffer)
	if err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Read", n,"bytes from", file1.Name(), ":", string(buffer[:n]))
	}

	// Sprawdzenie czy plik istnieje
	item, err := vfs.FindItem("/Documents/hello.txt")
	if err != nil {
			fmt.Println("\n ❌", err)
	} else {
			fmt.Println("\n Found file:", item.Name())
	}

	// Sprawdzenie czy katalog istnieje
	item, err = vfs.FindItem("/Documents")
	if err != nil {
			fmt.Println("\n ❌", err)
	} else {
			fmt.Println("\n Found directory:", item.Name())
	}

	// Sprawdzenie czy symlink działa
	item, err = vfs.FindItem("/hello-link.txt")
	if err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Found symlink target:", item.Name())
	}
	
	// Usunięcie pliku
	if err := docsDir.RemoveItem("hello.txt"); err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Removed hello.txt from Documents directory")
	}

	// Sprawdzenie czy plik został usunięty
	_, err = vfs.FindItem("/Documents/hello.txt")
	if err != nil {
		fmt.Println("\n File not found after removing:", err)
	} else {
		fmt.Println("\n ❌ File still exists after deletion")
	}

	// Usunięcie katalogu
	if err := vfs.RemoveItem("Documents"); err != nil {
		fmt.Println("\n ❌", err)
	} else {
		fmt.Println("\n Removed Documents directory")
	}

	// Sprawdzenie czy katalog został usunięty
	_, err = vfs.FindItem("/Documents")
	if err != nil {
		fmt.Println("\n Directory not found after removing:", err)
	} else {
		fmt.Println("\n ❌ Directory still exists after deletion")
	}

	// Usunięcie nieistniejącego pliku
	if err := vfs.RemoveItem("non-existing.txt"); err != nil {
		fmt.Println("\n Removing non-existing file error:", err)
	} else {
		fmt.Println("\n ❌ Removed non-existing file")
	}
}