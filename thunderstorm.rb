class Thunderstorm < Formula
  desc "Open source CLI for buford to send push notifications to Apple devices through HTTP/2"
  homepage "https://github.com/macteo/thunderstorm"
  url "https://github.com/macteo/thunderstorm.git",
    :tag => "v0.1.3"

  head "https://github.com/macteo/thunderstorm.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    
    goCommand = "/usr/local/go/bin/go"
    
    system goCommand, "get", "golang.org/x/crypto/pkcs12"
    system goCommand, "get", "github.com/codegangsta/cli"
    system goCommand, "get", "github.com/RobotsAndPencils/buford"
    system goCommand, "get", "github.com/macteo/buford"
    system goCommand, "build", "-o", "thunderstorm"
    bin.install "thunderstorm"
  end
  
  test do
    system "#{bin}/thunderstorm", "-v"
  end
end