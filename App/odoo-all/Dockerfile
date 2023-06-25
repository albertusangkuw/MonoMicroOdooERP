FROM debian:bullseye-slim

# Generate locale C.UTF-8 for postgres and general locale data
ENV LANG C.UTF-8

# Install some deps, lessc and less-plugin-clean-css, and wkhtmltopdf
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        software-properties-common \
        ca-certificates \
        curl \
        dirmngr \
        fonts-noto-cjk \
        gnupg \
        libssl-dev \
        libldap2-dev\
        libpq-dev\
        libsasl2-dev\ 
        node-less \
        npm \
        nano \
        python3\
        python3-num2words \
        python3-pdfminer \
        python3-pip \
        python3-phonenumbers \
        python3-pyldap \
        python3-qrcode \
        python3-renderpm \
        python3-setuptools \
        python3-slugify \
        python3-vobject \
        python3-watchdog \
        python3-xlrd \
        python3-xlwt \
        build-essential \
        ldap-utils tox \
        lcov valgrind\
        xz-utils \
        python3-dev\
    && curl -o wkhtmltox.deb -sSL https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_arm64.deb \
    && apt-get install -y --no-install-recommends ./wkhtmltox.deb \
    && rm -rf /var/lib/apt/lists/* wkhtmltox.deb

# Install rtlcss (on Debian buster)
RUN npm install -g rtlcss

RUN python3 -m pip install --upgrade pip
RUN python3 -m pip install wheel
RUN python3 -m pip install pyjwt
RUN pip install pip setuptools wheel Cython==3.0.0a10
RUN pip install gevent==20.9.0 --no-build-isolation
RUN python3 -m pip install --upgrade --force pip

COPY ./requirements.txt ./requirements.txt

# Install the Python dependencies
RUN pip3 install -r requirements.txt


RUN mkdir /home/odoo
RUN mkdir /home/odoo/shared_data

COPY . /home/odoo

WORKDIR /home/odoo

# Expose the default Odoo port
EXPOSE 8069

# Set the default command to start Odoo
CMD ["/bin/bash", "-c", "while true; do echo Service Monolith...; sleep 1; done"]