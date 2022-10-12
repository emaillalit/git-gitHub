#!/usr/bin/env bash

sed -i '/^sudoers:/d' /etc/nsswitch.conf
cat << E0F >> /etc/nsswitch.conf

# Arena SSSD update by firstboot.
sudoers:    files sss
E0F