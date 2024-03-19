

run: build_native gen
    go run --tags ij,jp .

build_native:
    cd lib && cargo build --release

gen:
    go generate .

build-docker:
    docker build -t hanamaru:latest .