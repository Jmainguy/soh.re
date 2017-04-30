# Maintainer Jonathan Mainguy jon@soh.re
FROM centos:centos7
ADD gotty /bin/gotty
RUN yum install -y openssh-clients bash-completion
RUN useradd joshua
USER joshua
RUN mkdir /home/joshua/.ssh
ADD known_hosts /home/joshua/.ssh/known_hosts
ADD id_rsa /home/joshua/.ssh/id_rsa
EXPOSE 8080
ENV TERM=xterm
WORKDIR /home/joshua
ADD .gotty /home/joshua/.gotty
CMD ["/bin/gotty", "--timeout", "10", "--once", "-w", "/bin/bash"]
