FROM registry.fedoraproject.org/fedora:45 AS ttyd-builder

# First upstream revision containing the mobile viewport fix from ttyd PR #1564.
ARG TTYD_REF=2922cb89f518bae4d0fcf4d757a7419638fc71fc

RUN dnf install -y --setopt=install_weak_deps=False \
        cmake \
        curl \
        gcc \
        json-c-devel \
        libuv-devel \
        libwebsockets-devel \
        openssl-devel \
        tar \
        zlib-devel \
    && dnf clean all \
    && mkdir -p /src/ttyd \
    && curl --fail --location --proto '=https' --tlsv1.2 \
        "https://github.com/tsl0922/ttyd/archive/${TTYD_REF}.tar.gz" \
        | tar --extract --gzip --directory /src/ttyd --strip-components=1 \
    && cmake -S /src/ttyd -B /src/ttyd/build \
        -DCMAKE_BUILD_TYPE=Release \
        -DCMAKE_INSTALL_PREFIX=/usr/local \
    && cmake --build /src/ttyd/build --parallel \
    && strip /src/ttyd/build/ttyd

FROM registry.fedoraproject.org/fedora-minimal:46 AS cli-assets

ARG TARGETARCH
ARG ASCIQUARIUM_VERSION=2.0.0
ARG BIBLE_REF=4d107ce3516a1b93028bfc1229993a823bac7fea
ARG BIBLE_VERSION=1.0.3
ARG WEATHERSPECT_VERSION=2.0.0

RUN microdnf install -y --setopt=install_weak_deps=0 \
        ca-certificates \
        cpio \
        curl-minimal \
        rpm \
    && microdnf clean all \
    && mkdir -p /out/asciiquarium /out/bible /out/bible-db /out/weatherspect \
    && case "${TARGETARCH}" in \
        amd64) \
            rpm_arch=x86_64; \
            asciiquarium_sha=1349ff0daa23dd55d75e98c28758d2a9bee1b6c0fd3635e24f2313344cb926be; \
            bible_sha=407ad679717ecd883db29f6ae8226a90f8ddbcd5a9959276725259b5305ccfe4; \
            weatherspect_sha=4a56e1d439d092a3c6075bfb84630ba8cdeb2db35fb1b37bf41aefde62dd3f6a ;; \
        arm64) \
            rpm_arch=aarch64; \
            asciiquarium_sha=eaa56c58acb34450a40acf6d4f54d321df40fb194b51801264bd0be37ffe5a88; \
            bible_sha=28a898d24dfc5396841325b73bb645121f88b7323a57511ae95cddd570bf2ab1; \
            weatherspect_sha=c08d17e3540155bdf3f7001ca8692abc5fe4dd10f8dbedd87e077b83fdba608e ;; \
        *) echo "unsupported architecture: ${TARGETARCH}" >&2; exit 1 ;; \
    esac \
    && curl --fail --location --proto '=https' --tlsv1.2 \
        "https://github.com/Jmainguy/asciiquarium/releases/download/v${ASCIQUARIUM_VERSION}/asciiquarium-${ASCIQUARIUM_VERSION}-1.${rpm_arch}.rpm" \
        --output /tmp/asciiquarium.rpm \
    && echo "${asciiquarium_sha}  /tmp/asciiquarium.rpm" | sha256sum --check --strict \
    && cd /out/asciiquarium \
    && rpm2cpio /tmp/asciiquarium.rpm | cpio --extract --make-directories --quiet \
    && curl --fail --location --proto '=https' --tlsv1.2 \
        "https://github.com/Jmainguy/bible/releases/download/v${BIBLE_VERSION}/bible-${BIBLE_VERSION}-1.${TARGETARCH}.rpm" \
        --output /tmp/bible.rpm \
    && echo "${bible_sha}  /tmp/bible.rpm" | sha256sum --check --strict \
    && cd /out/bible \
    && rpm2cpio /tmp/bible.rpm | cpio --extract --make-directories --quiet \
    && curl --fail --location --proto '=https' --tlsv1.2 \
        "https://raw.githubusercontent.com/Jmainguy/bible/${BIBLE_REF}/database/bible.db" \
        --output /out/bible-db/bible.db \
    && echo "4846347239cf5e324690c3fbdb02768646496ddf99b9e9a094d4cabaf3c4fb96  /out/bible-db/bible.db" \
        | sha256sum --check --strict \
    && curl --fail --location --proto '=https' --tlsv1.2 \
        "https://github.com/Jmainguy/weatherspect/releases/download/v${WEATHERSPECT_VERSION}/weatherspect-${WEATHERSPECT_VERSION}-1.${rpm_arch}.rpm" \
        --output /tmp/weatherspect.rpm \
    && echo "${weatherspect_sha}  /tmp/weatherspect.rpm" | sha256sum --check --strict \
    && cd /out/weatherspect \
    && rpm2cpio /tmp/weatherspect.rpm | cpio --extract --make-directories --quiet

FROM registry.fedoraproject.org/fedora-minimal:46

LABEL org.opencontainers.image.authors="Jonathan Seth Mainguy <jon@soh.re>" \
      org.opencontainers.image.description="Interactive terminal for soh.re" \
      org.opencontainers.image.licenses="GPL-2.0" \
      org.opencontainers.image.source="https://github.com/Jmainguy/soh.re" \
      org.opencontainers.image.title="soh.re" \
      org.opencontainers.image.url="https://github.com/Jmainguy/soh.re" \
      org.opencontainers.image.vendor="Jmainguy"

# Keep the interactive shell useful while omitting compilers, Python, Perl,
# and other build-only dependencies from the runtime image.
RUN microdnf upgrade -y \
    && microdnf install -y --setopt=install_weak_deps=0 \
        bash \
        bash-completion \
        ca-certificates \
        cmatrix \
        coreutils \
        findutils \
        git-core \
        grep \
        gzip \
        iproute \
        json-c \
        less \
        libuv \
        libwebsockets \
        openssl-libs \
        procps-ng \
        sed \
        shadow-utils \
        tar \
        which \
    && microdnf clean all \
    && rm -rf /var/cache/dnf /var/log/dnf* /tmp/*

COPY --from=ttyd-builder /src/ttyd/build/ttyd /usr/local/bin/ttyd
COPY --from=cli-assets /out/asciiquarium/usr/bin/asciiquarium /usr/local/bin/asciiquarium
COPY --from=cli-assets /out/bible/usr/bin/bible /usr/local/bin/bible
COPY --from=cli-assets /out/bible-db/bible.db /home/jmainguy/.bible/bible.db
COPY --from=cli-assets /out/weatherspect/usr/bin/weatherspect /usr/local/bin/weatherspect

COPY Jonathan_Mainguy_Profile/ /home/jmainguy/
COPY bashrc /home/jmainguy/.bashrc
COPY motd /etc/motd

RUN useradd --create-home --home-dir /home/jmainguy --shell /bin/bash jmainguy \
    && chown -R jmainguy:jmainguy /home/jmainguy

USER jmainguy
WORKDIR /home/jmainguy
EXPOSE 8080
ENV TERM=xterm
ENV WEATHERSPECT_LOCATION="Raleigh, NC"
CMD ["/usr/local/bin/ttyd", "--once", "--writable", "-p", "8080", "/bin/bash"]
