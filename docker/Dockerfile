# Maintainer Jonathan Mainguy jon@soh.re
FROM fedora as build
RUN dnf install -y bash-completion git gcc perl-CPAN perl-Curses ncurses-devel cmake gcc-c++
# asciiquarium
RUN curl https://raw.githubusercontent.com/Jmainguy/asciiquarium/main/asciiquarium --output /bin/asciiquarium
RUN echo "y" | cpan Term::Animation || echo "moving on"
RUN cp -R /root/.cpan/build/Term-Animation-2.6-0/lib/Term/Animation* /usr/share/perl5/vendor_perl/Term/
RUN chmod +x /usr/bin/asciiquarium
# cmatrix
RUN git clone https://github.com/abishekvashok/cmatrix
RUN mkdir -p cmatrix/build && cd cmatrix/build && cmake .. && make && make install
# Add rest of things
ADD Jonathan_Mainguy_Profile/* /home/jmainguy/
RUN mkdir -p /home/jmainguy/bin
ADD .gotty /home/jmainguy/.gotty
ADD bashrc /home/jmainguy/.bashrc
# Bible
RUN curl -L https://github.com/Jmainguy/bible/releases/download/v0.3.1/bible_Linux_x86_64.tar.gz --output /tmp/bible.tar.gz
RUN tar zxvf /tmp/bible.tar.gz
RUN cp bible /home/jmainguy/bin/
RUN mkdir -p /home/jmainguy/.bible/
RUN curl -L https://github.com/Jmainguy/bible/raw/main/database/bible.db --output /home/jmainguy/.bible/bible.db
# Bak
RUN curl https://raw.githubusercontent.com/Jmainguy/bak/main/bak --output /home/jmainguy/bin/bak
RUN chmod +x /home/jmainguy/bin/*
RUN dnf install -y python3-xattr

FROM fedora as middle
ADD motd /etc/motd
COPY --from=build /home/jmainguy /home/jmainguy
RUN useradd jmainguy
RUN chown -R jmainguy:jmainguy /home/jmainguy

ADD gotty /bin/gotty
USER jmainguy
WORKDIR /home/jmainguy
EXPOSE 8080
ENV TERM=xterm
CMD ["/bin/gotty", "--once", "-w", "/bin/bash"]
