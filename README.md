# Task 5: Creating a Builder

Welcome to the Paketo Buildpacks tutorial.

This repository consists of a set of branches,
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
- Task 5: **Creating a Builder**
- Task 6: Rolling Your Own Implementation Buildpack
- Task 7: The Packit Library

## Prerequisites

For this task you will need a couple additional pieces of software
 - [pack](https://buildpacks.io/docs/install-pack/)
   - This is the CLI that orchestrates the running of each Paketo buildpack
 - [docker](https://docs.docker.com/get-docker/)
   - Provides an image registry on all platforms.
 
 ## Creating a Builder
 
 Here we are going to write the configuration files needed to create our own custom **Builder**. Recall a builder is a combination of buildpacks, and a stack. For this exercise we are going to use the `bionic` stack, which is maintained by the [Cloud Native Buildpacks Project](https://buildpacks.io/).
 
 We will be using the following `pack` command to create a builder:
 ```
 pack create-builder <builder-name> -c <path/to/builder.toml>
 ```
 
 The general form of our `builder.toml` file is:
 
 ``` toml
 description = "Onboarding Builder"

[[buildpacks]]
image = "gcr.io/paketo-buildpacks/node-engine:1.2.3"
version = "1.2.3"

[[buildpacks]]
image = "gcr.io/<image-org>/<image-name>:<version>"
version = "<version>"


[[order]]
  [[order.group]]
  id = "paketo-buildpack/node-engine"
  version = "1.2.3"

  [[order.group]]
  id = "<id from buildpack.toml>"
  version = "<version>"

# Stack that will be used by the builder
[stack]
id = "io.buildpacks.stacks.bionic"
run-image = "gcr.io/paketo-buildpacks/run:base-cnb"
build-image = "gcr.io/paketo-buildpacks/build:base-cnb"
 ```
 
 ## Task
 
For this task we would like you to do the following,

1) Create a builder using the `pack create-builder` command. This builder should only contain the `node-engine` and `npm` buildpacks.
2) run a successful `pack build` of the `sample_application` using the builder created in step 1.
3) Create a builder using the `pack create-builder` command. This builder should only contain the `nodejs` metabuildpack.
4) run a successful `pack build` of the `sample_application` using the builder created in step 3.



Hints:
-
-  Again if stuck check out the branch named `creating_a_builder_solution` for a walk through.

