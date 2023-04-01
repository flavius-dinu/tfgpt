# Terraform GPT Helper

This Golang program is a CLI tool that integrates Terraform with OpenAI's GPT-3.5 Turbo to provide explanations for Terraform commands and concepts.

## Overview

The tool offers support for the following Terraform commands:

* plan
* validate
* destroy
* init
* show


Additionally, it includes functions to interact with OpenAI's GPT-3.5 Turbo API to request explanations for Terraform commands and concepts. As adoption will grow for GPT-4, this will change to use GPT-4 instead.




https://user-images.githubusercontent.com/57911064/229308688-0082f72a-ddb6-47cf-bab4-7f891c869f84.mp4



## Prerequisites
To use this tool, you need to have the following:

* Golang installed on your system

* Terraform installed on your system

* OpenAI API key


To generate the OpenAI API key you have to go to https://platform.openai.com/account/api-keys and create an api key, similar to:



https://user-images.githubusercontent.com/57911064/229308740-cdbf8574-992e-40e2-9ef6-f617c46fe657.mp4



### Setup

* Clone the repository

* Build the program using go build.

```bash
go build -o bin/tfgpt cmd/tfgpt/main.go
```


* Add tfgpt to your path

* Make sure you have your OpenAI API key stored as an environment variable or in a credentials file within the user's home directory (e.g., ~/.tfgpt/credentials).

* Run the compiled program with the desired Terraform command or concept. If you want to use it in conjunction with a Terraform command, you must be in the directory containing the terraform configuration.

This will not run destroy on your behalf, it will just show you a plan based on your configuration


## Usage

To use the tool, simply run the program with the desired Terraform command or concept:

```bash
tfgpt [command] 
```

Or

```bash
tfgpt concept terraform_concept
```


## Examples

```bash
tfgpt plan
tfgpt destroy
tfgpt init
tfgpt validate

tfgpt concept resource
tfgpt concept datasource
```
