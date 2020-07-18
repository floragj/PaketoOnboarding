# Task 2: Packaging a Buildpack

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
- Task 2: **Packaging a Buildpack**
- Task 3: Packaging a Metabuildpack
- Task 4: Create a custom Metabuildpack
- Task 5: Creating a Builder
- Task 6: Rolling Your Own Implementation Buildpack
- Task 7: The Packit Library

## Packaging a Buildpack

Ok, this is the first task where we will ask you the reader to complete a task, with a few hints (hopefully helpful). We will be trying to accomplish the following.

Build the sample application using the versions of the `node-engine` and `npm` buildpacks that you package, up from source.

## Prerequisites

For this task you will need a couple additional pieces of software
 - [pack](https://buildpacks.io/docs/install-pack/)
   - This is the CLI that orchestrates the running of each Paketo buildpack
 - [docker](https://docs.docker.com/get-docker/)
   - Provides an image registry on all platforms.
 - [sample_application](https://github.com/dwillist/onboarding_application)
   - just a simple application nodejs application we are going to build
   -    This app will be used throughout this tutorial so it is recommended that you use it
- [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
    - likely you system has this installed
- [paketo-buildpacks/node-engine repository](https://github.com/paketo-buildpacks/node-engine)
    - this can be clone by running
    ```git clone https://github.com/paketo-buildpacks/node-engine```
- [paketo-buildpacks/npm repository](https://github.com/paketo-buildpacks/npm)
    - this can be clone by running
    ```git clone https://github.com/paketo-buildpacks/npm```

## Packaging a Buildpack

For this task we would like you to do the following,
 1) run a successful `pack build` of the `sample_application` using the `node-engine` and `npm` **buildpack archives** that you package.

 2) run a successful `pack build` of the `sample_application` using the `node-engine` and `npm` **buildpackage images** that you package.

Hints:
-
- there is a script in both the `node-engine` and `npm` repos called `package.sh`, you should probably use this
- Recall order of buildpacks matter when running `pack build`
- If you are truly stuck, there is a secondary branch in this repo named `packaging_a_buildpack_solution`, the `README.md` file in this repo will contain a process to get these files

Extra:
-
- Try and use [`skopeo`](https://github.com/containers/skopeo) to import the `buildpackage.cnb` artifacts into your local docker registry. Then build using these.



