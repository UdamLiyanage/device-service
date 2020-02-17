IMAGE := udamliyanage/device-service:v1.1:latest

test:
	true

image:
	docker build -t $(IMAGE) .

push-image:
	docker push $(IMAGE)


.PHONY: image push-image test