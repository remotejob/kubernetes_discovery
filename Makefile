ll: push

# 0.0 shouldn't clobber any released builds
TAG =0.2
PREFIX = gcr.io/jntlserv0/cvserver

binary: cvserver.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o cvserver

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	gcloud docker push $(PREFIX):$(TAG)

set: push
	 kubectl set image deployment/cv-app cv-app=$(PREFIX):$(TAG)

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
