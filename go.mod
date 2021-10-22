module github.com/Azure/helmrelease-workflow-executor

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	k8s.io/client-go v0.21.3
	k8s.io/kubectl v0.21.0
	sigs.k8s.io/controller-runtime v0.9.5
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v1.4.2-0.20200203170920-46ec8731fbce
)
