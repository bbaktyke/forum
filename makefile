create:
	docker build -t beka .
run:
	docker run -p 8080:8080 --rm --name bbaktyke beka 
stop:
	docker stop bbaktyke
start:
	docker start  bbaktyke
prune:
	docker container prune