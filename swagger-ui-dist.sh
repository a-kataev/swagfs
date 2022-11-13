#!/usr/bin/env bash

version=${1:-4.15.5}

set -x

curl -Ls -o swagger-ui.zip "https://github.com/swagger-api/swagger-ui/archive/refs/tags/v${version}.zip"

unzip -o swagger-ui.zip "swagger-ui-${version}/dist/*.*" -d tmp

rm "tmp/swagger-ui-${version}/dist/swagger-initializer.js"

cp -rf "tmp/swagger-ui-${version}/dist" files/

rm -rf tmp swagger-ui.zip
