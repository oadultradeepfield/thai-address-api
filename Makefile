SERVICE_NAME := thai-address-api
REGION := asia-southeast1
IMAGE := gcr.io/$(GCP_PROJECT_ID)/$(SERVICE_NAME)

.PHONY: run lint docker-build docker-push deploy

run:
	go run cmd/server/main.go

lint:
	golangci-lint run --fix

docker-build:
	docker buildx build \
		--platform linux/amd64 \
		-t $(IMAGE) \
		--load .

docker-push: docker-build
	docker push $(IMAGE)

deploy: docker-push
	gcloud run deploy $(SERVICE_NAME) \
		--image $(IMAGE) \
		--platform managed \
		--region $(REGION) \
		--memory 128Mi \
		--cpu 0.5 \
		--max-instances 1 \
		--min-instances 0 \
		--timeout 30s \
		--allow-unauthenticated
