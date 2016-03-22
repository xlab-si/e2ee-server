#
# Cookbook Name:: e2ee
# Recipe:: setup_keys
#	Sets up RSA keypair for E2EE Server and configures paths to keys in configuration files
#
# Copyright 2016, XLAB
#
# All rights reserved - Do Not Redistribute
#

require 'openssl'

gopath = node['go']['gopath']
e2ee_path = "#{gopath}/src/github.com/xlab-si/e2ee-server"

rsa = OpenSSL::PKey::RSA.new(2048)
private_key = rsa.to_pem
public_key = rsa.public_key.to_pem

private_key_path = "#{node['e2ee']['private_key']}"
public_key_path = "#{node['e2ee']['public_key']}"

file "#{private_key_path}" do
	content "#{private_key}"
end

file "#{public_key_path}" do
	content "#{public_key}"
end

# Set paths according to where certificate and key are stored
# Replace existing paths in environment configuration files
template "#{e2ee_path}/settings/pre.json" do
  source 'pre.erb'
  variables ({
    :private_key_path => "#{private_key_path}",
    :public_key_path => "#{public_key_path}"
  })
end

template "#{e2ee_path}/settings/prod.json" do
  source 'prod.erb'
  variables ({
    :private_key_path => "#{private_key_path}",
    :public_key_path => "#{public_key_path}"
  })
end

template "#{e2ee_path}/settings/tests.json" do
  source 'tests.erb'
  variables ({
    :private_key_path => "#{private_key_path}",
    :public_key_path => "#{public_key_path}"
  })
end