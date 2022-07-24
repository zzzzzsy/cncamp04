# include .env
# Image URL to use all building/pushing image targets
IMG ?= zzzzzsy/cncamp04:latest
TARGET_PORT ?= 8080
LOCAL_PORT ?= 8443

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

test: fmt vet ## Run tests.
	go test ./... -coverprofile cover.out

docker-build: test ## Build docker image
	docker build -t ${IMG} .

docker-run: docker-build ## Run service on local port 8443
	docker run --name httpclient -d -p ${LOCAL_PORT}:${TARGET_PORT} ${IMG}

docker-login: ## Run docker login before push the image to the dockerhub
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD)

docker-push: docker-login ## Push docker image
	docker push ${IMG}

pre-apply: ## create namespace first
    kubectl apply -f module08/manifests/namespace.yaml

apply: pre-apply ## apply to target k8s cluster
	kubectl apply -f module08/manifests

cleanup:
    kubectl delete ns cncamp

debug:
	echo ${TARGET_PORT}
	echo ${IMG}
	echo ${LOCAL_PORT}
	echo ${GOBIN}
