# Task 4: Creating a Metabuildpack

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
- Task 4: **Create a custom Metabuildpack**
- Task 5: Creating a Builder
- Task 6: Rolling Your Own Implementation Buildpack
- Task 7: The Packit Library

## Prerequisites

For this task you will need a couple additional pieces of software
 - [pack](https://buildpacks.io/docs/install-pack/)
   - This is the CLI that orchestrates the running of each Paketo buildpack
 - [docker](https://docs.docker.com/get-docker/)
   - Provides an image registry on all platforms.
 - [sample_application](https://github.com/dwillist/onboarding_application)
   - just a simple application nodejs application we are going to build
   -    This app will be used throughout this tutorial so it is recommended that you use it
- the BSD `tar` binary, everything will work with `gnutar`, though some flags may differ.


## Packaging a Metabuildpack

We are going to write a metabuildpack definition and use `pack` to turn it into a functional metabuildpack.


We are going to need two files to do this, they are the `package.toml` file and a `metabuildpack.tgz` archive. Let's take a closer look at these.

#### `metabuildpack.tgz`
This is an archive containing a single `buildpack.toml` file at the archive root. This `buildpack.toml` file contains the ordering and group information for the metabuildpack.

Here is an rough example of the `buildpack.toml` file in the `paketo-buildpack/nodejs`

``` toml
api = "0.2"

[buildpack]
  id = "paketo-buildpacks/nodejs"
  name = "Node.js Buildpack"
  version = "1.2.3"

[[order]]

  [[order.group]]
    id = "paketo-buildpacks/node-engine"
    version = "1.2.3"

  [[order.group]]
    id = "paketo-buildpacks/yarn-install"
    version = "1.2.3"

[[order]]

  [[order.group]]
    id = "paketo-buildpacks/node-engine"
    version = "1.2.3"

  [[order.group]]
    id = "paketo-buildpacks/npm"
    version = "1.2.3"

```

Notice how the `orders` and `groups` in the above `toml` corresponds to the image used in Task 0.

<img src="assets/metabuildpackage.png">

#### `package.toml`

This is a file, consumed by the `pack package-buildpack` command. It basically tells pack exactly where to find all of the buildpackages `buildpack.toml` file above, as well as a `uri` to the `metabuildpackage.tgz` archive.

This file has contents similar to the following.
``` toml
[buildpack]
  uri = "path/to/metabuildpack.tgz"

[[dependencies]]
  image = "gcr.io/paketo-buildpacks/node-engine:0.0.237"

[[dependencies]]
  image = "some/image/reference:version"


```

## Task

For this task we would like you to do the following,

1) Create a `metabuildpack.tgz` archive that references the `node-engine` and `npm` implementation buildpacks.
2) Create a `package.toml` file that uses the `metabuildpackage.tgz` created in step 1, and the `node-engine` and `npm` buildpackages.
3) Create a custom metabuildpackage using the `pack package-buildpack` command with the `package.toml` created in step 2 
4) Run a successful `pack build` of the `sample_application` using the metabuildpackage created in step 3.

Hints:
-
- check out the current `nodejs metabuildpack` [`package.toml` file](https://github.com/paketo-buildpacks/nodejs/blob/master/package.toml)
-  check out the current `nodejs` metabuildpack [`buildpack.toml` file](https://github.com/paketo-buildpacks/nodejs/blob/master/buildpack.toml)
-  Again if really stuck check out the branch named `creating_a_metabuildpack_solution` for a walk through.

