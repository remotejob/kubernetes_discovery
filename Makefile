all: push

# 0.0 shouldn't clobber any released builds
TAG =0.3
PREFIX = remotejob/kubernetes-discovery

binary: cvserver.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o cvserver

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

set: 
	 kubectl set image deployment/kubernetes-discovery kubernetes-discovery=$(PREFIX):$(TAG)

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
