# Generate client language models

npx @openapitools/openapi-generator-cli generate --global-property models -i  type-definition.yaml --skip-validate-spec  -g python -o python/aisStream --package-name aisStream
npx @openapitools/openapi-generator-cli generate --global-property models -i  type-definition.yaml --skip-validate-spec  -g go -o golang/aisStream --package-name aisStream
npx @openapitools/openapi-generator-cli generate --global-property models -i  type-definition.yaml --skip-validate-spec  -g java -o java/aisStream --package-name aisStream
npx @openapitools/openapi-generator-cli generate --global-property models -i  type-definition.yaml --skip-validate-spec  -g typescript -o typescript/aisStream --package-name aisStream
npx @openapitools/openapi-generator-cli generate --global-property models -i  type-definition.yaml --skip-validate-spec  -g markdown -o markdown/aisStream --package-name aisStream


