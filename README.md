A simple file copy CLI tool written in Go.  
This tool mimics the basic behavior of the `cp` command, with a clearer interface using `--from` and `--to` flags.

---

## Features

- Copy a file from source to destination
- Clear interface: `--from` and `--to`
- Prints success or error messages in English

---

## Usage

```bash
go run main.go --from source.txt --to destination.txt
```
Or after building:
```bash
go build -o mycp
./mycp --from source.txt --to destination.txt
```
```bash
$ ./mycp --from notes.txt --to backup/notes.txt
Copy completed: notes.txt → backup/notes.txt
```

## Requirements
- Go 1.18 or later
- Linux (recommended for testing)

License
MIT
> README drafted with help from ChatGPT
