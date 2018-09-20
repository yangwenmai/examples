build:
	# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
	GOOS=linux GOARCH=amd64 go build	
	# 根据当前目录下的 Dockerfile 生成名为 local/api 的镜像
	docker build -t yangwen13/api .
	rm api

push:
	docker push yangwen13/api
	
deploy: build
	kubectl apply -f k8s-api.yaml
