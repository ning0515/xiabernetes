# 记录个人学习kubernetes的项目

对现在的我来说，kubernetes目前已经过于复杂，我希望跟着K8S早期的源码持续学习K8S。
但是我既没有GCE的环境，也不希望过早的引入etcd和docker client。
我认为用可视化的方式去模拟K8S的工作流程是不错的学习方式，一方面练习GO语言及标准库的使用，另一方面强化对于K8S的设计的理解。
通过理解k8s源码，自己尝试写一个mock-k8s：
    1.通过创建文件的方式模拟etcd的功能，可以直观感受k8s的功能
    2.通过在文件中写入数据的方式模拟创建POD和容器


目前还是在模仿K8S第一版开源的代码，持续学习中。。。


# 快速开始

1.创建pod
启动APIserver
go run .\apiserver.go -nodes "1.1.1.1"

执行xiaberctl命令

go run xiaberctl.go -a "http://127.0.0.1:8001" -f "../../template/createpod.json" create /pods

2.列出pod  

go run xiaberctl.go list /pods

#TODO 

1.打印不太标准
2.给apiserver发的path不对时，没有正确的日志，影响调试
3.没有写单元测试