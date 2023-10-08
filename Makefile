ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CHDIR_SHELL := $(SHELL)
define chdir
   $(eval _D=$(firstword $(1) $(@D)))
   $(info $(MAKE): cd $(_D)) $(eval SHELL = cd $(_D); $(CHDIR_SHELL))
endef

.PHONY: all
all : buildtheme screenshots clean

.PHONY: buildtheme
buildtheme :
	cd mktheme; go run . ../SPEC.md; cd $(ROOT_DIR)

.PHONY: pkg
pkg :
	vsce package -o clarion.vsix


.PHONY: screenshots
screenshots : pkg
	docker run --rm -t \
		-v $(ROOT_DIR)/img:/home/codeuser/shots \
		-v $(ROOT_DIR)/clarion.vsix:/home/codeuser/extensions/clarion.vsix \
		-v $(ROOT_DIR):/home/codeuser/code \
		flowchartsman/vscode-docker-screenshots \
		./makeshots --samplefile mktheme/theme.go --fileline 273 \
		"Clarion White,Clarion-White.jpg" \
		"Clarion Blue,Clarion-Blue.jpg" \
		"Clarion Orange,Clarion-Orange.jpg" \
		"Clarion Peach,Clarion-Peach.jpg" \
		"Clarion Red,Clarion-Red.jpg"


.PHONY: clean
clean :
	rm clarion.vsix

.PHONY: install_local
install_local : pkg
	code --install-extension clarion.vsix
	rm clarion.vsix
