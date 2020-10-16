# Top level namespace, everything is addressable through @cockroach//...
workspace(name = "test_repo")

# Load the things that lets us load other things.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# Load gazelle.
#
# TODO(irfansharif): Point to a proper release instead bazelbuild picks up up
# https://github.com/bazelbuild/bazel-gazelle/pull/933
git_repository(
    name = "bazel_gazelle",
    remote = "https://github.com/bazelbuild/bazel-gazelle",
    commit = "987c29003d1616c0b2499fa7f5ce0f3bc11bc9a3",
)

# Load go bazel tools.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2697f6bc7c529ee5e6a2d9799870b9ec9eaeb3ee7d70ed50b87a2c2c97e13d9e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.8/rules_go-v0.23.8.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.8/rules_go-v0.23.8.tar.gz",
    ],
)

# Load go rules, and invoke them.
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

# Load gazelle dependencies.
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

# Load protobuf depedency.
#
# Ref: https://github.com/bazelbuild/rules_go/blob/0.19.0/go/workspace.rst#proto-dependencies
#      https://github.com/bazelbuild/bazel-gazelle/issues/591
#
# XXX: We're not using this, are we?
git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
)
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()

# Load bazel utility that lets us build C/C++ projects using cmake.
#
# TODO(irfansharif): Get this merged upstream (it adds autoconf support), and
# point to it instead.
git_repository(
   name = "rules_foreign_cc",
   commit = "f6a15abd55be915b914aa618b50831bf5981340f",
   remote = "https://github.com/otan-cockroach/rules_foreign_cc",
)
load("@rules_foreign_cc//:workspace_definitions.bzl", "rules_foreign_cc_dependencies")
rules_foreign_cc_dependencies()

BUILD_ALL_CONTENT = """filegroup(name = "all", srcs = glob(["**"]), visibility = ["//visibility:public"])"""

new_local_repository(
  name = "jemalloc",
  path = "c-deps/jemalloc",
  build_file_content = BUILD_ALL_CONTENT,
)

new_local_repository(
   name = "cryptopp",
   path = "c-deps/cryptopp",
   build_file_content = BUILD_ALL_CONTENT,
)

new_local_repository(
    name = "rocksdb",
    path = "c-deps/rocksdb",
    build_file_content = BUILD_ALL_CONTENT,
)

new_local_repository(
    name = "snappy",
    path = "c-deps/snappy",
    build_file_content = BUILD_ALL_CONTENT,
)

new_local_repository(
    name = "libroach",
    path = "c-deps/libroach",
    build_file_content = BUILD_ALL_CONTENT,
)

new_local_repository(
    name = "googletest",
    path = "c-deps/googletest",
    build_file_content = BUILD_ALL_CONTENT,
)

BUILD_PROTOBUF_CONTENT = """filegroup(name = "all", srcs = glob(["**"], exclude=["src/google/protobuf/compiler/js/well_known_types_embed.cc"]), visibility = ["//visibility:public"])"""
new_local_repository(
   name = "protobuf",
   path = "c-deps/protobuf",
   build_file_content = BUILD_PROTOBUF_CONTENT,
)
