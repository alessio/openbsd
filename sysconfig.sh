touch /etc/sysctl.conf
cat >>/etc/sysctl.conf <<EOF
# shared memory limits (chrome needs a ton)
kern.shminfo.shmall=3145728
kern.shminfo.shmmax=2147483647
kern.shminfo.shmmni=1024

# semaphores
kern.shminfo.shmseg=1024
kern.seminfo.semmns=4096
kern.seminfo.semmni=1024

kern.maxproc=32768
kern.maxfiles=65535
kern.bufcachepercent=90
kern.maxvnodes=262144
kern.somaxconn=2048
EOF

touch /etc/login.conf
cat >>/etc/login.conf <<EOF

staff:\
  :datasize-cur=infinity:\
  :datasize-max=infinity:\
  :maxproc-cur=512:\
  :maxproc-max=1024:\
  :openfiles-cur=102400:\
  :openfiles-max=102400:\
  :stacksize-cur=32M:\
  :ignorenologin:\
  :requirehome@:\
  :tc=default:
EOF

touch /etc/doas.conf
echo 'permit nopass keepenv :wheel' >> /etc/doas.conf

sed -i.bak 's/rw/rw,softdep,noatime/g' /etc/fstab
diff -u /etc/fstab.bak /etc/fstab
