FROM mcr.microsoft.com/devcontainers/go:1-1.23-bullseye

# RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash 
# RUN apt-get update 
# RUN apt-get install migrate

# RUN apt-get update && apt-get install netcat redis-tools -y
# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# RUN go install github.com/air-verse/air@latest
# RUN go install github.com/rakyll/hey@latest
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
