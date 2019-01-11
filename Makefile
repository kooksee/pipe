.PHONY: version build build_linux docker_login docker_build docker_push_dev docker_push_pro
.PHONY: rm_stop

test:
	@echo "test"
	go test -v .
