# provider-cloudflare

`provider-cloudflare` is a [Crossplane][crossplane] provider that is built using
[Upjet][upjet] code generation tools and exposes XRM-conformant managed
resources for the Cloudflare API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release][marketplace]:

```console
up ctp provider install cloudhippie/provider-cloudflare:v1.0.0
```

Alternatively, you can use declarative installation:

```console
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudflare
spec:
  package: cloudhippie/provider-cloudflare:v1.0.0
EOF
```

You can see the API reference [here][apidocs].

## Developing

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build statically linked binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue][issue].

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2024 Cloudhippie <info@cloudhippie.de>
```

[crossplane]: https://crossplane.io/
[upjet]: https://github.com/crossplane/upjet
[marketplace]: https://marketplace.upbound.io/providers/cloudhippie/provider-cloudflare
[apidocs]: https://doc.crds.dev/github.com/cloudhippie/provider-cloudflare
[issue]: https://github.com/cloudhippie/provider-cloudflare/issues
