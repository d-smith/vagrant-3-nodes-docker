# -*- mode: ruby -*-
# vi: set ft=ruby :
#
#
#


# Specify Vagrant version, Vagrant API version, and Vagrant clone location
Vagrant.require_version '>= 1.6.0'
VAGRANTFILE_API_VERSION = '2'

require 'yaml'
require 'fileutils'
require 'erb'

servers = YAML.load_file(File.join(File.dirname(__FILE__), 'servers.yaml'))

# Create and configure the VMs
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.provider "virtualbox" do |v|
     v.memory = 1024
  end

  config.ssh.insert_key = false
  
  config.hostmanager.enabled = true
  config.hostmanager.manage_guest = true

  servers.each do |server|

    config.vm.define server['name'] do |srv|
      srv.vm.box_check_update = false
      srv.vm.hostname = server['name']
      srv.vm.box = server['box']

      if ENV['http_proxy'].nil? == false
        puts 'Setting http proxy'
        config.proxy.http = ENV['http_proxy']
        config.proxy.https = ENV['http_proxy']
        config.proxy.no_proxy="localhost," + server['priv_ip'] + ",/var/run/docker.sock"
      else
        puts 'Environment does not specify http_proxy value...'
      end

      srv.vm.network 'private_network', ip: server['priv_ip']
      srv.vm.synced_folder '.', '/vagrant'

      srv.vm.provision "shell", path: "provision.sh"
      srv.vm.provision :reload
    end
  end
end
