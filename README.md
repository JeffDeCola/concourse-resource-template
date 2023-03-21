# CONCOURSE RESOURCE TEMPLATE

[![Tag Latest](https://img.shields.io/github/v/tag/jeffdecola/concourse-resource-template)](https://github.com/JeffDeCola/concourse-resource-template/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/JeffDeCola/concourse-resource-template.svg)](https://pkg.go.dev/github.com/JeffDeCola/concourse-resource-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/concourse-resource-template)](https://goreportcard.com/report/github.com/JeffDeCola/concourse-resource-template)
[![codeclimate Maintainability](https://api.codeclimate.com/v1/badges/5abc7e41bcf4e122e7f7/maintainability)](https://codeclimate.com/github/JeffDeCola/concourse-resource-template/maintainability)
[![codeclimate Issue Count](https://codeclimate.com/github/JeffDeCola/concourse-resource-template/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/concourse-resource-template/issues)
[![Docker Pulls](https://badgen.net/docker/pulls/jeffdecola/concourse-resource-template?icon=docker&label=pulls)](https://hub.docker.com/r/jeffdecola/concourse-resource-template/)
[![MIT License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)
[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)

_Can be used as a template for developing a concourse resource._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/concourse-resource-template#overview)
  * [CHECK](https://github.com/JeffDeCola/concourse-resource-template#check)
  * [IN](https://github.com/JeffDeCola/concourse-resource-template#in)
  * [OUT](https://github.com/JeffDeCola/concourse-resource-template#out)
* [BUILD AND PUSH THE RESOURCE](https://github.com/JeffDeCola/concourse-resource-template#build-and-push-the-resource)
* [TEST THIS RESOURCE](https://github.com/JeffDeCola/concourse-resource-template#test-this-resource)

Documentation and Reference

* [concourse-resource-template docker image](https://hub.docker.com/r/jeffdecola/concourse-resource-template)
  on DockerHub
* This repos
  [github webpage](https://jeffdecola.github.io/concourse-resource-template/)
  _built with
  [concourse](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci-README.md)_

## OVERVIEW

A concourse resource is a docker image.

It requires 3 kinds of scripts or executables,

* **check** - Detecting new versions of the resource (i.e. git version)
* **in** - Fetching something
* **out** - Updating something

You build your Resource Docker Image with a Dockerfile by using the
`concourse docker base image` and adding your executables to `/opt/resource`.

The three scripts/executables can be written with bash or go,

* The go is build is located in
  [/build-resource-using-go](https://github.com/JeffDeCola/concourse-resource-template/tree/master/build-resource-using-go)
  (In development)
* The bash build is located in
  [/build-resource-using-bash](https://github.com/JeffDeCola/concourse-resource-template/tree/master/build-resource-using-bash)

### CHECK

[check](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/check-in-out/check)
will get a list of versions from a resource.

CHECK stdin,

```json
{
  "source": {
    "user": "username",
    "password": "mypassword"
  },
  "version": {
    "ref": "123",
  }
}
```

123 is the current version.

CHECK stdout,

```json
[
  { "ref": "123" },
  { "ref": "3de" },
  { "ref": "456" },
  { "ref": "777" }
]
```

777 is the latest version that will be used.

The last number 777 will become the current ref version that will be used by IN.

### IN

[in](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/check-in-out/in)
will mimic **fetching a resource** and placing a file in the working directory.

IN Parameters,

* `param1`: Just a placeholder.
* `param2`: Just a placeholder.

IN stdin,

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

IN stdout,

```json
{
  "version":{ "ref": "777" },
  "metadata": [
    { "name": "nameofmonkey", "value": "Larry" },
    { "name": "author","value": "Jeff DeCola" }
  ]
}
```

The IN will mimic a fetch and place a fake file `fetched.json` file
in the working directory:

### OUT

[out](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/check-in-out/out)
will mimic **updating a resource**.

OUT Parameters,

* `param1`: Just a placeholder.
* `param2`: Just a placeholder

OUT stdin,

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

OUT stdout,

```json
{
  "version":{ "ref": "777" },
  "metadata": [
    { "name": "nameofmonkey","value": "Henry" },
    { "name": "author","value": "Jeff DeCola" }
  ]
}
```

Where 777 is the version you wanted to update.

## BUILD AND PUSH THE RESOURCE

I am using bash to build the resource. go is still in development.

To
[build.sh](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/build/build.sh)
the resource docker image using the
[Dockerfile](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/build/Dockerfile),

```bash

```bash
cd build-resource-using-bash/build
sh build-resource.sh
```

Note how a concourse base image is used to build the resource.

To
[push.sh](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/push/push.sh)
the resource docker image to dockerhub,

```bash
cd build-resource-using-bash/push
sh push.sh
```

## TEST THIS RESOURCE

To [set-pipeline.sh](https://github.com/JeffDeCola/concourse-resource-template/blob/master/test-this-resource/set-pipeline.sh)
for concourse,

```bash
cd test-this-resource
sh set-pipeline.sh.
```

Where the
[pipeline.yml](https://github.com/JeffDeCola/concourse-resource-template/blob/master/test-resource/pipeline.yml),

```yaml
jobs:
...
- name: your-job-name
  plan:
    ...
    - put: concourse-resource-template
      params: 
        param1: "hello jeff"
        param2: "How are you?"

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
      source2: foo2
```
