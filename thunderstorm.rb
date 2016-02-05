class Thunderstorm < Formula
  desc "Open source CLI for buford to send push notifications to Apple devices through HTTP/2"
  homepage "https://github.com/macteo/thunderstorm"
  url "https://github.com/macteo/thunderstorm.git",
    :tag => "v0.1.1" #, :revision => "8a8336ae08b3ac5bb4ddce4a5ebbf16049961ad2"

  head "https://github.com/macteo/thunderstorm.git"

  # bottle do
  #   cellar :any_skip_relocation
  #   sha256 "09d708d36d4a1267de3d99f8318412ded60e7549d8915e079e2a2e522f29330f" => :el_capitan
  # end

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