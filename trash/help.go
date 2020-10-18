package cli

const helpText = `

fns - a command line utility for managing serverless functions.
https://www.github.com/gbdubs/fns

Commands (* indicates not yet implemented)

 *	create_project		creates a new collection of functions in a new folder
 *	create_fn		creates a new function within a project
 *	deploy_project		deploys a project and all of its functions to the cloud
 *	deploy_fn		deploys a single function to the cloud
 *	invoke			invokes a given function with a set input 
 *	test			invokes all tests against the cloud, verifying expected results 
 *	destroy_fn		deletes a given function from the cloud + locally 
 *	destory_project		deletes a project and all of its functions from the cloud 

Use "fns help <command>" for detials about the command.

`

const createProjectHelpText = `

create_project - creates a new folder to house a collection of functions.

fns create_project --project=my_project_name [--dir=~/go/src/.../myproject]

A project is a collection of resources (functions, storage, ACLs, domains) that can all communicate with one another with trust by deafult. On your local machine, all of the configuration for these resources live in a top-level directory with the project's name. On a cloud service provider, all of the project resources will be prefixed by the project name, and will be located in the project's configured region. Projects can be deployed, tested, or destroyed as a unit. Projects can be version controlled, but they do not need to be.

Required Flags
	--project=my_project_name	The name of the directory to house the project.

Optional Flags
	--where=path/from/here		The path that the project should be created at. If not 
					specified, defaults to the present working directory.

Use "fns help" to learn about other commands.

`

const createFnHelpText = `

create_fn - creates a new function within a project.

fns create_fn --fn=my_fn_name [--project=my_project_name] [--fn_type=UNAUTH|AUTH|SYSTEM|RPC]

A function is a piece of code that can be run on any cloud service provider's serverless computing solution. This command creates a new folder within the project directory with the function's name, and creates the files to deploy, test and invoke the function. This package supports four types of functions, 3 "serving" flavors, and 1 "internal" flavor:

- UNAUTH: Unauthenticated HTTP Functions, implements standard HTTP interface, used for serving websites over HTTP.
- AUTH: Authenticated HTTP Functions, provides authorization data in interface, used for interacting with users who are already logged in.
- SYSTEM: HTTP functions, called by trusted in-house systems (ex: a database trigger or asynchronous queue delivery).
- RPC: Internal remote procedure call functions, used to provide microservice interfaces between lambdas. Unlike the other function flavors, this type is "collapsible".

Collapsible functions are a beautiful thing - they can either be compiled into all of thier calling functions, or they can be invoked as microservices. The benefit of this architechure is that it allows you to not be bound to a specific service configuration. For small projects, you can bundle and deploy all of your bizlogic into a single serving function/binary, without hemming yourself into this architechure long term. As a future goal, I want to support real-trafic analysis to support and inform how different functions are built and deployed - ideally trying to maximize some form of user expressed constraints, build size, and end-user latency. After initial development, you should expect RPC to be the primary type of function you create.

Required Flags
	--fn=my_new_fn_name		The name of the new directory to house the function.

Optional Flags
	--project=my_project_name	The name of the project to push this function from. Defaults 
					to the project the command is invoked within
	--fn_type=...			The type of the function to create (defaults to UNAUTH), see 
					above for function types and their properties

Use "fns help" to learn about other commands.

`

const deployProject = `

deploy_project - push a project to one or more cloud-service-providers, actualizing local configuration.

fns deploy_project [--project=my_project_name] [--provider=AWS|GCP] [--dryrun]

Pushes the local project configuration to all clouds that it is currently configured for. If provider is specified, this command only pushes to that provider. On the first push to a cloud provider, you will need to specify the --provider flag.

Optional Flags
	--project=my_project_name	The name of the project to push, defaults to PWD project
	--provider=AWS|GCP		The cloud service provider to push this project to.
	--dryrun			If set, no actual pushes will be performed, but a summary
					of what WOULD be performed will be printed to the console.

Use "fns help" to learn about other commands.

`

const deployFnHelpText = `

deploy_fn - push a function to one or more cloud-service-providers, actualizing local configuration.

fns deploy_fn [--project=my_project_name] [--fn=my_fn_name] [--provider=AWS|GCP] [--dryrun]

Pushes the local function configuration to all cloud service providers that it is presently configured for. If the provider is specified, the command only pushes to that provider. On the first push to a cloud provider, you will need to specify the --provider flag.

Optional Flags
	--project=my_project_name	The name of the project to push, defaults to the PWD project.
	--fu=my_function_name		The name of the function to push, defaults to the PWD fn.
	--provider=AWS|GCP		The cloud service provider to push this project to.
	--dryrun			If set, no actual pushes will be performed, but a summary
					of what WOULD have been performed will be printed to the
					console.

Use "fns help" to learn about other commands.

`
