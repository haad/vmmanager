# vmware fusion player/pro Manual for Mac OS X

## Intro 

We would like to use vmware fusion to run a vm where we can provision docker + setup build systems and sync our sources. 


## vmware Fusion 

vmware fusion is a vm manager for mac os utilizing native virtualization.framework with support for both amd64 and arm64 based macs.


### Install vmware fusion and other requirements

``` 
brew install --cask vmware-fusion
brew install qemu
```

Visit [vmware fusion](https://customerconnect.vmware.com/en/evalcenter?p=fusion-player-personal-13) website to get a [free license](https://www.vmware.com/go/get-fusionplayer-key) for non commercial fusion usage. If you want to use it for business pls make sure to buy a license. 

```https://www.vmware.com/go/get-fusionplayer-key```

### Get ubuntu cloud image

Ubuntu cloud are hosted images daily built by ubuntu for their releases. They can be found [here](https://cloud-images.ubuntu.com/). LAtest LTS release is [jammy](https://cloud-images.ubuntu.com/jammy/current/). These images are for arm64 architecture built as qcow2 disks for qemu. To use it in vmware fusion we have to convert it to vmdk this is where we can use `qemu-img`tool.

```
wget https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-arm64.img
qemu-img convert -p -f qcow2 -O vmdk ubuntu-22.04-server-cloudimg-arm64.img ubuntu-22.04-server-cloudimg-arm64.vmdk
``` 

## Prepare VM for usage

Converted VM images can be added to newly created VM templates and will boot instantly. They expect to run on a cloud and that's why we have to provide them with some initial cloud-init information to provision our user. Cloud-init tools has an excellent manual about that [here](https://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html). 

```
mkdir /tmp/seed
echo -e "instance-id: iid-local01\nlocal-hostname: cloudimg" > /tmp/seed/meta-data
echo -e "#cloud-config\nhostname: cloudimg\npassword: passw0rd\nchpasswd: { expire: False }\nssh_pwauth: True\n" > /tmp/seed/user-data

hdiutil makehybrid -o /tmp/seed.iso -hfs -joliet -iso -default-volume-name cidata /tmp/seed/
```

## Initial boot

Before we boot our vm for a first time we need to attach created `seed.iso` file to it as a CD. After first boot with CD iso image attached our default user will have a new password set `passw0rd`. 


### First login 

After initial boot try to login to VM with your default user e.g. `ubuntu` and new password. If this works you can figure out VM ip address from a gui window or you can use following command. You can find vmrun documentation [here](https://docs.vmware.com/en/VMware-Fusion/13/com.vmware.fusion.using.doc/GUID-24F54E24-EFB0-4E94-8A07-2AD791F0E497.html).

```
vmrun list

# Command below will get you an ip address of a first vm you run.
vmrun getGuestIPAddress $(vmrun list | grep vmx | head -n 1)
```

## Usage and administration

### Configuring static ip address

To properly configure static ip address to our vmware fusion vm we first need to figure out actual ip/dhcpd setup. We can do it with following commands.

```
# Use command to list all vmnet-dhcpd processes look for different instances and check for a `vmnet8` interface.
ps ax | grep vmnet

```

For my instalation path to configuration file looks like this: `/Library/Preferences/VMware Fusion/vmnet8/dhcpd.conf`.  

```
❯ cat '/Library/Preferences/VMware Fusion/vmnet8/dhcpd.conf' | grep range
range 192.168.240.128 192.168.240.254;
```

DHCPd configuration file above is configured to only server ips from range `192.168.240.128 192.168.240.254`. If we want to statically assign ip it should be lower than `> 192.168.240.5 and <= 192.168.240.127.

Use `/etc/systemd/network/interface_name.network` file to configure it.

```
[Match]
# You can also use wildcards. Maybe you want enable dhcp
# an all eth* NICs
Name=eth0
[Network]
#DHCP=v4
# static IP
# 192.168.100.2 netmask 255.255.255.0
Address=192.168.100.2/24
Gateway=192.168.100.1
DNS=192.168.100.1

```

### Using vmrun to manipulate vm state

```
vmrun start /Devel/vms/primary.vmwarevm/primary.vmx nogui
vmrun stip /Devel/vms/primary.vmwarevm/primary.vmx

```

### Setup Mutagen to automatically sync your files to new machine

```
Host linux
    Hostname            192.168.168.5
    User                haad
    Ciphers             aes128-ctr,aes256-ctr,chacha20-poly1305@openssh.com
    ForwardAgent        yes
    TCPKeepAlive        yes
    ServerAliveInterval 10
    ServerAliveCountMax 300
    Compression         yes
    IdentityFile        ~/.ssh/id_key
```


```
mutagen sync create --name=php ~/Devel/php linux:~/Devel/php
mutagen sync create --name=pixel ~/Devel/chilli/pixel linux:~/chilli/pixel
mutagen sync create --name=go ~/Devel/go/src/github.com/haad linux:~/Devel/go/src/github.com/haad
```

### Using docker from mac machine

Set DOCKER_HOST in your environment and use you use docker command directly from mac.

```
DOCKER_HOST=ssh://linux
```