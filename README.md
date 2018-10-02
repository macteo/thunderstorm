# Thunderstorm

You should be able to install Thunderstorm either compiling it directly or using [homebrew](http://brew.sh) on a Mac.

```bash
brew install https://raw.github.com/macteo/thunderstorm/master/thunderstorm.rb
```

If you want to compile it yourself, please [install Go](https://golang.org/doc/install#install) then on the terminal get the required dependancies.

```bash
go get -u golang.org/x/net/http2
go get -u golang.org/x/crypto/pkcs12
go get -u github.com/aai/gocrypto/pkcs7
go get -u github.com/RobotsAndPencils/buford
go get -u github.com/codegangsta/cli
```

Then build the executable

```bash
go build -ldflags="-w"
```
