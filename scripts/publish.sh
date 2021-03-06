#!/usr/bin/env bash

VERSION=$(git describe --tags)
echo "Publishing $VERSION..."

mkdir dist
mkdir releases
gox -osarch="linux/amd64" -osarch="linux/386" -osarch="darwin/amd64" -osarch="freebsd/amd64" -osarch="freebsd/386" -ldflags "-X main.Version=$VERSION" -output "dist/{{.OS}}_{{.Arch}}/osiw-server" ./cmd/osiw-server
gox -osarch="linux/amd64" -osarch="linux/386" -osarch="darwin/amd64" -osarch="freebsd/amd64" -osarch="freebsd/386" -ldflags "-X main.Version=$VERSION" -output "dist/{{.OS}}_{{.Arch}}/osiw-client" ./cmd/osiw-client

for i in dist/* ; do
  if [ -d "$i" ]; then
   ARCH=$(basename "$i")

   # Server
   mkdir "osiw-server_$VERSION"
   cp "dist/$ARCH/osiw-server" "osiw-server_$VERSION"
   zip -r "releases/osiw-server_$VERSION-$ARCH.zip" "osiw-server_$VERSION"
   cd releases
   sha256sum "osiw-server_$VERSION-$ARCH.zip" > "osiw-server_$VERSION-$ARCH.zip.sha256sum"
   cd ..
   rm -rf "osiw-server_$VERSION"

   # Client
   mkdir "osiw-client_$VERSION"
   cp "dist/$ARCH/osiw-client" "osiw-client_$VERSION"
   zip -r "releases/osiw-client_$VERSION-$ARCH.zip" "osiw-client_$VERSION"
   cd releases
   sha256sum "osiw-client_$VERSION-$ARCH.zip" > "osiw-client_$VERSION-$ARCH.zip.sha256sum"
   cd ..
   rm -rf "osiw-client_$VERSION"
  fi
done

#ghr -t "$GITHUB_TOKEN" -u jirwin -r osiw --replace "$VERSION" releases/

#rm -rf dist releases
