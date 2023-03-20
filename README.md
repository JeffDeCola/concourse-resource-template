# CONCOURSE RESOURCE TEMPLATE

[![Tag Latest](https://img.shields.io/github/v/tag/jeffdecola/concourse-resource-template)](https://github.com/JeffDeCola/concourse-resource-template/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/JeffDeCola/concourse-resource-template.svg)](https://pkg.go.dev/github.com/JeffDeCola/concourse-resource-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/concourse-resource-template)](https://goreportcard.com/report/github.com/JeffDeCola/concourse-resource-template)
[![codeclimate Maintainability](https://api.codeclimate.com/v1/badges/5abc7e41bcf4e122e7f7/maintainability)](https://codeclimate.com/github/JeffDeCola/concourse-resource-template/maintainability)
[![codeclimate Issue Count](https://codeclimate.com/github/JeffDeCola/concourse-resource-template/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/concourse-resource-template/issues)
[![Docker Pulls](https://badgen.net/docker/pulls/jeffdecola/concourse-resource-template?icon=docker&label=pulls)](https://hub.docker.com/r/jeffdecola/crypto-wallet-status/)
[![MIT License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)
[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)

_Used as a template for developing a Concourse resource type._

Table of Contents

* tbd

Documentation and Reference

* [concourse-resource-template Docker Image](https://hub.docker.com/r/jeffdecola/resource-template)
  on DockerHub
* This repos
  [github webpage](https://jeffdecola.github.io/concourse-resource-template/)
  _built with
  [concourse](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci-README.md)_

## USE EITHER BASH SCRIPT OR GO

This resource type can use either bash script or go.

Change _ci/Dockerfile_ to either ADD _/assets-go_ or _/assets-bash_.

### USING BASH

The 3 bash script files located in _/assets-bash_.

### USING GO (default)

The 3 bash scripts are located in _/assets-go_ that run _main.go_ with
the second argument being _check_, _in_ or _out_ respectively.
Hence only one file _main.go_ need to be maintained, rather
than three separate files.

## SOURCE CONFIGURATION

These are just placeholders that you can update where your source is.

* `source1`: Just a placeholder.

* `source2`: Just a placeholder.

## BEHAVIOR

### CHECK (a resource version(s))

CHECK will mimic getting the list of versions from a resource.

#### CHECK stdin

```json
{
  "source": {
    "source1": "sourcefoo1",
    "source2": "sourcefoo2"
  },
  "version": {
    "ref": "123 ",
  }
}
```

123 is the current version.

#### CHECK stdout

```json
[
  { "ref": "123" },
  { "ref": "3de" },
  { "ref": "456" }
  { "ref": "777" }
]
```

777 is the latest version that will be used.

The last number 777 will become the current ref version that will be used by IN.

#### CHECK - go run

```bash
echo '{
"params": {"param1": "Hello Clif","param2": "Nice to meet you"},
"source": {"source1": "sourcefoo1","source2": "sourcefoo2"},
"version":{"ref": "123"}}' |
go run main.go check $PWD
```

### IN (fetch a resource)

IN will mimic fetching a resource and placing a file in the working directory.

#### IN Parameters

* `param1`: Just a placeholder.

* `param2`: Just a placeholder.

#### IN stdin

```json
{
  "params": {
    "param1": "Hello Jeff",
    "param2": "Nice to meet you"
  },
  "source": {
    "source1": "sourcefoo1",
    "source2": "sourcefoo2"
  },
  "version": {
    "ref": "777",
  }
```

#### IN stdout

```json
{
  "version":{ "ref": "777" },
  "metadata": [
    { "name": "nameofmonkey", "value": "Larry" },
    { "name": "author","value": "Jeff DeCola" }
  ]
}
```

#### file fetched (fetch.json)

The IN will mimic a fetch and place a fake file `fetched.json` file
in the working directory:

#### IN - go run

```bash
echo '{
"params": {"param1": "Hello Clif","param2": "Nice to meet you"},
"source": {"source1": "sourcefoo1","source2": "sourcefoo2"},
"version":{"ref": "777"}}' |
go run main.go in $PWD
```

### OUT (update a resource)

OUT will mimic updating a resource.

#### OUT Parameters

* `param1`: Just a placeholder.

* `param2`: Just a placeholder

#### OUT stdin

```json
{
  "params": {
    "param1": "Hello Jeff",
    "param2": "How are you?"
  },
  "source": {
    "source1": "sourcefoo1",
    "source2": "sourcefoo2"
  },
  "version": {
    "ref": ""
  }
}
```

#### OUT stdout

```json
{
  "version":{ "ref": "777" },
  "metadata": [
    { "name": "nameofmonkey","value": "Henry" },
    { "name": "author","value": "Jeff DeCola" }
  ]
}
```

where 777 is the version you wanted to update.

#### OUT - go run

```bash
echo '{
"params": {"param1": "Hello Jeff","param2": "How are you?"},
"source": {"source1": "sourcefoo1","source2": "sourcefoo2"},
"version":{"ref": ""}}' |
go run main.go out $PWD
```

## PIPELINE EXAMPLE USING PUT

```yaml
jobs:
...
- name: your-job-name
  plan:
    ...
  - put: resource-template
    params: { param1: "hello jeff", param2: "How are you?" }

resource_types:
  ...
- name: jeffs-resource
  type: docker-image
  source:
   repository: jeffdecola/resource-template
   tag: latest

resources:
  ...
- name: resource-template
  type: jeffs-resource
  source:
    source1: foo1
    source1: foo2
```

GET would look similiar.




## TESTED, BUILT & PUSHED TO DOCKERHUB USING CONCOURSE

To automate the creation of the `resource-template` docker image, a concourse pipeline
will,

* Update README.md for resource-template github webpage.
* Unit Test the code.
* Build the docker image `resource-template` and push to DockerHub.

![IMAGE - resource-template concourse ci pipeline - IMAGE](docs/pics/resource-template-pipeline.jpg)

As seen in the pipeline diagram, the _resource-dump-to-dockerhub_ uses
the resource type
[docker-image](https://github.com/concourse/docker-image-resource)
to push a docker image to dockerhub.

`resource-template` also contains a few extra concourse resources:

* A resource (_resource-slack-alert_) uses a [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* A resource (_resource-repo-status_) use a [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.
* A resource ([_`resource-template`_](https://github.com/JeffDeCola/resource-template))
  that can be used as a starting point and template for creating other concourse
  ci resources.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).

TEST IT

This is located under ci-test