# atv-simulator
All-terrain vehicle simulator

How to build locally:
```bash
docker-compose up --build -d
```

How to run container locally:
```bash
docker exec -it atv-simulator bash
```

How to test locally: inside the container, run the following:
```bash
go run main.go
```