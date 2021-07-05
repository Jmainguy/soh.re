# Maintainer Jonathan Mainguy jon@soh.re
FROM centos:centos8
ENV version 0704202101
ADD gotty /bin/gotty
RUN dnf install -y epel-release
RUN dnf install -y openssh-clients bash-completion weechat sudo git gcc python39 perl-CPAN perl-Curses ncurses-devel cmake gcc-c++
RUn dnf update -y
RUN alternatives --set python /usr/bin/python3
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
ENV PATH=/home/linuxbrew/.linuxbrew/bin:$PATH
# Bible
RUN brew install jmainguy/tap/bible
# certificatedownloader
RUN brew install jmainguy/tap/certificatedownloader
# Bak
RUN curl https://raw.githubusercontent.com/Jmainguy/bak/main/bak --output /bin/bak
# asciiquarium
RUN curl https://raw.githubusercontent.com/Jmainguy/asciiquarium/main/asciiquarium --output /bin/asciiquarium
RUN echo "y" | cpan Term::Animation || echo "moving on"
RUN cp -R /root/.cpan/build/Term-Animation-2.6-0/lib/Term/Animation* /usr/share/perl5/vendor_perl/Term/
# cmatrix
RUN git clone https://github.com/abishekvashok/cmatrix
RUN mkdir -p cmatrix/build && cd cmatrix/build && cmake .. && make && make install
# indent
RUN brew install jmainguy/tap/indent
# czechdns
RUN brew install jmainguy/tap/czechdns
# Cleanup homebrew
RUN cp /home/linuxbrew/.linuxbrew/bin/* /bin/
RUN chmod +x /bin/*
RUN rm -rf /home/linuxbrew
# Add rest of things
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
