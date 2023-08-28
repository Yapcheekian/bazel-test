java_binary(
    name = "HelloWorld",
    srcs = [
        "src/HelloWorld.java",
    ],
    deps = [":LibraryExample"],
)

java_library(
    name = "LibraryExample",
    srcs = ["src/IntMultiplier.java"],
)

java_test(
    name = "LibraryExampleTest",
    srcs = ["test/IntMultiplierTest.java"],
    test_class = "IntMultiplierTest",
    deps = [
        ":LibraryExample",
        "//third_party:junit4",
    ],
)
