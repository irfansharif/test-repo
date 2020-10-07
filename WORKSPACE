workspace(name = "testrepo")

# Load the thing that lets us load other things.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# Load go bazel tools.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2697f6bc7c529ee5e6a2d9799870b9ec9eaeb3ee7d70ed50b87a2c2c97e13d9e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.8/rules_go-v0.23.8.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.8/rules_go-v0.23.8.tar.gz",
    ],
)

# Load gazelle.

git_repository(
    name = "bazel_gazelle",
    branch = "fix_cxx_cpp",
    remote = "https://github.com/otan-cockroach/bazel-gazelle",
)

# Load go rules.
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

# Load gazelle dependencies.
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

new_local_repository(
    name = "libroach",
    path = "c-deps/libroach",
    build_file = "libroach.BUILD",
)

new_local_repository(
    name = "rocksdb",
    path = "c-deps/rocksdb",
    build_file = "rocksdb.BUILD",
)

local_repository(
    name = "gtest",
    path = "c-deps/googletest",
)
