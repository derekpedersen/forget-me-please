export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

test:
	.tools/scripts/test.sh

set-version:
	./.helm/set-version.sh

build:
	.tools/scripts/build.sh

run: build
	./bin/forget-me-please

docker: build
	docker build ./ -t forget-me-please:latest --no-cache

publish:
	docker tag forget-me-please us.gcr.io/${GCLOUD_PROJECT_ID}/forget-me-please:${GIT_COMMIT_SHA}
	gcloud docker -- push us.gcr.io/${GCLOUD_PROJECT_ID}/forget-me-please:${GIT_COMMIT_SHA}

deploy: set-version
	helm upgrade forget-me-please .helm

secret:
	kubectl create -f .kubernetes/secret.yaml

kubernetes: build test docker publish deploy