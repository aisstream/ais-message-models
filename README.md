# ais-message-models
This repository contains the OpenAPI definitions for all data models used by aisstream.io and generated libraries.

Getting Started
To use these data models in your own project, you can either:

- Use the generated libraries available in multiple languages (Java, Python, etc.). Many of these models are available as libraries in the respective languages package manager. 
- Use the OpenAPI definitions to generate your own libraries

## Prerequisites
To generate your own libraries, you will need to install the following tools:
- [openapi-generator](https://github.com/OpenAPITools/openapi-generator)

To generate your own libraries, you can use the OpenAPI Generator as follows:

Copy code
```
openapi-generator generate -i openapi.yaml -g <language> -o <output directory>
Replace <language> with the programming language of your choice (e.g. java, python, etc.), and <output directory> with the directory where you want the generated code to be stored.
```

## Contributing
We welcome contributions to the ais-message-models repository. If you have any suggestions or bug reports, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
