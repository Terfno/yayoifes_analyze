# yayoifes_analyze
来場者アナライズ

## Make
### docker build
Dockerfileをもとにimageを錬成するmake magic
```sh
$ make build
```

### docker run
make buildで錬成したコンテナでgoを実行する
```sh
$ make gotest
```

### docker images / ps -a
めんどいのでdlsで一覧表示させる
```sh
$ make dls
```
