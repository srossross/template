class Template < Formula
  desc "Command line template util"
  homepage "https://srossross.github.io/template"
  url "https://github.com/srossross/template/releases/download/v0.0.8/template-darwin-amd64.tgz"
  version "v0.0.8"
  sha256 "db64e7a6d8f85f3be68fa1865787cc19d07cf1a2ee19bf4d518d1c32b74c246d"

  def install
    system "mv", "template-darwin-amd64", "template"
    bin.install "template"
  end

  test do
    system "#{bin}/template", "version"
  end
end
