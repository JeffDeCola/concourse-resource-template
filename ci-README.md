# CONCOURSE CONTINUOUS INTEGRATION

I use concourse ci to,

* Copy and edit `README.md` to `/docs/_includes/README.md` for
  [GitHub Webpage](https://jeffdecola.github.io/concourse-resource-template/)
* TEST code
* BUILD docker image
* PUSH to dockerhub
* Alert me of the progress via repo status and slack

## PIPELINE

The concourse
[pipeline.yml](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - concourse-resource-template concourse ci pipeline - IMAGE](docs/pics/concourse-resource-template-pipeline.jpg)

## JOBS, TASKS AND RESOURCE TYPES

The concourse `jobs` and `tasks` are,

* `job-readme-github-pages` runs task
  [task-readme-github-pages.yml](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci/tasks/task-readme-github-pages.yml)
  that kicks off shell script
  [readme-github-pages.sh](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci/scripts/readme-github-pages.sh)
* `job-build-push` runs task
  [task-build-push.yml](https://github.com/JeffDeCola/concourse-resource-template/blob/master/ci/tasks/task-build-push.yml)
  that kicks off shell script
  [build-push.sh](https://github.com/JeffDeCola/concourse-resource-template/tree/master/ci/scripts/build-push.sh)

The concourse `resources types` are,

* `concourse-resource-template` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github
* `resource-dump-to-dockerhub` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/docker-image-resource/)
  to PUSH a docker image to dockerhub.
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit

For more information on using concourse for continuous integration,
refer to my
[concourse-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet).
