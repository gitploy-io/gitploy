# GITPLOY_TLS_CERT

Path to an SSL certificate used by the server to accept HTTPS requests. This parameter is of type string and is optional.

Please note that the cert file should be the concatenation of the server’s certificate, any intermediates, and the certificate authority’s certificate.

```
GITPLOY_TLS_CERT=/path/to/cert.pem
```