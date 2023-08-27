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
  deps = [
      ":LibraryExample",
      "//third_party:junit4",
  ],
  test_class = "IntMultiplierTest",
)
