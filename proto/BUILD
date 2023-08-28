load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "message_proto",
    srcs = ["message.proto"],
    visibility = ["//visibility:public"],
)

java_proto_library(
    name = "message_java_proto",
    deps = [":message_proto"],
)

load("@io_bazel_rules_go//proto:def.bzl", "go_proto_compile", "go_proto_library")

go_proto_library(
    name = "message_go_proto",
    importpath = "github.com/Yapcheekian/bazel-test/proto",
    proto = ":message_proto",
    visibility = ["//visibility:public"],
)

go_proto_compile(
    name = "person_go_proto",
    protos = ["@rules_proto_grpc//example/proto:person_proto"],
)

go_library(
    name = "proto",
    embed = [":message_go_proto"],
    importpath = "github.com/Yapcheekian/bazel-test/proto",
    visibility = ["//visibility:public"],
)
