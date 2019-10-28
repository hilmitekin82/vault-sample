# Uncomment this to have Agent run once (e.g. when running as an initContainer)
exit_after_auth = true
pid_file = "/home/vault/pidfile"

vault {
    tls_skip_verify = "true"
}

auto_auth {
    method "kubernetes" {
        mount_path = "auth/kubernetes"
        config = {
            role = "VAULT_ROLE"
        }
        wrap-ttl = 120
    }

    sink "file" {
        config = {
            path = "/home/vault/.vault-token"
        }
    }
}