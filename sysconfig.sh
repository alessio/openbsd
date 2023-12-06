cp sysctl.conf /etc
cp login.conf.d/* /etc/login.comf.d/

touch /etc/doas.conf
echo 'permit nopass keepenv :wheel' >> /etc/doas.conf

sed -i.bak 's/rw/rw,softdep,noatime/g' /etc/fstab
diff -u /etc/fstab.bak /etc/fstab
