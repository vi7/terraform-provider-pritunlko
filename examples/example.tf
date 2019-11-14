provider "pritunlko" {
  pritunl_host = "192.168.200.20"
  pritunl_token = "ey0h5PtLZpCvpiWcOpwoCWYdu5rdR3XT"
  pritunl_secret = "MVKcPmv3H5E2Xc5sdn0cZPOX8ARTbUEl"
  pritunl_insecure = true
}

resource "pritunlko_organization" "tf_org" {
  name = "tf-org02"
}
