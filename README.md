Terraform Provider Pritunlko
============================

Unofficial Terraform provider for the [Pritunl VPN server](https://pritunl.com) because an [official one](https://github.com/pritunl/terraform-provider-pritunl) is not really functional<br/><br/>

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="400px">
<img src="https://files.readme.io/VrFcaFRleaC8nYnHZp4Q_logo_full2.png" width="400px">

Maintainers
-----------

[vi7](http://github.com/vi7)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x (earlier versions may work as well)
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

Usage
---------------------

### TL;DR

```
provider "pritunlko" {
  pritunl_host = "192.168.200.20"
  pritunl_token = "ey0h5PtLZpCvpiWcOpwoCWYdu5rdR3XT"
  pritunl_secret = "MVKcPmv3H5E2Xc5sdn0cZPOX8ARTbUEl"
  # Optional:
  pritunl_insecure = false  # Default: `false`
}

resource "pritunlko_organization" "tf_org" {
  name = "tf-org02"
}

```

also check [examples](examples)

### Provider arguments

- `pritunl_host` - Pritunl hostname or IP address
- `pritunl_token` - API token
- `pritunl_secret` - API secret
- `pritunl_insecure` - (Optional) Skip HTTPS cert check. Useful when Pritunl endpoint uses self-signed cert. Default: `false`

### Resources

### pritunlko_organization

Creates organization

**Arguments**

- `name` - organization name

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/vi7/terraform-provider-pritunlko`

```sh
$ mkdir -p $GOPATH/src/github.com/vi7; cd $GOPATH/src/github.com/vi7
$ git clone git@github.com:vi7/terraform-provider-pritunlko
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/vi7/terraform-provider-pritunlko
$ make build
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-pritunlko
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

TODO list
---------

- Implement resources for other Pritunl API objects except Organization
- Implement datasources
- Implement resource imports
