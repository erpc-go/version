# version 程序版本信息 #
程序版本信息记录，配合makefile，自动在编译程序时填入程序版本信息。参考example

```
// 编译完成后的二进制程序，使用-v参数，可查看程序版本信息
./example -v    

out: 

{
  "name": "example",
  "gitTag": "v1.0",
  "gitCommit": "081a4355c05f1abea3c3dd63723e13a7a464b81d",
  "gitTreeState": "clean",
  "gitBranch": "master",
  "buildAuthor": "a",
  "buildDate": "2020-08-30T16:02:54+0800",
  "goVersion": "go1.13",
  "compiler": "gc",
  "platform": "darwin/amd64"
}
```
