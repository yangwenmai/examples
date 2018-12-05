# 分布式全链路跟踪系统实战

## 分布式全链路跟踪系统是什么？

分布式链路跟踪的核心基本都是 Google Dapper 论文所述，使用全局 TraceID 表示一条调用链，连接各个服务调用（用 SPAN 表示），通过分析 SPAN 之间的父子关系形成跟踪树。
另外通过中间件的埋点和业务自定义的 Annotation ，记录日志并采用收集器进行离线和在线分析，从而实现调用链跟踪、优化决策等信息。

## Dapper 是什么？

[Dapper](https://github.com/bigbully/Dapper-translation) 是 Google 发表的分布式链路跟踪的论文。

## Jaeger 是什么？

[Jaeger](https://github.com/jaegertracing/jaeger) 是 Uber 基于 Google Dapper 开发的分布式链路跟踪系统的实现。

## 其他

- Ziplin

----

# 实战演示

1. 本地运行 Docker `jaegertracing/all-in-one:latest`，打开 `http://localhost:16686` 如果能够看到 Jaeger UI 即表示运行成功了；
2. `git clone https://github.com/yangwenmai/jaeger-opentracing-examples.git`；
3. `GOOS=linux GOARCH=amd64 go build .` 将生成 `jaeger-opentracing-examples` 的二进制可执行文件；
4. `docker build -t jaeger-opentracing-examples:dev .` 制作一个本地 jaeger-opentracing-examples 的 Docker 镜像；
5. `docker run -it --name=jaeger-opentracing-examples -p 0.0.0.0:8080:8080 jaeger-opentracing-examples:dev`
6. 本地访问 `http://localhost:8080`，然后打开 `http://localhost:16686` Search 即可查看到你注册的全链路跟踪请求了。

----

# 坑/问题

1. Docker 运行要特别注意网络互通，默认 Docker 网络的 IP 是 `172.17.0.x`。
2. 实际运行中可能存在一定时间的延时，所以你访问请求之后需要等一会儿才可以在 Jaeger UI 上查看到请求。
3. 注意上报服务器地址是 `Reporter中的LocalAgentHostPort`；
4. 注意本文是使用 dep 包管理工具，特别要注意 vendor 中的 `github.com/apache/thrift/` 版本，如果你使用 `github.com/apache/thrift` master 分支，那么你将编译不通过；

# 参考资料

1. http://ginobefunny.com/post/learning_distributed_systems_tracing/
2. https://bigbully.github.io/Dapper-translation/
3. https://tech.meituan.com/mt-mtrace.html
4. https://github.com/jaegertracing/jaeger
5. https://github.com/alextanhongpin/go-jaeger-trace
6. http://github.com/jukylin/trace_example
7. http://github.com/jukylin/blog/blob/master/Uber分布式追踪系统Jaeger使用介绍和案例【PHP%20%20%20Hprose%20%20%20Go】.md
8. http://www.infoq.com/cn/articles/evolving-distributed-tracing-at-uber-engineering
