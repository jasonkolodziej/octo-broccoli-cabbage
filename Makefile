# TODO: Replace with your new repo name
NEW_REPO_PATH='yourscmprovider.com/youruser/yourrepo'




BUF_VERSION:=$(shell curl -sSL https://api.github.com/repos/bufbuild/buf/releases/latest \
                   | grep '"name":' \
                   | head -1 \
                   | cut -d : -f 2,3 \
                   | tr -d '[:space:]\",')

CURRENT_REPO_PATH=$(shell go mod why | tail -n1)

GIT_REPO_PATH:=$(shell git config --get remote.origin.url)


# extract the protocol
proto=$(shell echo ${GIT_REPO_PATH} | grep :// | sed -e's,^\(.*://\).*,\1,g')
# remove the protocol -- updated
url=$(shell echo ${GIT_REPO_PATH} | sed -e s,${proto},,g)

no_tail=$(shell echo ${url} | sed 's/.git.*//')
# extract the user (if any)
user=$(shell echo ${url} | grep @ | cut -d@ -f1)
# # extract the host and port -- updated
# hostport=$(echo $url | sed -e s,$user@,,g | cut -d/ -f1)
# # by request host without port
# host="$(echo $hostport | sed -e 's,:.*,,g')"
# # by request - try to extract the port
# port="$(echo $hostport | sed -e 's,^.*:,:,g' -e 's,.*:\([0-9]*\).*,\1,g' -e 's,[^0-9],,g')"
# # extract the path (if any)
# path="$(echo $url | grep / | cut -d/ -f2-)"
remote_url:
	@echo ${no_tail}
generate:
	buf --debug --verbose generate

lint:
	buf lint
	buf breaking --against 'https://github.com/johanbrandhorst/grpc-gateway-boilerplate.git#branch=master'

# Installs buf.build
# "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)"
install:
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"

update:
#! Deprecated
#	buf --debug --verbose mod update
	buf --debug --verbose dep update

init:
	buf --debug --verbose mod init


clean: # lint
	GO111MODULE=on go mod tidy
	rm -rf bin/*

build: clean
	GO111MODULE=on go build -o bin/protoc-gen-go-redact ./cmd/protoc-gen-go-redact
	cp ./cmd/protoc-gen-go-redact/internal/template.tmpl bin/template.tmpl

buf_clean:
	buf --debug --verbose mod clear-cache
# 1> Install buf with make install, which is necessary for us to generate the Go and OpenAPIv2 files.
# 2> If you forked this repo, or cloned it into a different directory from the github structure,
#	 you will need to correct the import paths.
#	 Here's a nice find one-liner for accomplishing this
#    (replace yourscmprovider.com/youruser/yourrepo with your cloned repo path):
# find . -path ./vendor -prune -o -type f \( -name '*.go' -o -name '*.proto' \) -exec sed -i -e "s;github.com/johanbrandhorst/grpc-gateway-boilerplate;yourscmprovider.com/youruser/yourrepo;g" {} +
# find . -path ./vendor -prune -o -type f \( -name '*.go' -o -name '*.proto' \) -exec sed -i -e "s;${CURRENT_REPO_PATH};${NEW_REPO_PATH};g" {} +

adjust_template:
ifeq ($(NEW_REPO_PATH),'yourscmprovider.com/youruser/yourrepo')
	@read -p "What is your new/cloned/forked repository's path? (e.g. ${NEW_REPO_PATH}): " new_repo; \
	NEW_REPO_PATH=$$new_repo; \
	find . -path ./vendor -prune -o -type f \( -name '*.go' -o -name '*.proto' -o -name 'go.mod' \) -exec sed -i -e "s;${CURRENT_REPO_PATH};$$NEW_REPO_PATH;g" {} +
else
	find . -path ./vendor -prune -o -type f \( -name '*.go' -o -name '*.proto' -o -name 'go.mod' \) -exec sed -i -e "s;${CURRENT_REPO_PATH};${NEW_REPO_PATH};g" {} +
endif

# purge_old removes the excess files and should be used after adjust_template
purge_old:
	find . -path ./vendor -o -type f \( -name '*.go-e' -o -name '*.proto-e' -o -name 'go.mod-e' \) | xargs rm

latest_version:
	@echo ${BUF_VERSION}
