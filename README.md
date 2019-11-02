# yayoifes_analyze
来場者アナライズ

## start
### database
read Makefile.
```sh
$ cd database
$ make init
```

### server
using ngrok.
```sh
$ ./ngrok http 8080
$ cd app
$ chmod +x ./run.sh
$ ./run.sh
```

## stop
### database
readmake file.
```sh
$ cd database
$ make d/down
```

### server
this shell script is not die. please enter `control + c` twice quickly. :fire:
