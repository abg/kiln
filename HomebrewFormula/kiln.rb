# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kiln < Formula
  desc ""
  homepage ""
  version "0.69.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.69.0/kiln-darwin-arm64-0.69.0.tar.gz"
      sha256 "d4d77b9db89df12f42a829db10be4bb72a2664a958282ecbbef5a1568b367c72"

      def install
        bin.install "kiln"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.69.0/kiln-darwin-amd64-0.69.0.tar.gz"
      sha256 "094f61fb20620d149b99a542ef5159d0c296db438c1778e181c110da4338e893"

      def install
        bin.install "kiln"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.69.0/kiln-linux-amd64-0.69.0.tar.gz"
      sha256 "c49a5c580705bed11c7de0d0cca0f7745c96910515b8a947cfe7d561d1e57bfd"

      def install
        bin.install "kiln"
      end
    end
  end

  test do
    system "#{bin}/kiln --version"
  end
end
