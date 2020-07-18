# Task 6: Rolling Your Own Implementation Buildpack

Welcome to the Paketo Buildpacks tutorial.

This repository consists of a couple of branches,
Each branch has a particular task designed to familiarize
you with the Paketo Buildpacks project.

Each subsequent task gets progressively more technical,
so feel free to dive as deep as you would like.

Index:
-
- Task 0: Introduction & The Application Image
- Task 1: Paketo Artifacts Overview
- Task 2: Packaging a Buildpack
- Task 3: Packaging a Metabuildpack
- Task 4: Create a custom Metabuildpack
- Task 5: Creating a Builder
- Task 6: **Rolling Your Own Implementation Buildpack**
- Task 7: The Packit Library

## Prerequisites

For this task you will need a couple additional pieces of software
 - [pack](https://buildpacks.io/docs/install-pack/)
   - This is the CLI that orchestrates the running of each Paketo buildpack
 - [docker](https://docs.docker.com/get-docker/)
   - Provides an image registry on all platforms.
- [`go`](https://golang.org/dl/)
    - the go programming language!
- the BSD `tar` binary


## Rolling Your Own Implementation Buildpack
Here we are going to be writing our very own Buildpack using the `go` programming language.

Before we strike out, let's take a quick look at the structure of a Buildpack.

A buildpack must contain at least three files arranged as follows:

```
buildpack-root
├── bin
│   ├── build
│   └── detect
└── buildpack.toml
```

The two files in the `bin` folder must be executable.

`pack` executes each of these binaries in two separate phases.

- First all `detect` binaries are run for all buildpacks,
- Second the `build` binaries are run for the buildpacks in the highest priority `group` that passed detection.

The goal of this task is to familiarize yourself with the inputs and outputs of these two binaries.

## Task

In these repositories there are `build` and `detect` directories. The code in each of these is responsible for building the `detect` and `build` binaries for our buildpack.

There are two test files that contain a list of tests at `./detect/detect_test.go` and `./build/build_test.go`.  Please read these test files, they have context about what you are attempting to implement.

Complete the `./detect/detect.go` and `./build/build.build.go` implementation and get all tests to pass.

These tests may be run using the following from the root of this repository
- `go test ./detect/... -count=1 -v`
- `go test ./build/... -count=1 -v`

If you get these tests to pass. You will have a bare bones buildpack!

Useful resources:
- [buildpack specification](https://github.com/buildpacks/spec/blob/main/buildpack.md)
- [detect interface](https://github.com/buildpacks/spec/blob/main/buildpack.md#detection)
- [detect process outline](https://github.com/buildpacks/spec/blob/main/buildpack.md#phase-1-detection)
- [build interface](https://github.com/buildpacks/spec/blob/main/buildpack.md#build)
- [build process outline](https://github.com/buildpacks/spec/blob/main/buildpack.md#phase-3-build)
- [golang TOML library](https://godoc.org/github.com/BurntSushi/toml)
- [golang JSON library](https://golang.org/pkg/encoding/json/)

It is recommended that you start with the `detect` tests.

## Part 2: Validation
Upon completion the `./scripts/package.sh` will compile the `detect` and `build` binaries
into a file at `./buildpack.tgz`

We can then use this buildpack in a `pack build` by running

```
pack build <name-of-app-image> -p <path_to_onboarding_sample_app> --buildpack <path_to_buildpack.tgz>
```

Run our application image to make sure that it is working:
```
docker run -d --rm -p 8080:8080 my-buildpack-test 'node server.js'
```

you should now be able to curl our application and get some output
```
curl localhost:8080
```
or visit `localhost:8080` in your browser

## Cleanup

Awesome job completing Task K!

- please remember to kill the container running your application
by running `docker kill <container-id>`








