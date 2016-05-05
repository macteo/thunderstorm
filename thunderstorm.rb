class Thunderstorm < Formula
  desc "Open source CLI for buford to send push notifications to Apple devices through HTTP/2"
  homepage "https://github.com/macteo/thunderstorm"
  url "https://github.com/macteo/thunderstorm.git",
    :tag => "v0.2.0"

  head "https://github.com/macteo/thunderstorm.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath

    # goCommand = "go"

    system "go", "get", "golang.org/x/net/http2"
    system "go", "get", "golang.org/x/crypto/pkcs12"
    system "go", "get", "github.com/codegangsta/cli"
    system "go", "get", "github.com/RobotsAndPencils/buford"
    system "go", "build", "-o", "thunderstorm"
    bin.install "thunderstorm"
  end

  test do
    system "#{bin}/thunderstorm", "-v"
  end
end
