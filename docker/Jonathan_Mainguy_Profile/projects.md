# Projects
This is a incomplete list of projects and tools I have worked on, for a full list you should visit https://github.com/jmainguy

## This website

This website uses gotty (for the interactive shell), docker (to contain it, keep people from breaking things), iptables (to limit network), a custom router I wrote named soh-router (this intercepts websocket requests, then manages docker containers, then sends the socket to an available connection), sqlite3 (to store the available connections), systemd (to run the router as a service)

[https://github.com/Jmainguy/soh.re](https://github.com/Jmainguy/soh.re "This Website")

## Bible
A command line bible. It reads from an sqlite3 database. Supports multiple translations. I spend most my time in the command line, and figured having a bible available there would be useful.

This is installed on this system for you to use.

[https://github.com/Jmainguy/bible](https://github.com/Jmainguy/bible "Bible")

## Bak
bak is a single file/directory backup utility I wrote in python. The idea came about after taking a training course where the recommended we run really long commands to essentially cp a file and append .bak to the end, bak provides this and more in a much shorter and easier to remember command.

This is installed on this system for you to use.

[https://github.com/jmainguy/bak](https://github.com/jmainguy/bak "Bak")
