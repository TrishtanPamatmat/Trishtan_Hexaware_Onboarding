go_binary(
  name = "main",
  srcs = ["main.go"],
  deps = [
        ":fizzbizz",
    ],
)

go_library(
  name = "fizzbizz",
  srcs = ["fizzbizz.go"],
)

go_test(
    name = "fizzbizz_test",
    srcs = ["fizzbizz_test.go"],
 deps = [
        ":fizzbizz",
        "//third_party/go:assert",
    ],

)