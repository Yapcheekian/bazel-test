load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/Yapcheekian/bazel-test
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

gazelle(
    name = "gazelle-update",
    command = "update",
)
