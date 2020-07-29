include Makefile.f5cs

### Customized Makefile
.PHONY: gensdk
gensdk:
	docker run --rm -v "${PWD}/generated:/generated" -v "${PWD}/schema:/schema" openapitools/openapi-generator-cli generate -i ./schema/subscription-service.openapi3.json -g go -o ./generated/subscription --package-name=subscription
	docker run --rm -v "${PWD}/generated:/generated" -v "${PWD}/schema:/schema" openapitools/openapi-generator-cli generate -i ./schema/authentication-service.openapi3.json -g go -o ./generated/authentication/ --package-name=authentication