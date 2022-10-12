#!/usr/bin/env bash

# Enable auto-home dir creation.
pushd /etc/pam.d/ >/dev/null
pam_auto_homedir='session     optional      pam_oddjob_mkhomedir.so umask=0077'
for f in password-auth system-auth; do
	cp -p ${f} ${f}.bak
	sed -i '/pam_oddjob_mkhomedir/d' ${f}
	echo "${pam_auto_homedir}" >> ${f}
done
popd >/dev/null

systemctl restart oddjobd