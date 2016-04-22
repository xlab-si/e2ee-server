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
template "#{e2ee_path}/config/config.json" do
  source 'config.erb'
  variables ({
	:db_host => "#{node['postgresql']['config']['listen_addresses']}",
	:db_pw => "#{node['postgresql']['password']['postgres']}",
	:https_port => "#{node['e2ee']['https']['port']}",
	:https_cert_path => "#{node['e2ee']['https']['cert_path']}",
	:https_cert_prefix => "#{node['e2ee']['https']['cert_name']}"
  })
end