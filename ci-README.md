# CONCOURSE CONTINUOUS INTEGRATION

I use concourse to automate,

* Use README for
  [GitHub Webpage](https://jeffdecola.github.io/jeffcoin/)
* TEST code
* Alert me of the progress via repo status and slack

## PIPELINE

The concourse
[pipeline.yml](https://github.com/JeffDeCola/jeffcoin/blob/master/ci/pipeline.yml)
shows the entire ci flow,

![IMAGE - jeffcoin concourse ci pipeline - IMAGE](docs/pics/jeffCoin-pipeline.jpg)

## JOBS, TASKS AND RESOURCE TYPES

Concourse Jobs and Tasks

* `job-readme-github-pages` runs task
  [task-readme-github-pages.yml](https://github.com/JeffDeCola/jeffcoin/blob/master/ci/tasks/task-readme-github-pages.yml)
  that kicks off shell script
  [readme-github-pages.sh](https://github.com/JeffDeCola/jeffcoin/blob/master/ci/scripts/readme-github-pages.sh)
* `job-unit-tests` runs task
  [task-unit-tests.yml](https://github.com/JeffDeCola/jeffcoin/blob/master/ci/tasks/task-unit-tests.yml)
  that kicks off shell script
  [unit-tests.sh](https://github.com/JeffDeCola/jeffcoin/tree/master/ci/scripts/unit-tests.sh)

Concourse Resources

* `jeffcoin` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/jeffdecola/github-status-resource-clone)
  that will update your git status for that particular commit
