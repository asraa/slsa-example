#!/bin/bash -eu

# build project
# e.g.
# ./autogen.sh
# ./configure
# make -j$(nproc) all

# build fuzzers
# e.g.
# $CXX $CXXFLAGS -std=c++11 -Iinclude \
#     /path/to/name_of_fuzzer.cc -o $OUT/name_of_fuzzer \
#     $LIB_FUZZING_ENGINE /path/to/library.a
go get github.com/AdamKorcz/go-118-fuzz-build/utils

compile_native_go_fuzzer github.com/asraa/slsa-example/tests FuzzParseYaml fuzz_parseyaml gofuzz