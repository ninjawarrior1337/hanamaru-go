

run: build_native gen
    go run --tags ij,jp .

build_native:
    cd lib && cargo build --release

gen:
    go generate .

build_docker:
    docker build -t ghcr.io/ninjawarrior1337/hanamaru-go:main .