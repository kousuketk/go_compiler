# go_compiler

### Docker build, run
```
$ docker build -t ubuntu:ccpp .
$ docker run -itv /Users/takahashi/Go2/go_compiler:/home/projects 01c8ea667b7f /bin/bash
```

### execute
```
# echo -n "10+2;" | go run main.go | ./asrun
-------- a.s ----------------

 .global main
main:
 mov $10, %rax
 mov $2, %rcx
 add %rcx, %rax
 ret
-------- result -------------
12

# echo -n "10-2;" | go run main.go | ./asrun
-------- a.s ----------------

 .global main
main:
 mov $10, %rax
 mov $2, %rcx
 sub %rcx, %rax
 ret
-------- result -------------
8
```