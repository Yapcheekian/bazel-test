## Binary
### Build
bazel build android:HelloWorld

### Run
./bazel-bin/android/HelloWorld

### Build and Run
bazel run android:HelloWorld

## Library
### Build
bazel build android:LibraryExample

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

## Clean
bazel clean

