terraform {
  required_providers {
    tftest = {
      version = "0.1.0"
      source  = "local/prashantv/tftest"
    }
  }
}

resource "tftest_dummy" "d" {
  job {
    name = "ignored"
    q = "foo "
  }
}

