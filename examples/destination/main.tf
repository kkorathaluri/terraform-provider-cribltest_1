terraform {
  required_providers {
    cribl-terraform = {
      source = "kkorathaluri/cribltest-new"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
  server_url ="https://app.cribl-playground.cloud/organizations/beautiful-nguyen-y8y4azd/workspaces/main/app/api/v1/m/default"
    bearer_auth = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjhIb3ctRnFoOUZ6Y3U4cGotd1V2OCJ9.eyJodHRwczovL2NyaWJsLmNsb3VkL3JvbGUiOlsiYWRtaW4iLCJvcmdfYWRtaW4iXSwiaHR0cHM6Ly9jcmlibC5jbG91ZC9vcmdhbml6YXRpb25JZCI6ImJlYXV0aWZ1bC1uZ3V5ZW4teTh5NGF6ZCIsImh0dHBzOi8vY3JpYmwuY2xvdWQvb3JnYW5pemF0aW9uIjp7Im5hbWUiOiJiZWF1dGlmdWwtbmd1eWVuLXk4eTRhemQifSwiaHR0cHM6Ly9jcmlibC5jbG91ZC9lbWFpbCI6Imtrb3JhdGhhbHVyaUBjcmlibC5pbyIsImh0dHBzOi8vY3JpYmwuY2xvdWQvbmFtZSI6Imtrb3JhdGhhbHVyaUBjcmlibC5pbyIsImh0dHBzOi8vY3JpYmwuY2xvdWQvY2xpZW50UmVxdWVzdCI6dHJ1ZSwiaXNzIjoiaHR0cHM6Ly9sb2dpbi5jcmlibC1wbGF5Z3JvdW5kLmNsb3VkLyIsInN1YiI6IlN2Q0NZaFAzZFdDUmNFR21PY0R2SzN0dHJIaUY4cWFoQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2FwaS5jcmlibC1wbGF5Z3JvdW5kLmNsb3VkIiwiaWF0IjoxNzQzNDg3NjYyLCJleHAiOjE3NDM1NzQwNjIsInNjb3BlIjoidXNlcjpyZWFkOmNsaWVudHMgdXNlcjpjcmVhdGU6Y2xpZW50cyB1c2VyOnVwZGF0ZTpjbGllbnRzIHVzZXI6cmVhZDp3b3JrZXJncm91cHMgdXNlcjp1cGRhdGU6d29ya2VyZ3JvdXBzIHVzZXI6cmVhZDpjb25uZWN0aW9ucyB1c2VyOmNyZWF0ZTpjb25uZWN0aW9ucyB1c2VyOnVwZGF0ZTpjb25uZWN0aW9ucyB1c2VyOmRlbGV0ZTpjb25uZWN0aW9ucyB1c2VyOnVwZGF0ZTp3b3Jrc3BhY2VzIHVzZXI6cmVhZDp3b3Jrc3BhY2VzIHVzZXI6ZGVsZXRlOndvcmtzcGFjZXMgdXNlcjpjcmVhdGU6d29ya3NwYWNlcyIsImd0eSI6ImNsaWVudC1jcmVkZW50aWFscyIsImF6cCI6IlN2Q0NZaFAzZFdDUmNFR21PY0R2SzN0dHJIaUY4cWFoIiwicGVybWlzc2lvbnMiOlsidXNlcjpyZWFkOmNsaWVudHMiLCJ1c2VyOmNyZWF0ZTpjbGllbnRzIiwidXNlcjp1cGRhdGU6Y2xpZW50cyIsInVzZXI6cmVhZDp3b3JrZXJncm91cHMiLCJ1c2VyOnVwZGF0ZTp3b3JrZXJncm91cHMiLCJ1c2VyOnJlYWQ6Y29ubmVjdGlvbnMiLCJ1c2VyOmNyZWF0ZTpjb25uZWN0aW9ucyIsInVzZXI6dXBkYXRlOmNvbm5lY3Rpb25zIiwidXNlcjpkZWxldGU6Y29ubmVjdGlvbnMiLCJ1c2VyOnVwZGF0ZTp3b3Jrc3BhY2VzIiwidXNlcjpyZWFkOndvcmtzcGFjZXMiLCJ1c2VyOmRlbGV0ZTp3b3Jrc3BhY2VzIiwidXNlcjpjcmVhdGU6d29ya3NwYWNlcyJdfQ.ahyrtFZEEz9UF3KVUnWskziVdOob36_gpkzMWcr5YUNHFzMIoTslzHhkBVXg6ctFxD8zPbPG851_sFHJ-QVombBh9-7penSB-ordn2catTVbZhm5wmd66sUPgFSF5GJlRNe39bwe2bXU37NQAc2uS6OW-P4oklhCPE6xcy6em44m1h_Ltko31mTrEKJ58upyb6vb-EE6GZDRAyZHbvVwrDePNwVoBkdh6c1Nh4ikjOQo1r-9V19hx-gWX85QThLkFmtvQOFJZ6Xr7zhK7WWXVNtpD0KlmQ0i6-CQDQOznW_SUK5WBeYgbXvVN0DzjL_5TdbjsVn3hn05gf01haHVqw"
}


resource "cribl-terraform_destination" "my_destination" {
   output_cribl_http = {
    compression            = "gzip"
    concurrency            = 11.56
    description            = "test"
    dns_resolve_period_sec = 5488.51
    environment            = "test"
    exclude_fields = [
      "test"
    ]
    exclude_self = true
    extra_http_headers = [
      {
        name  = "test"
        value = "test"
      }
    ]
    failed_request_logging_mode   = "none"
    flush_period_sec              = 2.35
    id                            = "test"
    load_balance_stats_period_sec = 12.74
    load_balanced                 = false
    max_payload_events            = 6.83
    max_payload_size_kb           = 2827.75
    on_backpressure               = "block"
    pipeline                      = "default"
    pq_compress                   = "none"
    pq_controls = {
      # ...
    }
    pq_max_file_size                  = 1000
    pq_max_size                       = 1000
    pq_mode                           = "always"
    pq_on_backpressure                = "drop"
    pq_path                           = "/tmp/test.pq"
    reject_unauthorized               = true
    response_honor_retry_after_header = false
    response_retry_settings = [
      {
        backoff_rate    = 18.82
        http_status     = 531.39
        initial_backoff = 117330.79
        max_backoff     = 52596.23
      }
    ]
    safe_headers = [
      "authorization"
    ]
    streamtags = [
      "test"
    ]
    system_fields = [
      "test"
    ]
    timeout_retry_settings = {
      backoff_rate    = 19.69
      initial_backoff = 173747.34
      max_backoff     = 163675.23
      timeout_retry   = true
    }
    timeout_sec = 8157421429180492
    tls = {
      ca_path             = "test"
      cert_path           = "test"
      certificate_name    = "test"
      disabled            = false
      max_version         = "TLSv1.3"
      min_version         = "TLSv1.3"
      passphrase          = "test"
      priv_key_path       = "/tmp/test.key"
      reject_unauthorized = false
      servername          = "test"
    }
    token_ttl_minutes = 38.63
    type              = "cribl_http"
    url               = "https://test.com"
    urls = [
      {
        url    = "https://test.com"
        weight = 0.48
      }
    ]
    use_round_robin_dns = false
  }
}

output "destination_details" {
  value = {
    id = cribl-terraform_destination.my_destination.output_default.id
    type = cribl-terraform_destination.my_destination.output_default.type
    default_id = cribl-terraform_destination.my_destination.output_default.default_id
    environment = cribl-terraform_destination.my_destination.output_default.environment
    pipeline = cribl-terraform_destination.my_destination.output_default.pipeline
  }
}