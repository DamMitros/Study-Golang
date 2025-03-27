# Zadanie: Model Systemu Plików

Celem zadania jest stworzenie modelu systemu plików.

## Wymagania

W zadaniu należy zaimplementować cztery typy struktur, które spełniają interfejsy zdefiniowane w pliku `interfejs.go`:

### 1. **Plik**
- Implementuje: 
  - `FileSystemItem`
  - `Readable`
  - `Writable`

### 2. **Katalog**
- Implementuje: 
  - `Directory`

### 3. **SymLink**
- Implementuje: 
  - `FileSystemItem`
- Zawiera referencję do innej struktury implementującej `FileSystemItem`.

### 4. **PlikDoOdczytu**
- Implementuje: 
  - `FileSystemItem`
  - `Readable`
- Nie implementuje: 
  - `Writable`

## Struktura VirtualFileSystem

Należy stworzyć strukturę `VirtualFileSystem`, która umożliwia:
- Tworzenie plików i folderów.
- Otwieranie plików i folderów.
- Znajdowanie plików i folderów.
- Usuwanie plików i folderów.