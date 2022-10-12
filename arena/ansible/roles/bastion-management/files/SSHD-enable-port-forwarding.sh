#!/usr/bin/env bash

# Update /etc/ssh/sshd_config to set 'GatewayPorts yes'
sed -i '/^GatewayPorts /d' /etc/ssh/sshd_config
cat << E0F >> /etc/ssh/sshd_config
# Firstboot: Allow port forwarding to bind to non-localhost address
GatewayPorts yes
E0F

systemctl restart sshd
