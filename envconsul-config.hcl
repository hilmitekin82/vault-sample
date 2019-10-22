secret {
    no_prefix = true
    path = "SECRET_PATH"
}

vault {
  address     = "VAULT_ADDR"
  token       = "VAULT_TOKEN"
  renew_token = false

  ssl {
    enabled = true
    verify  = false
  }
}