# Projects
This is a incomplete list of projects and tools I have worked on, for a full list you should visit https://github.com/jmainguy

## This website

This website uses gotty (for the interactive shell), docker (to contain it, keep people from breaking things), iptables (to limit network), haproxy (to intercept http and websocket connections seperate), a custom router I am calling soh-router for now (this intercepts websocket requests as passed by haproxy, then manages docker containers, then sends the socket to an available connection), sqlite3 (to store the available connections), systemd (to run the router as a service)

[https://github.com/Jmainguy/soh.re](https://github.com/Jmainguy/soh.re "This Website")

## Ansible
I am proud to be a module maintainer and regular contributor to Ansible. Ansible is a radically simple IT automation system. It handles configuration-management, application deployment, cloud provisioning, ad-hoc task-execution, and multinode orchestration

## Bible
A command line bible. It reads from an sqlite3 database. Supports multiple translations. I spend most my time in the command line, and figured having a bible available there would be useful.

This is installed on this system for you to use.

[https://github.com/Jmainguy/bible](https://github.com/Jmainguy/bible "Bible")

## Repeatafterme
A golang daemon to like / retweet anything a list of users post. Useful for important twitter users to you which you think all your followers should also know about.

[https://github.com/Jmainguy/repeatafterme](https://github.com/Jmainguy/repeatafterme "Repeatafterme")

## Bak
bak is a single file/directory backup utility I wrote in python. The idea came about after taking a training course where the recommended we run really long commands to essentially cp a file and append .bak to the end, bak provides this and more in a much shorter and easier to remember command.

This is installed on this system for you to use.

[https://github.com/jmainguy/bak](https://github.com/jmainguy/bak "Bak")

## Indent
indent is a cli tool to add or remove spaces from a file, and output it to stdout. I wrote this to assist with writing helm / argoCD applications, where I was often adding or removing indentation from a kubernetes yaml file to match expected indentation for helm / argocd.

This is installed on this system for you to use.

[https://github.com/jmainguy/indent](https://github.com/jmainguy/indent "Indent")

## CertificateDownloader
certificatedownloader is a cli tool to download a cert from a https site, and optionally on a fedora/rhel/centos server, have the OS trust it by default (needs sudo, which you dont have here). I wrote this tool for assisting in setting up a workspace for interacting with a newly created kubernetes cluster, before certManager had been run to give it valid certs. Although this tool can be quite useful for an intranet sites your company might be using.

This is installed on this system for you to use.

[https://github.com/jmainguy/certificatedownloader](https://github.com/jmainguy/certificatedownloader "CertificateDownloader")

## czechDNS
czechdns is a cli tool to test a number of dns records, against a specific list of nameservers. I wrote this tool to provide trust to my team when we were migrating a large number of records from one amazon R53 zone to multiple. I first created the test set with a list of hostnames and hit their current zone. Then ran the test set against the new zones. Once trust was built, we flipped the switch to send traffic to the new zones.

This is installed on this system for you to use. See https://github.com/Jmainguy/czechdns/blob/main/example.yaml for an example testFile to use.

[https://github.com/jmainguy/czechdns](https://github.com/jmainguy/czechdns "czechDNS")

