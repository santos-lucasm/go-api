# MANGADEX GO-API

## Description

{WIP}

## How to install

1. Install git: [Instructions](https://phoenixnap.com/kb/how-to-install-git-windows)
2. Install go: [Instructions](https://go.dev/doc/install)
3. Check your environment GOPATH variable, it if does not exist, create one pointing to wherever you wish to code your project
4. Inside your GOPATH folder, create `bin`, `pkg` and `src` folders
5. Clone this repository inside your clone folder

## How to execute the program

1. Go inside src/go-api folder you just cloned
2. Create a file named credentials.env, open it with a text editor and type the following
```
USER=<YOUR-MANGADEX-USER>
EMAIL=<YOUR-MANGADEX-EMAIL>
PASSWORD=<YOUR-MANGADEX-PASSWORD>
```
3. Save the file
4. In the same folder, type the following in the terminal
```
go install
go run go-api
```

## Observations

* Do not commit your credentials.env file to the repository
* Always push for revisions!
* Questions? Send me a message!
* * Iuqueta#9802
* * santos.lucasmmatheus@gmail.com 