CMDS=$(shell find ./cmd/* -maxdepth 1 -type d -exec basename {} \; )

cmd_targets = $(addprefix ./cmd/, $(CMDS))

all:
	go install -v $(cmd_targets)

publish:
	scripts/publish.sh

.PHONY: all publish
