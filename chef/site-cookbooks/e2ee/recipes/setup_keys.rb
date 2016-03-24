#
# Cookbook Name:: e2ee
# Recipe:: setup_keys
#	Sets up RSA keypair for E2EE Server 
#
# Copyright 2016, XLAB
#
# All rights reserved - Do Not Redistribute
#

require 'openssl'

gopath = node['go']['gopath']
e2ee_path = "#{gopath}/src/github.com/xlab-si/e2ee-server"

# Create new RSA keypair for server
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
