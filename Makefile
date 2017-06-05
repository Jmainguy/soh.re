# This expects a few requirements
# one, that https://github.com/Jmainguy/docker_rpmbuild is cloned into ~/Github/docker_rpmbuild
# two, that docker is installed and running
# three, that ~/Github/docker_rpmbuild/dockerbuild/build.sh centos7 has been run
rpm:
	@tar -czvf ~/Github/docker_rpmbuild/rpmbuild/SOURCES/soh-router.tar.gz ../soh.re
	@cp soh-router.spec ~/Github/docker_rpmbuild/rpmbuild/SPECS/soh-router.spec
	@cd ~/Github/docker_rpmbuild/; ./run.sh centos7 soh-router
	@ls -ltrh ~/Github/docker_rpmbuild/rpmbuild/RPMS/x86_64/soh-router*
