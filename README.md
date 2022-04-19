# fumofuzzer
这只是一个练手的项目...

This is just a practice project...

# build
```shell
git clone https://github.com/AFKL-CUIT/fumofuzzer.git
cd fumofuzzer
go build .
```

# example
本项目大致模仿了[wfuzz](https://github.com/xmendez/wfuzz/)项目的使用方法。

```shell
./fumofuzzer -p "list,1-2-3-4" -o "./tests/assets/test.json" -t "http://127.0.0.1:9000/FUZ0Z"
```
上面的命令会将`FUZ0Z`分别替换为1，2，3，4。并分别请求。将`FUZZ`的结果以`json`格式放入`./tests/assets/test.json`。

```shell
./fumofuzzer -p "list,1-2-3-4" -p "list,a-b-c-d" -i "zip" -o "./tests/assets/test.json" -t "http://127.0.0.1:9000/FUZ0Z/FUZ1Z"
```
上面的命令会分别请求`/1/a`, `/2/b`...等，就像`python`的`zip`函数一样。

# report a bug
欢迎在 issues 上提交bug，也欢迎PR。

Bugs are welcome on issues, and PR are welcome.

# report a FUMO
Q: FUMO?

A: FUMO

# known problem
1. json格式输出的数据有待完善。
2. 缺少类似于`wfuzz`可以在命令行输出表格的格式。
3. 对于`FUZZ`其他地方（如：json、method...）的支持。