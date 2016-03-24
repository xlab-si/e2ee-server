#
# Cookbook Name:: e2ee
# Recipe:: database
# 	This recipe is a wrapper for setting up postgres backend for E2EE server
#
# Copyright 2016, XLAB
#
# All rights reserved - Do Not Redistribute
#

include_recipe "postgresql::server"
include_recipe "database::postgresql"

postgres_conn = {
  :host => node['postgresql']['config']['listen_addresses'],
  :port => 5432,
  :username => 'postgres',
  :password => node['postgresql']['password']['postgres']
}

# Create a named database
postgresql_database 'e2ee' do
  connection postgres_conn
  action :create
end

