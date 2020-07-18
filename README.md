# Task 0: Introduction and The Application Image

Welcome to the Paketo Buildpacks tutorial.

This repository consists of a set of branches,
Each branch has a particular task designed to familiarize
you with the Paketo Buildpacks project.

Each subsequent task gets progressively more technical,
so feel free to dive as deep as you would like.

Index:
-
- Task 0: **Introduction and The Application Image**
- Task 1: Paketo Artifacts Overview
- Task 2: Packaging a Buildpack
- Task 3: Packaging a Metabuildpack
- Task 4: Create a custom Metabuildpack
- Task 5: Creating a Builder
- Task 6: Rolling Your Own Implementation Buildpack
- Task 7: The Packit Library


Without further ado let's jump in.

## The Application image

**Application Images** are the output artifact that the Paketo Buildpacks project builds.

For our purposes an application image is just a collection of layers. Each layer performs some addition (or deletion) to the filesystem. Each application image has three types of layers, an **app layer**, **dependency layers**, and an **OS layer**.


<img src="assets/app_image.png" width="200">

Notice the topology of the above image implies dependence. The application layer is dependent on the dependencies layers, and on the OS layer. While the dependencies layers are just reliant on other lower dependency layers and the OS layer.

The Paketo Buildpacks are responsible for providing the dependency layers, while the Paketo Stacks are responsible for the OS layer.

With this image in mind, let's walk through the steps for building an application image using the artifacts the Paketo Buildpacks project produces.

## Prerequisites

For this task we need a couple additional pieces of software
 - [pack](https://buildpacks.io/docs/install-pack/)
   - This is the CLI that orchestrates the running of each Paketo buildpack
 - [docker](https://docs.docker.com/get-docker/)
   - Provides an image registry on all platforms.
 - [sample_application](https://github.com/dwillist/onboarding_application)
   - just a simple application nodejs application we are going to build
   - This app will be used throughout this tutorial so it is recommended that you use it


## Task

We are going to build an app-image using `pack` three different ways
1. Using a **builder**
1. Using a **metabuildpackage**
1. Using **implementation buildpackages**


##  Using a Builder
Once you have installed the above prerequisites and started the docker daemon. We can now set up `pack` to use a Paketo Buildpacks Builder.

list all recommended builders:

```
pack list-trusted-builders
```

For this tutorial we are going to use the `gcr.io/paketo-buildpacks/builder:base` builder.
To set this as the default we run

```
pack set-default-builder gcr.io/paketo-buildpacks/builder:base
```

Ok great! Now from the root of the `sample_application` repository:
```
pack build onboarding-test-image
```

After a bit of output, our build succeeds and we have produced an application image. Run:
```
docker images
```
lists all images and we will see that indeed the `onboarding-test-image` application image is present.

## Using a metabuildpackage

A **metabuildpackage** is an artifact consisting of a group of buildpackages and an order that they should be executed in.

We are going to `pack build` using the `gcr.io/paketo-buildpacks/nodejs` **metabuildpackage**
This buildpackage contains the `node-engine`, `npm` and `yarn` implementation buildpackages.

from the root of the `sample_application` run:
```
pack build metabuildpack-build-test --buildpack gcr.io/paketo-buildpacks/nodejs
```

## Using implementation buildpackages

And finally we can un-group the **metabuildpackages** used in the above `pack build` build and use the **implementation buildpacages** directly.

Note: ordering of the implementation buildpackages matters
```
pack build implementation-build-test --buildpack gcr.io/paketo-buildpacks/node-engine --buildpack gcr.io/paketo-buildpacks/npm
```
