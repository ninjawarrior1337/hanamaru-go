run:
    go run --tags ij,jp .

gen:
    go generate .

build-docker:
    docker build -t hanamaru:latest .