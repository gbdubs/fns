# Fns - multicloud serverless utilities

## TL;DR

This is a command line tool to quickly create, deploy, modify and configure cloud functions written in go. By taking an opinionated stance on every element of function configuration, it allows you (the developer) to just think about your business logic.

## Status

This project is in its very early stages, I don't yet recommend its use.

## Install

```
go get -u github.com/gbdubs/fns
go build ~/go/src/github.com/gbdubs/fns
alias fns=./~/go/src/github.com/gbdubs/fns
fns help
```

## Concepts and Basic Usage

### Function

A function is a piece of code that runs on-demand on a cloud provider. The functions in this package can be run on Amazon Web Services (AWS) as a AWS Lambda, or Google Cloud Platform (GCP) as a GoogleCloudFunction.

### Function Types

This package supports four types of functions, 3 "serving" flavors, and 1 "internal" flavor:

- `UNAUTH`: Unauthenticated HTTP Functions, implements standard HTTP interface, used for serving websites over HTTP.
- `AUTH`: Authenticated HTTP Functions, provides authorization data in interface, used for interacting with users who are already logged in.
- `SYSTEM`: HTTP functions, called by trusted in-house systems (ex: a database trigger or asynchronous queue delivery).
- `RPC`: Internal remote procedure call functions, used to provide microservice interfaces between lambdas. Unlike the other function flavors, this type is "collapsible".

### Projects

Functions are packaged together into "projects" - functions in the same project are trusted to call one another, share resources, etc. Every function must belong to exactly one project. Each project can be deployed to a single cloud platform.

### Function Lifecycle

First, create a project:

```
fns create_project --project=my_project
```
This will create a new directory relative to your current directory, and creates an initial project manifest `project.json` It's recommended that you create your project from the source directory of your github so that the packaging plays nicely with go and github (ex: from pwd=`~/go/src/github.com/myusername` call `create_project` to create a new project with the package `github.com/myusername/my_project`)

```
fns create_fn --fn=my_fn --type=UNAUTH|AUTH|SYSTEM|RPC [--project=my_project]
```
Creates a new function wihin your project. You must specify the project name if you are not invoking the command from the directory of your project. This will create the following files:

- `service.proto`: [if type is RPC] specifies the service that your function will be serving.
- `function.go`: The function (where your business logic will run)
- `function_test.go`: Tests for your function
- `function.json`: Configuraiton for the function.

```
fns push_project [--project=my_project] [--platform=AWS|GCP]
```
Pushes all of the functions in your project to the specified cloud services provider. If this is the first time you're deploying, you'll have to specify the --platform argument.

Push should be thought of as similar to a git push - it changes the contents of the remote resource (the cloud service provider) to match the local state of the resource configurations. This means if you wanted to delete a given function on AWS, you'd just delete the `aws.json` file, and then push the changes to the cloud.

```
fns invoke --fn=my_fn --input=file_with_input.json [--output=my_output.json] [--project=my_project]
```
Sends a request to the running function with your input, and either prints the result, or saves it to an output file.

## Goodies

```
fns debug --fn=my_fn --input=file_with_input.json [--project=my_project]
```

Does golog-ing (see github.com/gbdubs/golog).

```
fns test [--fn=my_fn] [--project=my_project]
```

Runs all input output tests in the directory.

## Other Commands

```
fns push_fn [--fn=my_fn] [--project=my_project]
fns destroy_fn [--fn=my_fn] [--project=my_project]
fns destroy_project [--project=my_project]
```


