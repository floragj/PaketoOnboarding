# Task 3: Packaging a Metabuildpack

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
- Task 3: **Packaging a Metabuildpack**
- Task 4: Create a custom Metabuildpack
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
- [paketo-buildpacks/nodejs repository](https://github.com/paketo-buildpacks/nodejs)
    - this can be clone by running
    ```git clone https://github.com/paketo-buildpacks/nodejs```



## Packaging a Metabuildpack

This is a simple continuation from the last exercise, where we are going to package up a `nodejs` metabuildpackage.

For this task we would like you to do the following,
 1) package the `nodejs` metabuildpack from source into a buildpackage
 2) run a successful `pack build` of the `sample_application` using the `nodejs` **metabuildpackage** that you packaged in step 1.

Hints:
-
- there is a script in the `nodejs` repo called `package.sh`, try using it.
- Again if stuck check out the branch named `packaging_a_metabuildpack_solution`

