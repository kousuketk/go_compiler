# go_compiler

### Docker build, run
```
$ docker build -t ubuntu:ccpp .
$ docker run -itv /Users/takahashi/Go2/go_compiler:/home/projects 01c8ea667b7f /bin/bash
```

### execute
```
# # echo -n 40 | go run main.go | ./asrun 
-------- a.s ----------------

 .global main
main:
 mov $40, %rax
 ret
-------- result -------------
40

```