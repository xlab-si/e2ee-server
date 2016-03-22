#
# Cookbook Name:: e2ee
# Recipe:: default
#	Installs, tests and runs E2EE server
#
# Copyright 2016, XLAB
#
# All rights reserved - Do Not Redistribute
#

gopath = node['go']['gopath']
e2ee_path = "#{gopath}/src/github.com/xlab-si/e2ee-server"

# Modify paths to JSON configuration files in settings.go 
template "#{e2ee_path}/settings/settings.go" do
  source 'settings.erb'
  variables ({
    :settings_path => "#{e2ee_path}/settings"
  })
end

# Install E2EE server
execute "Installing E2EE Server" do
	cwd "#{e2ee_path}"
	command "go install"
end

# Run tests
execute "Running unit tests" do
	cwd "#{e2ee_path}/tests/unit_tests"
	command "go test"
end

execute "Running API tests" do
	cwd "#{e2ee_path}/tests/api_tests"
	command "go test"
	ignore_failure true # odstrani, ko bo pomergano spet
end

# Run E2EE Server in background, redirect ouptut to file 
service "e2ee-server" do
	start_command "cd && #{node['go']['gobin']}/e2ee-server > #{node['e2ee']['log_file']} &"
	action :start
end