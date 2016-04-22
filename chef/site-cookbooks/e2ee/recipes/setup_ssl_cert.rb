#
# Cookbook Name:: e2ee
# Recipe:: setup_keys
#	Sets up SSL certificate for http use with E2EE Server 
#
# Copyright 2016, XLAB
#
# All rights reserved - Do Not Redistribute
#

require 'openssl'
require 'rubygems'

cert_path = "#{node['e2ee']['https']['cert_path']}"
cert_name = "#{node['e2ee']['https']['cert_name']}"

openssl_x509 "#{cert_path}/#{cert_name}" do
  common_name 'E2EE-Server'
  org 'xlab-si'
  org_unit 'SPECS'
  country 'SI'
end

