> PRO: Multi-Subcommand Provisioning Tool

> [!TIP]
> `digitalocean` is the only **platform** implemented currently.

```sh
Usage:
  $0 [platform] [options]

Subcommands:
  digitalocean  Manage DigitalOcean droplets
    create      Create a new droplet
    list        List existing droplets
    ssh         SSH into a selected droplet
    destroy     Destroy a selected droplet
  
  aws           Manage AWS EC2 instances
    create      Launch a new EC2 instance with cloud-init
    list        List existing EC2 instances
    ssh         SSH into a selected EC2 instance
    terminate   Terminate a selected EC2 instance
  
  azure         Manage Azure VMs
    create      Deploy a new Azure VM with cloud-init
    list        List existing Azure VMs
    ssh         SSH into a selected Azure VM
    delete      Delete a selected Azure VM
  
  gcp           Manage Google Cloud Compute Engine instances
    create      Launch a new GCE instance with cloud-init
    list        List existing GCE instances
    ssh         SSH into a selected GCE instance
    delete      Delete a selected GCE instance
  
  openstack     Manage OpenStack instances
    create      Create a new OpenStack instance with cloud-init
    list        List existing OpenStack instances
    ssh         SSH into a selected OpenStack instance
    delete      Delete a selected OpenStack instance

Options:
  -h, --help    Show help information
  -v, --version Show version information

Examples:
  # DigitalOcean
  $0 digitalocean create --repo GistWiz/cli --name QuickSearch --tags custom-tag1,custom-tag2
  $0 digitalocean list
  $0 digitalocean ssh
  $0 digitalocean destroy

  # AWS EC2
  $0 aws create --ami ami-12345678 --instance-type t2.micro --key-name my-key
  $0 aws list
  $0 aws ssh --instance-id i-0abcd1234
  $0 aws terminate --instance-id i-0abcd1234

  # Azure VMs
  $0 azure create --image UbuntuLTS --size Standard_B1s --name my-vm
  $0 azure list
  $0 azure ssh --name my-vm
  $0 azure delete --name my-vm

  # Google Cloud Compute Engine
  $0 gcp create --image-family ubuntu-2004-lts --machine-type e2-micro --name my-instance
  $0 gcp list
  $0 gcp ssh --name my-instance
  $0 gcp delete --name my-instance

  # OpenStack
  $0 openstack create --flavor m1.small --image UbuntuLTS --name my-instance
  $0 openstack list
  $0 openstack ssh --name my-instance
  $0 openstack delete --name my-instance
```
