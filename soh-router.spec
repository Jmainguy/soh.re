%if 0%{?rhel} == 7
  %define dist .el7
%endif
%define _unpackaged_files_terminate_build 0
Name:	 soh-router
Version: 0.5
Release: 1%{?dist}
Summary: A golang daemon to run soh.re router

License: GPLv2
URL: https://github.com/jmainguy/soh.re
Source0: soh-router.tar.gz

%description
A golang daemon to run soh.re router

%prep
%setup -q -n soh.re
%build
export GOPATH=/usr/src/go
cd soh-router
go build
%install
mkdir -p $RPM_BUILD_ROOT/usr/sbin
mkdir -p $RPM_BUILD_ROOT/usr/lib/systemd/system
mkdir -p $RPM_BUILD_ROOT/opt/soh-router
mkdir -p $RPM_BUILD_ROOT/etc/soh-router
install -m 0644 $RPM_BUILD_DIR/soh.re/service/soh-router.service %{buildroot}/usr/lib/systemd/system
install -m 0755 $RPM_BUILD_DIR/soh.re/soh-router/soh-router %{buildroot}/usr/sbin
install -m 0644 $RPM_BUILD_DIR/soh.re/config.yaml %{buildroot}/etc/soh-router/

%files
/usr/sbin/soh-router
/usr/lib/systemd/system/soh-router.service
%dir /opt/soh-router
%dir /etc/soh-router
%config(noreplace) /etc/soh-router/config.yaml
%doc

%pre
getent group soh-router >/dev/null || groupadd -r soh-router
getent passwd soh-router >/dev/null || \
    useradd -r -g soh-router -d /opt/soh-router -s /sbin/nologin \
    -c "User to run soh-router service" soh-router && \
    groupadd docker && \
    usermod -aG docker soh-router
exit 0
%post
chown -R soh-router:soh-router /opt/soh-router
if [ -f /usr/bin/systemctl ]; then
  systemctl daemon-reload && systemctl restart docker
fi
semanage port -a -t http_port_t -p tcp 8085

%changelog
