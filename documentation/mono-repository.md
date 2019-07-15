## Mono-Repository

All code and documentation for this project exists ENTIRELY within this repository -- except security credentials.

This repo is organized around Uncle Bob Martin's concept of being able to grasp what a project does, at a glance, from the names of items in the top level of the project.

While that's an oldish concept, this repo takes it one step further. Even the deployent and production configuration of this system exists in this repository.

Also, common developer tools that the team has found useful are in this repository. See atom/, git/, compatibility/, and load/ for examples of this philosophy.

An important premise to note is that all commands and scripts in this repository are root-relative, intended to be invoked only from the root of the repository. This makes it unnecessary to change directories during development, and allows all nested script calls to be relative to root, providing for very uniform and easy to follow script chains.
