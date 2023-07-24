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

# Add automated stuff
RUN mkdir -p /usr/bin /usr/lib64/perl5/vendor_perl /usr/share/perl5/vendor_perl/Exporter /usr/lib64/perl5/vendor_perl/auto/Data/Dumper /usr/lib/locale/C.utf8 /usr/share/terminfo/x /usr/lib64/perl5/vendor_perl/Data /usr/lib64/python3.11/collections/__pycache__ /usr/share/locale /usr/lib/locale/C.utf8/LC_MESSAGES /usr/share/perl5 /usr/local/share/perl5/5.36/Term/Animation /usr/share/perl5/vendor_perl /usr/lib64/python3.11 /usr/lib/python3.11 /usr/lib64/python3.11/re/__pycache__ /usr/local/bin /bin /usr/lib64/perl5 /usr/share/perl5/Getopt /usr/lib64/perl5/vendor_perl/auto/Curses /usr/lib64/python3.11/encodings/__pycache__ /etc /lib64 /usr/local/share/perl5/5.36/Term /home/jmainguy/bin /home/jmainguy /usr/lib64/python3.11/__pycache__ /usr/lib64/python3.11/lib-dynload /usr/lib64/python3.11/site-packages/xattr/__pycache__ /usr/lib64/gconv /usr/share/perl5/warnings . /usr/lib64 /usr/lib64/python3.11/site-packages /usr/lib64/python3.11/site-packages/xattr
COPY --from=build /usr/local/share/perl5/5.36/Term/Animation.pm /usr/local/share/perl5/5.36/Term/Animation.pm
COPY --from=build /usr/lib64/perl5/Config.pm /usr/lib64/perl5/Config.pm
COPY --from=build /usr/share/perl5/XSLoader.pm /usr/share/perl5/XSLoader.pm
COPY --from=build /usr/share/perl5/bytes.pm /usr/share/perl5/bytes.pm
COPY --from=build /usr/lib64/python3.11/__pycache__/types.cpython-311.pyc /usr/lib64/python3.11/__pycache__/types.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/lib-dynload/zlib.cpython-311-x86_64-linux-gnu.so /usr/lib64/python3.11/lib-dynload/zlib.cpython-311-x86_64-linux-gnu.so
COPY --from=build /usr/lib64/python3.11/site-packages/xattr/__pycache__/__init__.cpython-311.pyc /usr/lib64/python3.11/site-packages/xattr/__pycache__/__init__.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/__pycache__/shutil.cpython-311.pyc /usr/lib64/python3.11/__pycache__/shutil.cpython-311.pyc
COPY --from=build /lib64/libtinfo.so.6 /lib64/libtinfo.so.6
COPY --from=build /usr/lib/locale/C.utf8/LC_MONETARY /usr/lib/locale/C.utf8/LC_MONETARY
COPY --from=build /usr/share/perl5/vars.pm /usr/share/perl5/vars.pm
COPY --from=build /lib64/libpanelw.so.6 /lib64/libpanelw.so.6
COPY --from=build /lib64/libmenuw.so.6 /lib64/libmenuw.so.6
COPY --from=build /usr/lib64/python3.11/encodings/__pycache__/aliases.cpython-311.pyc /usr/lib64/python3.11/encodings/__pycache__/aliases.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/__pycache__/warnings.cpython-311.pyc /usr/lib64/python3.11/__pycache__/warnings.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/lib-dynload/_bz2.cpython-311-x86_64-linux-gnu.so /usr/lib64/python3.11/lib-dynload/_bz2.cpython-311-x86_64-linux-gnu.so
COPY --from=build /home/jmainguy/bin/bak /home/jmainguy/bin/bak
COPY --from=build /usr/lib64/python3.11/encodings/__pycache__/utf_8.cpython-311.pyc /usr/lib64/python3.11/encodings/__pycache__/utf_8.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/__pycache__/functools.cpython-311.pyc /usr/lib64/python3.11/__pycache__/functools.cpython-311.pyc
COPY --from=build /lib64/libncursesw.so.6 /lib64/libncursesw.so.6
COPY --from=build /usr/lib/locale/C.utf8/LC_NUMERIC /usr/lib/locale/C.utf8/LC_NUMERIC
COPY --from=build /usr/lib64/python3.11/site-packages/xattr/__pycache__/lib.cpython-311.pyc /usr/lib64/python3.11/site-packages/xattr/__pycache__/lib.cpython-311.pyc
COPY --from=build /etc/ld.so.cache /etc/ld.so.cache
COPY --from=build /usr/lib/locale/C.utf8/LC_CTYPE /usr/lib/locale/C.utf8/LC_CTYPE
COPY --from=build /lib64/libcrypt.so.2 /lib64/libcrypt.so.2
COPY --from=build /usr/lib64/perl5/vendor_perl/auto/Curses/Curses.so /usr/lib64/perl5/vendor_perl/auto/Curses/Curses.so
COPY --from=build /usr/lib64/python3.11/encodings/__pycache__/__init__.cpython-311.pyc /usr/lib64/python3.11/encodings/__pycache__/__init__.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/re/__pycache__/_casefix.cpython-311.pyc /usr/lib64/python3.11/re/__pycache__/_casefix.cpython-311.pyc
COPY --from=build /lib64/libc.so.6 /lib64/libc.so.6
COPY --from=build /usr/bin/asciiquarium /home/jmainguy/bin/asciiquarium
COPY --from=build /usr/lib64/python3.11/site-packages/xattr/__pycache__/compat.cpython-311.pyc /usr/lib64/python3.11/site-packages/xattr/__pycache__/compat.cpython-311.pyc
COPY --from=build /usr/lib64/perl5/vendor_perl/Curses.pm /usr/lib64/perl5/vendor_perl/Curses.pm
COPY --from=build /usr/lib64/perl5/vendor_perl/Data/Dumper.pm /usr/lib64/perl5/vendor_perl/Data/Dumper.pm
COPY --from=build /usr/lib64/python3.11/__pycache__/copyreg.cpython-311.pyc /usr/lib64/python3.11/__pycache__/copyreg.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/site-packages/xattr /usr/lib64/python3.11/site-packages/xattr
COPY --from=build /lib64/libbz2.so.1 /lib64/libbz2.so.1
COPY --from=build /usr/lib/locale/C.utf8/LC_TELEPHONE /usr/lib/locale/C.utf8/LC_TELEPHONE
COPY --from=build /usr/lib/locale/C.utf8/LC_NAME /usr/lib/locale/C.utf8/LC_NAME
COPY --from=build /usr/lib/locale/C.utf8/LC_PAPER /usr/lib/locale/C.utf8/LC_PAPER
COPY --from=build /usr/share/perl5/vendor_perl/Exporter.pm /usr/share/perl5/vendor_perl/Exporter.pm
COPY --from=build /usr/share/perl5/Getopt/Std.pm /usr/share/perl5/Getopt/Std.pm
COPY --from=build /usr/lib64/python3.11/site-packages /usr/lib64/python3.11/site-packages
COPY --from=build /usr/lib64/python3.11/__pycache__/argparse.cpython-311.pyc /usr/lib64/python3.11/__pycache__/argparse.cpython-311.pyc
COPY --from=build /lib64/libm.so.6 /lib64/libm.so.6
COPY --from=build /usr/lib64/perl5/DynaLoader.pm /usr/lib64/perl5/DynaLoader.pm
COPY --from=build /usr/lib64/perl5/vendor_perl/auto/Data/Dumper/Dumper.so /usr/lib64/perl5/vendor_perl/auto/Data/Dumper/Dumper.so
COPY --from=build /usr/bin/python3 /usr/bin/python3
COPY --from=build /usr/lib64/python3.11/site-packages/xattr/_lib.abi3.so /usr/lib64/python3.11/site-packages/xattr/_lib.abi3.so
COPY --from=build /usr/lib/locale/C.utf8/LC_COLLATE /usr/lib/locale/C.utf8/LC_COLLATE
COPY --from=build /usr/share/perl5/warnings.pm /usr/share/perl5/warnings.pm
COPY --from=build /usr/lib64/python3.11 /usr/lib64/python3.11
COPY --from=build /usr/lib64/python3.11/re/__pycache__/__init__.cpython-311.pyc /usr/lib64/python3.11/re/__pycache__/__init__.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/__pycache__/operator.cpython-311.pyc /usr/lib64/python3.11/__pycache__/operator.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/re/__pycache__/_compiler.cpython-311.pyc /usr/lib64/python3.11/re/__pycache__/_compiler.cpython-311.pyc
COPY --from=build /lib64/liblzma.so.5 /lib64/liblzma.so.5
COPY --from=build /lib64/libffi.so.8 /lib64/libffi.so.8
COPY --from=build /usr/lib64/python3.11/__pycache__/locale.cpython-311.pyc /usr/lib64/python3.11/__pycache__/locale.cpython-311.pyc
COPY --from=build /usr/lib/python3.11/site-packages /usr/lib/python3.11/site-packages
COPY --from=build /usr/lib64/python3.11/__pycache__/reprlib.cpython-311.pyc /usr/lib64/python3.11/__pycache__/reprlib.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/re/__pycache__/_parser.cpython-311.pyc /usr/lib64/python3.11/re/__pycache__/_parser.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/__pycache__/fnmatch.cpython-311.pyc /usr/lib64/python3.11/__pycache__/fnmatch.cpython-311.pyc
COPY --from=build /lib64/libz.so.1 /lib64/libz.so.1
COPY --from=build /usr/lib64/python3.11/__pycache__/_compression.cpython-311.pyc /usr/lib64/python3.11/__pycache__/_compression.cpython-311.pyc
COPY --from=build /usr/share/locale/locale.alias /usr/share/locale/locale.alias
COPY --from=build /usr/lib/locale/C.utf8/LC_IDENTIFICATION /usr/lib/locale/C.utf8/LC_IDENTIFICATION
COPY --from=build /usr/share/perl5/vendor_perl/constant.pm /usr/share/perl5/vendor_perl/constant.pm
COPY --from=build /lib64/libpython3.11.so.1.0 /lib64/libpython3.11.so.1.0
COPY --from=build /usr/lib64/python3.11/__pycache__/lzma.cpython-311.pyc /usr/lib64/python3.11/__pycache__/lzma.cpython-311.pyc
COPY --from=build /usr/lib/locale/C.utf8/LC_MEASUREMENT /usr/lib/locale/C.utf8/LC_MEASUREMENT
COPY --from=build /usr/lib/locale/C.utf8/LC_MESSAGES/SYS_LC_MESSAGES /usr/lib/locale/C.utf8/LC_MESSAGES/SYS_LC_MESSAGES
COPY --from=build /usr/bin/perl /usr/bin/perl
COPY --from=build /usr/share/perl5/vendor_perl/Exporter/Heavy.pm /usr/share/perl5/vendor_perl/Exporter/Heavy.pm
COPY --from=build /usr/lib64/python3.11/__pycache__/enum.cpython-311.pyc /usr/lib64/python3.11/__pycache__/enum.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/site-packages/_cffi_backend.cpython-311-x86_64-linux-gnu.so /usr/lib64/python3.11/site-packages/_cffi_backend.cpython-311-x86_64-linux-gnu.so
COPY --from=build /usr/lib64/python3.11/collections/__pycache__/__init__.cpython-311.pyc /usr/lib64/python3.11/collections/__pycache__/__init__.cpython-311.pyc
COPY --from=build /usr/lib64/gconv/gconv-modules.cache /usr/lib64/gconv/gconv-modules.cache
COPY --from=build /usr/lib/locale/C.utf8/LC_TIME /usr/lib/locale/C.utf8/LC_TIME
COPY --from=build /bin/asciiquarium /bin/asciiquarium
COPY --from=build /usr/share/perl5/vendor_perl/Carp.pm /usr/share/perl5/vendor_perl/Carp.pm
COPY --from=build /etc/localtime /etc/localtime
COPY --from=build /usr/lib64/python3.11/encodings /usr/lib64/python3.11/encodings
COPY --from=build /home/jmainguy/bin /home/jmainguy/bin
COPY --from=build /usr/lib64/python3.11/re /usr/lib64/python3.11/re
COPY --from=build /usr/lib64/python3.11/re/__pycache__/_constants.cpython-311.pyc /usr/lib64/python3.11/re/__pycache__/_constants.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/__pycache__/bz2.cpython-311.pyc /usr/lib64/python3.11/__pycache__/bz2.cpython-311.pyc
COPY --from=build /usr/local/bin/cmatrix /home/jmainguy/bin/cmatrix
COPY --from=build /lib64/libformw.so.6 /lib64/libformw.so.6
COPY --from=build /usr/lib/locale/C.utf8/LC_MESSAGES /usr/lib/locale/C.utf8/LC_MESSAGES
COPY --from=build /usr/share/perl5/warnings/register.pm /usr/share/perl5/warnings/register.pm
COPY --from=build /usr/local/share/perl5/5.36/Term/Animation/Entity.pm /usr/local/share/perl5/5.36/Term/Animation/Entity.pm
COPY --from=build /usr/lib64/python3.11/__pycache__/gettext.cpython-311.pyc /usr/lib64/python3.11/__pycache__/gettext.cpython-311.pyc
COPY --from=build /usr/lib64/python3.11/lib-dynload/_lzma.cpython-311-x86_64-linux-gnu.so /usr/lib64/python3.11/lib-dynload/_lzma.cpython-311-x86_64-linux-gnu.so
COPY --from=build /usr/lib/locale/C.utf8/LC_ADDRESS /usr/lib/locale/C.utf8/LC_ADDRESS
COPY --from=build /usr/share/terminfo/x/xterm /usr/share/terminfo/x/xterm
COPY --from=build /lib64/libperl.so.5.36 /lib64/libperl.so.5.36
COPY --from=build /usr/share/perl5/strict.pm /usr/share/perl5/strict.pm
COPY --from=build /usr/share/perl5/overloading.pm /usr/share/perl5/overloading.pm
COPY --from=build /usr/lib64/python3.11/lib-dynload /usr/lib64/python3.11/lib-dynload
COPY --from=build /usr/lib64/python3.11/__pycache__/keyword.cpython-311.pyc /usr/lib64/python3.11/__pycache__/keyword.cpython-311.pyc

FROM scratch
COPY --from=middle / /
ADD gotty /bin/gotty
USER jmainguy
WORKDIR /home/jmainguy
EXPOSE 8080
ENV TERM=xterm
CMD ["/bin/gotty", "--once", "-w", "/bin/bash"]
