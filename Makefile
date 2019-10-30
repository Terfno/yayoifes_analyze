init:
	go install

dls:
	@echo ""
	docker images
	@echo ""
	docker ps -a
	@echo ""
