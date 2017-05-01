# Projects
This is a list of projects, ordered by what I think is most recent and interesting that I have worked on.

## This website

This website uses gotty (for the interactive shell), docker (to contain it, keep people from breaking things), iptables (to limit network), haproxy (to intercept http and websocket connections seperate), a custom router I am calling soh-router for now (this intercepts websocket requests as passed by haproxy, then manages docker containers, then sends the socket to an available connection), sqlite3 (to store the available connections), systemd (to run the router as a service)

[https://github.com/Jmainguy/soh.re](https://github.com/Jmainguy/soh.re "This Website")

## Ansible
I am proud to be a module maintainer and regular contributor to Ansible. Ansible is a radically simple IT automation system. It handles configuration-management, application deployment, cloud provisioning, ad-hoc task-execution, and multinode orchestration

## Repeatafterme
A golang daemon to like / retweet anything a list of users post. Useful for important twitter users to you which you think all your followers should also know about.

[https://github.com/Jmainguy/repeatafterme](https://github.com/Jmainguy/repeatafterme "Repeatafterme")

## Relevy
A Golang daemon to collect information on a host, and place into mongodb.

[https://github.com/Jmainguy/relevy](https://github.com/Jmainguy/relevy "Relevy")

## Relevy-web
A golang daemon to display information about hosts from a mongodb server, to be used in conjunction with Relevy. Useful for configuration Management Database. It removes hosts if they havent checked in for a bit, so you can dynamically add / remove servers and the list stays updated.

[https://github.com/Jmainguy/relevy-web](https://github.com/Jmainguy/relevy-web "Relevy-Web")

## Gophermound
A golang daemon to collect and display information about a zookeeper cluster. Useful for providing information over http that a developer without shell access would be unable to get.

[https://github.com/Jmainguy/gophermound](https://github.com/Jmainguy/gophermound "GopherMound")

## Multicraft Scripts
I wrote a bunch of scripts to enable the multicraft hosting panel to host many more games than just vanilla minecraft.

[https://github.com/Standouthost/Multicraft](https://github.com/Standouthost/Multicraft "Multicraft Scripts")

## Bak
bak is a single file/directory backup utility I wrote in python. The idea came about after taking a training course where the recommended we run really long commands to essentially cp a file and append .bak to the end, bak provides this and more in a much shorter and easier to remember command.

[https://github.com/jmainguy/bak](https://github.com/jmainguy/bak "Bak")

## rpmbuild
docker_rpmbuild is a solution I came up with easily build rpm's across multiple operating systems using docker.

[https://github.com/jmainguy/docker_rpmbuild](https://github.com/jmainguy/docker_rpmbuild "Rpmbuild")

## statuscode
Memorizing what each http response code means is silly, use this instead. A sinatrarb website to return http status codes with explanations.

[https://github.com/Jmainguy/openshift-statuscode](https://github.com/Jmainguy/openshift-statuscode "Statuscode)

## ppq
python package query (ppq) is a clone of the rpm command for debian based systems with limited functionality that I wrote. This is useful if you prefer the rpm -qa and rpm -ql style output present on fedora based systems, but manage debian systems.

[https://github.com/jmainguy/ppq](https://github.com/jmainguy/ppq "Ppq")

## sinatrarb-virt-manager
sinatrarb-virt-manager is a sinatrarb implementation to easily restart libvirt vm's.

[https://github.com/Jmainguy/sinatrarb-virt-manager](https://github.com/Jmainguy/sinatrarb-virt-manager "sinatrarb-virt-manager")

## goblin
A limnoria (irc bot) plugin to roll multiple dice with modifiers (DnD style dice)

[https://github.com/Jmainguy/Goblin](https://github.com/Jmainguy/Goblin "Goblin")

## Tabs
A github repo of music I have tabulated for guitar

[https://github.com/Jmainguy/tabs](https://github.com/Jmainguy/tabs "Tabs")

## Powerchord
A python script to ease tabulating of power chords. (needs work)

[https://github.com/Jmainguy/powerchord](https://github.com/Jmainguy/powerchord "Powerchord")

## systemd-minecraft
Minecraft running as a systemd service, deployed by ansible. Wrote this for an interview at a really neat gaming company, sadly they had a slow hiring process and I missed out on joining them during my last free-agency period. Still a fun solution

[https://github.com/Jmainguy/systemd-minecraft](https://github.com/Jmainguy/systemd-minecraft "Systemd-Minecraft")

