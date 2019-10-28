secret {
    no_prefix = true
    path = "SECRET_PATH"
}

vault {
  address     = "VAULT_ADDR"
  token       = "VAULT_TOKEN"
  renew_token = false
  unwrap_token = true
  ssl {
    enabled = true
    verify  = false
  }
}