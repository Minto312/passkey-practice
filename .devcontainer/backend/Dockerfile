FROM golang:1.24-bookworm


ARG username=vscode
ARG useruid=1000
ARG usergid=${useruid}

RUN groupadd --gid ${usergid} ${username} && \
    useradd -s /bin/bash --uid ${useruid} --gid ${usergid} -m ${username} && \
    mkdir /etc/sudoers.d && \
    echo ${username} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${username} && \
    chmod 0440 /etc/sudoers.d/${username} && \
    chown -R ${username}:${username} /home/${username}

USER ${username}
