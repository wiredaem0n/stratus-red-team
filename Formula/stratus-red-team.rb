# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class StratusRedTeam < Formula
  desc ""
  homepage "https://stratus-red-team.cloud"
  version "2.5.0"
  license "Apache-2.0"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.0/stratus-red-team_2.5.0_Darwin_x86_64.tar.gz"
      sha256 "5b73b7e98434e7269c628ad70e36181b01bd0ffe4a255fc7531b3bea2a611c37"

      def install
        bin.install "stratus"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.0/stratus-red-team_2.5.0_Darwin_arm64.tar.gz"
      sha256 "182c9b4b63bafed4949ef1dd8ed4eeca2ac506ebd82a70ff86207d8c571fb875"

      def install
        bin.install "stratus"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.0/stratus-red-team_2.5.0_Linux_arm64.tar.gz"
      sha256 "edcb1388a53d84bdaaaa95646f6ef1c6a8edfd137af5924088ad511c40a37f22"

      def install
        bin.install "stratus"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.0/stratus-red-team_2.5.0_Linux_x86_64.tar.gz"
      sha256 "e4dcf344a6c5e2af77de1f803a1d3d99ea72aaf6bfe2824d1b617aa804054676"

      def install
        bin.install "stratus"
      end
    end
  end
end
