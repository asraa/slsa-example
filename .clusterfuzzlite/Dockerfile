FROM gcr.io/oss-fuzz-base/base-builder-go:v1
RUN apt-get update && apt-get install -y make autoconf automake libtool
COPY . $SRC/slsa-example
WORKDIR slsa-example
COPY .clusterfuzzlite/build.sh $SRC/
