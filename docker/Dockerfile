# Maintainer Jonathan Mainguy jon@soh.re
FROM centos:centos7
ADD gotty /bin/gotty
RUN yum install -y epel-release
RUN yum install openssh-clients bash-completion weechat
RUN useradd joshua
USER joshua
RUN mkdir /home/joshua/.ssh
RUN mkdir /home/joshua/.weechat
ADD known_hosts /home/joshua/.ssh/known_hosts
ADD id_rsa /home/joshua/.ssh/id_rsa
ADD .weechat /home/joshua/.weechat
ADD Jonathan_Mainguy_Profile/* /home/joshua/
EXPOSE 8080
ENV TERM=xterm
WORKDIR /home/joshua
ADD .gotty /home/joshua/.gotty
CMD ["/bin/gotty", "--once", "-w", "/bin/bash"]
