build:
	docker build -t go_yf:latest .
	docker images

gotest:
	docker run -it --name yf go_yf:latest /bin/bash

dls:
	@echo ""
	docker images
	@echo ""
	docker ps -a
	@echo ""
