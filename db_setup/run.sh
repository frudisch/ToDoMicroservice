#build docker container
docker build -t todo_go_example .
#run docker container
docker run --rm -P -p 127.0.0.1:5432:5432 --name todo_go_example todo_go_example
