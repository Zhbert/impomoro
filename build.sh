#!/bin/bash

ver='0.1'

echo "Cleaning packages..."
rm -rf ./releases
mkdir ./releases

echo "Preparing the releases: $ver"
echo "Building Linux versions"
fyne-cross linux -arch=*
echo "Copying artifacts..."
cp fyne-cross/dist/linux-386/impomoro.tar.xz ./releases/impomoro-linux-i386-$ver.tar.xz
cp fyne-cross/dist/linux-amd64/impomoro.tar.xz ./releases/impomoro-linux-amd64-$ver.tar.xz
cp fyne-cross/dist/linux-arm/impomoro.tar.xz ./releases/impomoro-linux-arm-$ver.tar.xz
cp fyne-cross/dist/linux-arm64/impomoro.tar.xz ./releases/impomoro-linux-arm64-$ver.tar.xz
echo "Done."

echo "Building Windows versions"
fyne-cross windows -arch=*
echo "Copying artifacts..."
cp fyne-cross/dist/windows-386/impomoro.exe.zip ./releases/impomoro-windows-i386-$ver.zip
cp fyne-cross/dist/windows-amd64/impomoro.exe.zip ./releases/impomoro-windows-amd64-$ver.zip
cp fyne-cross/dist/windows-arm64/impomoro.exe.zip ./releases/impomoro-windows-arm64-$ver.zip
echo "Done."

echo "Building FreeBSD versions"
fyne-cross freebsd -arch=*
echo "Copying artifacts..."
cp fyne-cross/dist/freebsd-amd64/impomoro.tar.xz ./releases/impomoro-freebsd-amd64-$ver.zip
cp fyne-cross/dist/freebsd-arm64/impomoro.tar.xz ./releases/impomoro-freebsd-arm64-$ver.zip
echo "Done."

echo "Building macOS versions"
fyne-cross darwin --macosx-sdk-path ~/SDKs/MacOSX11.1.sdk/ -arch=* -app-id impomoro
echo "Copying artifacts..."artifacts
# shellcheck disable=SC2164
cd fyne-cross/dist/darwin-amd64/
zip -r impomoro-macos-amd64-$ver.zip impomoro.app
# shellcheck disable=SC2164
cd ../../../fyne-cross/dist/darwin-arm64
zip -r impomoro-macos-arm64-$ver.zip impomoro.app
cd ../../../
cp fyne-cross/dist/darwin-arm64/impomoro-macos-arm64-$ver.zip ./releases/impomoro-macos-arm64-$ver.zip
cp fyne-cross/dist/darwin-amd64/impomoro-macos-amd64-$ver.zip ./releases/impomoro-macos-amd64-$ver.zip
echo "Done."
echo "The assembly is complete."
