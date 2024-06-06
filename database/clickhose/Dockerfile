FROM python:3.9-slim

WORKDIR /app

COPY --chown=65534:65534 requirements.txt .
RUN chown -R 65534:65534 /app /var/log

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      build-essential \
      gcc \
      libssl-dev \
      libffi-dev \
      iputils-ping && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get clean

#RUN pip install --no-cache-dir -r requirements.txt

COPY --chown=65534:65534 . .

USER nobody
