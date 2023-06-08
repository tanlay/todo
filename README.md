# todo
Golang+Gorm+MySQL实现一个todo,
- 增加todo项
- 通过id查询todo
- 查询todo列表
- 更新todo名称、分类
- 更新todo状态（未完成，已完成）

## 安装
```shell
go mod tidy
go build
```

## 使用
```shell
github.com/tanlay/todo -h      
命令行todo

Usage:                                                                  
  github.com/tanlay/todo [command]                                                        
                                                                        
Available Commands:                                                     
  completion  Generate the autocompletion script for the specified shell
  create      创建todo项                                                
  describe    通过ID查询todo项                                          
  help        Help about any command                                    
  query       查询todo列表                                              
  status      更新todo状态                                              
  update      更新todo                                                  
                                                                        
Flags:                                                                  
  -h, --help   help for github.com/tanlay/todo                                            
                                                                        
Use "todo [command] --help" for more information about a command.
```



## config.toml
```toml
[database]
dsn="mysql://root:123456@tcp(localhost:3306)/todo?parseTime=True"

[logger]
env="pord"
level="debug"
output="log.txt"

```