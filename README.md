# anki-unpkg

A simple command-line tool to unpack Anki deck packages (`.apkg`) and collection packages (`.colpkg`).

It extracts the database file and all media files into a specified directory.

## Installation

With Go installed, you can install `anki-unpkg` with:

```sh
go install github.com/lftk/anki-unpkg@latest
```

## Usage

```sh
anki-unpkg <path> [<dir>]
```

-   `<path>`: The path to the Anki `.apkg` or `.colpkg` file you want to unpack.
-   `<dir>`: (Optional) The directory where the contents will be extracted. If not provided, a directory will be created with the same name as the package file (without the extension).

## Example

```sh
# Unpack my-deck.apkg into a new directory named "my-deck"
anki-unpkg my-deck.apkg

# Unpack collection.colpkg into a specific directory named "unpacked_collection"
anki-unpkg collection.colpkg unpacked_collection
```
