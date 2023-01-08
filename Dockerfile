FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:alpine
RUN apk add --update --no-cache vim git make musl-dev go~=1.17.10 curl openjdk11-jre-headless
RUN ["gcloud", "components", "install", "beta", "cloud-firestore-emulator", "--quiet"]
# RUN firestore emulator
# WAIT 10 seconds for the emulator to start

WORKDIR /app
COPY . .

#CMD ["gcloud", "beta", "emulators" ,"firestore","start","--quiet","--host-port","localhost:8020"]
CMD /bin/sh -c "./run-test.sh"
COPY ./pkg/setting/files/setting.develop.yaml /setting.yaml
ENV ENV_LOCATION /setting.yaml
ENV FIRESTORE_EMULATOR_HOST localhost:8020
