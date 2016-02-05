class Thunderstorm < Formula
  desc "Open source CLI for buford to send push notifications to Apple devices through HTTP/2"
  homepage "https://github.com/macteo/thunderstorm"
  url "https://github.com/macteo/thunderstorm.git",
    :tag => "v0.1.0" #, :revision => "8a8336ae08b3ac5bb4ddce4a5ebbf16049961ad2"

  head "https://github.com/macteo/thunderstorm.git"

  # bottle do
  #   cellar :any_skip_relocation
  #   sha256 "09d708d36d4a1267de3d99f8318412ded60e7549d8915e079e2a2e522f29330f" => :el_capitan
  # end

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = cached_download/".gopath"
    ENV.append_path "PATH", "#{ENV["GOPATH"]}/bin"

    # # FIXTHIS: do this without mutating the cache!
    # hack_dir = cached_download/".gopath/src/github.com/macteo"
    # rm_rf hack_dir
    # mkdir_p hack_dir
    # ln_s cached_download, "#{hack_dir}/macteo"

    go get "golang.org/x/crypto/pkcs12"
    go get "github.com/codegangsta/cli"
    go get "github.com/RobotsAndPencils/buford"
    go install
  end
end