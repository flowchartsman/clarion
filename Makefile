CHDIR_SHELL := $(SHELL)
define chdir
   $(eval _D=$(firstword $(1) $(@D)))
   $(info $(MAKE): cd $(_D)) $(eval SHELL = cd $(_D); $(CHDIR_SHELL))
endef

.PHONY: all
all : buildmktheme buildtheme

.PHONY: buildmktheme
buildmktheme :
	$(call chdir,mktheme)
	go build

.PHONY: buildtheme
buildtheme : buildmktheme
	$(call chdir,mktheme)
	./mktheme ../SPEC.md ..

.PHONY: clean
clean :
	$(call chdir,mktheme)
	go clean

.PHONY: install_local
inst_dir:=$(HOME)/.vscode/extensions/clarion_dev
install_local :
	mkdir -p $(inst_dir)
	cp -r icon.png package.json themes $(inst_dir)
	@echo "you may wish to restart vscode"
