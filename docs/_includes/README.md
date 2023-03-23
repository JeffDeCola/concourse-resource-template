  _built with
  [concourse](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci-README.md)_

# OVERVIEW

A concourse resource is a docker image.

It requires 3 kinds of scripts or executables,

* **check** - Detecting new versions of the STUFF (e.g. git version)
* **in** - GET STUFF
* **out** - PUT STUFF

You build your resource with a Dockerfile by using the
`concourse docker base image` and adding your scripts/executables to `/opt/resource`.

The three scripts/executables can be written with bash, go, etc.

* My go is build is located in
  [/build-resource-using-go](https://github.com/JeffDeCola/concourse-resource-template/tree/master/build-resource-using-go)
  _(In development)_
* My bash build is located in
  [/build-resource-using-bash](https://github.com/JeffDeCola/concourse-resource-template/tree/master/build-resource-using-bash)

![IMAGE - concourse-check-in-out-pipeline - IMAGE](pics/concourse-check-in-out-pipeline.jpg)

---

### CHECK

[check](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/check-in-out/check)
is performed before anything can use the resource. It is used to
determine if the STUFF has changed (checks version of STUFF).

#### PART 1 - Input

Concourse will send **stdin** for `check` to parse, where the source
and comes from the pipeline and the version comes from the check.

```json
{
  "source": {
    "source1": "source1 info",
    "source2": "source2 info",
    "source3": "source3 info"
  },
  "version": {
    "ref": "null"
  }
}
```

#### PART 2 - Check/Update Version

In this example, I will mimic a getting a new version and increment until
version 5.

#### PART 3 - Output

Check will send **stdout** that will be used in the next step in the pipeline.

```json
[
  {
    "ref": "1"
  }
]
```

---

### IN

The
[in](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/check-in-out/in)
is performed after a check has confirmed there is something there.
For my resource,
[in](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/check-in-out/in)
will mimic **fetching STUFF** and place a file in the working directory.

#### PART 1 - Input

Concourse will send **stdin** for `in` to parse, where the source
and params come from the pipeline and the version comes from the check.

```json
{
  "params": {
    "param1": "get param1",
    "param2": "get param2",
    "param3": "get param3"
  },
  "source": {
    "source1": "source1 info",
    "source2": "source2 info",
    "source3": "source3 info"
  },
  "version": {
    "ref": "1"
  }
}
```

#### PART 2 - GET Something

In this example, I will mimic a fetch some STUFF and place a file
`get_fetch.json` in the working directory.

#### PART 3 - Output

Input will send **stdout** that will be used in the next step in the pipeline.

```json
{
  "version": {
    "ref": "1"
  },
  "metadata": [
    { "name": "author", "value": "Jeff DeCola"},
    { "name": "author_date", "value": "March 2023"},
    { "name": "executable", "value": "in"},
    { "name": "version", "value": "1" }
  ]
}
```

---

### OUT

The
[out](https://github.com/JeffDeCola/concourse-resource-template/blob/master/build-resource-using-bash/check-in-out/out)
will mimic **updating STUFF**
and is performed after the task.

#### PART 1 - Input

Concourse will send **stdin** for `out` to parse, where the source
and params come from the pipeline.

```json
{
  "params": {
    "param1": "put param1",
    "param2": "put param2",
    "param3": "put param3"
  },
  "source": {
    "source1": "source1 info",
    "source2": "source2 info",
    "source3": "source3 info"
  }
}
```

#### PART 2 - PUT Something

In this example, I will mimic a push/deploy to STUFF and place a file
`put_fetch.json` in the working directory.

It is important you must recheck version here. **So you must get the version.**

#### PART 3 - Output

You send **stdout** that will be used in the next step in the pipeline.

```json
{
  "version": {
    "ref": "1"
  },
  "metadata": [
    { "name": "author", "value": "Addie DeCola"},
    { "name": "author_date", "value": "April 2021"},
    { "name": "executable", "value": "out"},
    { "name": "version", "value": "1" }
  ]
}
```

## BUILD AND PUSH THE RESOURCE

I am using bash shell scripts to build the resource docker image.
Using go is still in development.

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

You can check this docker image,

```bash
docker images jeffdecola/concourse-resource-template
docker run --name concourse-resource-template -dit jeffdecola/concourse-resource-template
docker exec -i -t concourse-resource-template /bin/bash
cd /opt/resource
tree
docker logs concourse-resource-template
docker rm -f concourse-resource-template
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

    # GET REPO FROM GITHUB
    - get: concourse-resource-template
      trigger: true

    # CONCOURSE RESOURCE TEMPLATE
    - get: concourse-resource-template-test
      params:
        param1: "get param1"
        param2: "get param2"
        param3: "get param3"

    # RUN TASK IN REPO USING ALPINE DOCKER IMAGE
    - task: task-test-concourse-resource-template
      file: concourse-resource-template/test-this-resource/tasks/task-test-concourse-resource-template.yml

      # TASK SUCCESS
      on_success:
        do:
          # CONCOURSE RESOURCE TEMPLATE
          - put: concourse-resource-template-test
            params:
              param1: "put param1"
              param2: "put param2"
              param3: "put param3"

#------------------------------------------------------------------------------------------
resource_types:

  - name: jeffs-resource
    type: docker-image
    source:
      repository: jeffdecola/concourse-resource-template
      tag: latest

#------------------------------------------------------------------------------------------
resources:

  - name: concourse-resource-template
    type: git
    icon: github
    source:
      uri: git@github.com:jeffdecola/concourse-resource-template.git
      branch: master
      private_key: ((git_private_key))

  - name: concourse-resource-template-test
    type: jeffs-resource
    source:
      source1: "source1 info"
      source2: "source2 info"
      source3: "source3 info"
```

Note: You will need to put your `((git_private_key))` in a .credentials
file in the root of this repo.
