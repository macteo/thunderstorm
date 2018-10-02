# Thunderstorm formula specifications
class Thunderstorm < Formula
  desc 'Open source CLI for buford to send push notifications
        to Apple devices through HTTP/2'
  homepage 'https://github.com/macteo/thunderstorm'
  # url "https://github.com/macteo/thunderstorm.git", :tag => "v0.2.0"

  url 'https://github.com/macteo/thunderstorm/archive/v0.2.1.tar.gz'
  sha256 '07294758c2a55c4c0dc45990e4951ed58f2bbb90867e8fb2542364abb7fb7d51'

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
