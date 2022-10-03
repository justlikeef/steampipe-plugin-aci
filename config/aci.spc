connection "aci" {
  plugin = "justlikeef/aci"

  # The url APIC lives at
  cluster_uri  = "192.168.122.233"

  # APIC username
  user  = "root"

  # APIC password
  password  = "s0Mep@ss"

  # APIC login Domain
  #login_domain = ""

  # TLS cert validation
  allow_unverified_ssl = true
}