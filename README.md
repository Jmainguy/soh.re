# soh.re
The soh.re website infra

## Run this yourself
You can run this yourself on Google Cloud Platform using the included terraform scripts.

This expects you to have populated ~/.secure/gcp.json with your gcp credentials, please follow https://cloud.google.com/community/tutorials/getting-started-on-gcp-with-terraform if you have never used terraform to understand how to generate these files.

This also expects you named your gcp project "vpsaddict-router" and that the ID matches
```
cd terraform/gce
terraform init
terraform apply
```
