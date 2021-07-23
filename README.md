# ranking

1.rpc.proto添加req,resp,service，method enum
2.cmd.go注册命令
3.gen_cmd_handle.sh,格式化，自定义修改
4.clientHandle.go注册分发器
5.gen_server_handle.sh，格式化，自定义修改


# todo
rankServer 参数校验
config 检查



协程退出 生成测试数据
proto3的默认字段返回 optiona处理 done
db操作封装到DB done
原子操作，比如dict和list的同步删除 
remrange 方法待优化，bug  done

RDB load/save 编码优化

# pprof:

_ "net/http/pprof"
go tool pprof http://localhost:8087/debug/pprof/heap
pprof -http 127.0.0.1:9090 http://127.0.0.1:8087/debug/pprof/heap

# load testing
