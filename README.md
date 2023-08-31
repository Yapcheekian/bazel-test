## Binary
### Build
bazel build android:HelloWorld

### Run
./bazel-bin/android/HelloWorld

### Build and Run
bazel run android:HelloWorld
bazel run backend:hello_world_go
bazel run client-server-go:echo_server
bazel run client-server-java:EchoClient

## Library
### Build
bazel build android:LibraryExample
bazel build proto:message_proto

### Test
bazel test android:LibraryExampleTest

### third_party
bazel build third_party:junit4
bazel build third_party:all

## All
bazel build android:all
bazel build android/...
bazel build third_party:all
bazel build ...
bazel build //...

## Clean
bazel clean

## Gazelle
gazelle -go_prefix github.com/Yapcheekian/bazel-test
gazelle update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune
bazel run //:gazelle
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
bazel run //:gazelle -- update
