# Define where your included makefile lives
GITURL = git.psu.edu/k8s/gitlab-ci-templates.git
GITPROJECT = swe-golang-binary
GITFILE = .go-project.mk

# Override your docker tag here, otherwise it'll use your PWD
# DOCKER_TAG = everything-is-awesome

-include $(shell test -f .go-project.mk || git archive --remote=ssh://git@$(GITURL) HEAD:$(GITPROJECT) $(GITFILE) | tar -x ; echo .go-project.mk)
