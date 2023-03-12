```
❯ go run main.go snack_to_camel
snackToCamel
```

```
❯ go run main.go to-camel-case snack_to_camel
snackToCamel
```

```
❯ cat test.xx
asdf_sdf
sdf_sdf_sdf
xcv_sdf_sdfsdf

❯ go run main.go to-camel-case file test.xx
asdfSdf
sdfSdfSdf
xcvSdfSdfsdf
```