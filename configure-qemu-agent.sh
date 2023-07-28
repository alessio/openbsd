flags="-t /var/run/qemu-ga -f /var/run/qemu-ga/qemu-ga.pid -m unix-listen"
rcctl set qemu_ga flags "$flags"
