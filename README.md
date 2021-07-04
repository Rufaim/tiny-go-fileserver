## Tiny Golang fileserver

A tiny sever to serve a static files.
Good for development or serving of static assets.

### How to use it
Download the repository and install it
```bash
go get github.com/Rufaim/tiny-go-fileserver 
```
Checkout to a source code location and install:
```bash
cd $GOPATH/src/github.com/Rufaim/tiny-go-fileserver
go install
```
Now you can pass a sequence of folders to serve.
```bash
tiny-go-fileserver "folder_name_1" "folder_name_2" "folder_name_3"
```
The easiest to check it on the current folder
```bash
tiny-go-fileserver .
```
Remember, if a provided path does not exist or not a folder, it will be ignored.
