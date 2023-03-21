  _built with
  [concourse](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci-README.md)_

# OVERVIEW

A concourse resource is a docker image.

It requires 3 kinds of scripts or executables,

* **check** - Detecting new versions of the resource (e.g. git version)
* **in** - Fetching something
* **out** - Updating something

You build your resource with a Dockerfile by using the
`concourse docker base image` and adding your scripts/executables to `/opt/resource`.

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

I am using bash shell scripts to build the resource docker image. Using go is still in development.

To
[build.sh](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/build/build.sh)
using the
[Dockerfile](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/build/Dockerfile),

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

Now we can test the resource in a concourse pipeline.

To [set-pipeline.sh](https://github.com/JeffDeCola/concourse-resource-template/blob/master/test-this-resource/set-pipeline.sh)
for concourse,

```bash
cd test-this-resource
sh set-pipeline.sh.
```

Where the
[pipeline.yml](https://github.com/JeffDeCola/concourse-resource-template/blob/master/test-this-resource/pipeline.yml),

```yml
#------------------------------------------------------------------------------------------
jobs:

#**********************************************
- name: job-test-concourse-resource-template
#**********************************************
  plan:
    # CONCOURSE RESOURCE TEMPLATE
    - get: concourse-resource-template-test
      trigger: true

    # RUN TASK IN REPO USING ALPINE DOCKER IMAGE
    - task: task-test-concourse-resource-template
      file: concourse-resource-template/test-resource/tasks/task-test-concourse-resource-template.yml

      # TASK SUCCESS
      on_success:
        do:
          # CONCOURSE RESOURCE TEMPLATE
          - put: concourse-resource-template-test
            params:
              param1: "Hello jeff"
              param2: "How are you?"

#------------------------------------------------------------------------------------------
resource_types:

  - name: jeffs-resource
    type: docker-image
    source:
      repository: jeffdecola/concourse-resource-template
      tag: latest

#------------------------------------------------------------------------------------------
resources:

  - name: concourse-resource-template-test
    type: jeffs-resource
    source:
      source1: "This is the info for source 1"
      source2: "This is the info for source 2"
```
