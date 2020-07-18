# Task 0: Introduction and The Application Image

Welcome to the Paketo Buildpacks tutorial.

This repository consists of a set of branches, 
Each branch has a particular task designed to framiliarize 
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


Without further ado lets jump in.

## The Application image

**Application Images** are the output artifact that the Paketo Buildpacks project builds.

For our purposes an application image is just a collection of layers. Each layer performs some addition (or deletion) to the filesystem. Each application image has three types of layers, an **app layer**, **dependency layers**, and an **OS layer**.


![](assets/app_image.png =300x)

Notice the topology of the above image implies dependence. The application is dependent on the dependencies layers, and on the OS layer. While the dependencies layers are just relient on other lower dependency layers and the OS layer.

The Paketo Buildpacks are responsible for providing the dependency layers, while the Paketo Stacks are responsible for the OS layer.

With this image in mind lets walk through the steps for building an application image using the artifacts that the Paketo Buildpacks project produces.

## Prerequisites

For this task you will need a couple additional peices of software
 - [pack](https://buildpacks.io/docs/install-pack/)
   - This is the CLI that orchestrates the running of each Paketo buildpack
 - [docker](https://docs.docker.com/get-docker/)
   - Provides a image registry on all platforms.
 - [sample_application](https://github.com/dwillist/onboarding_application)
   - just a simple application nodejs application we are going to build 
   - This app will be used throughout this tutorial so it is recommended that you use it


## Task

We are going to is just build the same application into an app-image using `pack` in three different ways
1. Build using a **builder**
1. Build using a **metabuildpackage**
1. Build using **implementation buildpackages**


##  Using a Builder
Given that you have installed the above pre-requisites and started the docker daemon. We can now set up pack to use a Paketo Buildpacks Builder.

list all recommended builders:

```
pack list-trusted-builder
```

For this tutrorial we are going to use the `gcr.io/paketo-buildpacks/builders:base` builder.
To set this as the default we run

```
pack set-default-builder gcr.io/paketo-buildpacks/builders:base
```

Ok great! Now from the root of the `sample_application` cloned as a prerequisite simply run:
```
pack build onboarding-test-image
```

After a bit of output our build succeeds and we have produced an application image. Running 
```
docker images
```
lists all images and we will see that indeed the `onboarding-test-image` application image is present.

## Using a metabuildpackage

A **metabuildpackage** is a artifact consisting of a group of buildpackages and an order that they should be excuted in.

We `pack build` using a specific **metabuildpackage** 

This time lets build our application using the `gcr.io/paketo-buildpacks/nodejs` metabuildpackage, which contains the `node-engine`, `npm` and `yarn` implementation buildpackages. 

this is as simple as sitting in the root of the `sample_application`
```
pack build metabuildpack-build-test --buildpack gcr.io/paketo-buildpacks/nodejs
```

## Using implementation buildpackages

And finally we can un-group the **implementation buildpackages** used in the above `pack build` build and use them directly.

Note: ordering of the implementation buildpackages matters
```
pack build implementation-build-test --buildpack gcr.io/paketo-buildpacks/nodejs --buildpack gcr.io/paketo-buildpacks/npm
```


