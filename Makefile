### Customized Makefile
.PHONY: gensdk
gensdk:
	docker run --rm -v "${PWD}/generated:/generated" -v "${PWD}/schema:/schema" openapitools/openapi-generator-cli generate -i ./schema/subscription-service.openapi3.json -g go -o ./generated/subscription --package-name=subscription
	docker run --rm -v "${PWD}/generated:/generated" -v "${PWD}/schema:/schema" openapitools/openapi-generator-cli generate -i ./schema/authentication-service.openapi3.json -g go -o ./generated/authentication/ --package-name=authentication
	docker run --rm -v "${PWD}/generated:/generated" -v "${PWD}/schema:/schema" openapitools/openapi-generator-cli generate -i ./schema/dashboard-service.openapi3.json -g go -o ./generated/dashboard/ --package-name=dashboard
	docker run --rm -v "${PWD}/generated:/generated" -v "${PWD}/schema:/schema" openapitools/openapi-generator-cli generate -i ./schema/account-service.openapi3.json -g go -o ./generated/account/ --package-name=account

.PHONY: version
version:
	$(ECHO) echo "Update changelog"
	standard-version
