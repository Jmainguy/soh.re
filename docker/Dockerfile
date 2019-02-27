# Maintainer Jonathan Mainguy jon@soh.re
FROM centos:centos7
ADD gotty /bin/gotty
RUN yum install -y epel-release
RUN yum install -y openssh-clients bash-completion weechat
RUN useradd jmainguy
RUN mkdir /home/jmainguy/.ssh
RUN mkdir /home/jmainguy/.weechat
ADD known_hosts /home/jmainguy/.ssh/known_hosts
ADD id_rsa /home/jmainguy/.ssh/id_rsa
ADD .weechat /home/jmainguy/.weechat
ADD Jonathan_Mainguy_Profile/* /home/jmainguy/
EXPOSE 8080
ENV TERM=xterm
WORKDIR /home/jmainguy
ADD .gotty /home/jmainguy/.gotty
ADD motd /etc/motd
ADD bashrc /home/jmainguy/.bashrc
ADD config /home/jmainguy/.ssh/config
RUN chown -R jmainguy:jmainguy /home/jmainguy
RUN chmod 600 /home/jmainguy/.ssh/*
USER jmainguy
CMD ["/bin/gotty", "--once", "-w", "/bin/bash"]
