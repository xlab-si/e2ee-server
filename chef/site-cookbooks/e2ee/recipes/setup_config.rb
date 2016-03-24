#
# Cookbook Name:: e2ee
# Recipe:: default
#	Sets up config.json for E2EE server based on attributes provided in e2ee-server role
#
# Copyright 2016, XLAB
#
# All rights reserved - Do Not Redistribute
#

gopath = node['go']['gopath']
e2ee_path = "#{gopath}/src/github.com/xlab-si/e2ee-server"

# Modify config with desired backend settings
template "#{e2ee_path}/config.json" do
  source 'config.erb'
  variables ({
	:private_key_path => "#{node['e2ee']['private_key']}",
    :public_key_path => "#{node['e2ee']['public_key']}",
    :redis_pw => "#{node['redisio']['default_settings']['requirepass']}",
	:db_host => "#{node['postgresql']['config']['listen_addresses']}",
	:db_pw => "#{node['postgresql']['password']['postgres']}"
  })
end