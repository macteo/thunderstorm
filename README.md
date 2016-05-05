# Thunderstorm

You should be able to install Thunderstorm either compiling it directly or, only if you already have Go 1.6 installed and have a mac, using [homebrew](http://brew.sh).

```bash
brew install https://raw.github.com/macteo/thunderstorm/master/thunderstorm.rb
```

If you want to compile it yourself, please install Go 1.6 and then

```bash
go get -u golang.org/x/crypto/pkcs12
go get -u github.com/aai/gocrypto/pkcs7
go get -u github.com/RobotsAndPencils/buford
```

More instructions to follow.

```bash
go build -gcflags="-newexport" -ldflags="-w"
```