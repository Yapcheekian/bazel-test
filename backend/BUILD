load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "hello_world_go",
    embed = [":backend_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "backend_lib",
    srcs = ["hello_world.go"],
    importpath = "github.com/Yapcheekian/bazel-test/backend",
    visibility = ["//visibility:private"],
)
