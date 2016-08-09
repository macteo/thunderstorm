# Thunderstorm formula specifications
class Thunderstorm < Formula
  desc 'Open source CLI for buford to send push notifications
        to Apple devices through HTTP/2'
  homepage 'https://github.com/macteo/thunderstorm'
  # url "https://github.com/macteo/thunderstorm.git", :tag => "v0.2.0"

  url 'https://github.com/macteo/thunderstorm/archive/v0.2.1.tar.gz'
  sha256 '0bcc539498c918a8d139e02effb396ef7a13964385a87a6b1e0af9f689bfa1cf'

  head 'https://github.com/macteo/thunderstorm.git'

  depends_on 'go' => :build

  def install
    ENV['GOPATH'] = buildpath

    system 'go', 'get', 'golang.org/x/net/http2'
    system 'go', 'get', 'golang.org/x/crypto/pkcs12'
    system 'go', 'get', 'github.com/codegangsta/cli'
    system 'go', 'get', 'github.com/RobotsAndPencils/buford'
    system 'go', 'build', '-o', 'thunderstorm'
    bin.install 'thunderstorm'
  end

  test do
    system "#{bin}/thunderstorm", '-v'
  end
end
