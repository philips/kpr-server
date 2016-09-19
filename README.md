# KPR Server

Kubernetes Package Registry Server is a prototype server to create a
push/pull/login experience for Kubernetes Packages.

## Theory of Operation

These objects and URL layouts are closely inspired by the Open Container
Initiative and the Docker Registry API. The basic idea is that a user request
for a name and tag like `example.com/org/app:v1.0.0` is resolved in the
following order:

```
Named Repo -> Tag (or Channel to tag) -> Manifest List -> Manifest -> Package Blob(s)
```

Once a name and tag are resolved to a Manifest List all of the objects are
simply content-addressable blobs. This means that the manifest list, or any
other object, can be signed or otherwise addressed directly by their digest.
This is how the Open Container Iniative Image Specification works (see the
image-layout.md and manifest.md).

### Definitions

- A `package` is a blob (usually tar.gz) that Helm processes to install resources on a cluster. This is built by the developer.
- A `repo` name like example.com/org/app where a number of tagged packages live. A repo name is in the DNS federated namespace e.g. example.com/org/app
- A `repo tag` represnts a package release that resolves to a package `manifest` or `manifest list` by digest.
  - A `manifest` pointing to the package blob by digest. This is built and the digest is optionally signed by the developer.
  - A `manifest list` pointing to one or more manifests. This is generated and the digest is optionally signed by the developer. This indirection will allow for new versions of the package format or handling different formats like DAB.## gRPC + REST Gateway Play


## Trying it out

To try it all out do this:

```
$ go get -u github.com/philips/kpr-server
$ kpr-server serve
```
