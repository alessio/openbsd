# install packages

pkg_add lxqt lxqt-extras consolekit2

# configure services

rcctl enable messagebus
rcctl start messagebus
