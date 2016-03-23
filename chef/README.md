## Requirements
* Vagrant
* Chef Development Kit
* librarian-chef (`gem install librarian-chef`)

## Running E2EE Server with Vagrant
From chef/ directory, run
```bash
$ librarian-chef install
```
Which will fetch dependencies and download them to the cookbooks/ folder. This command should also be run if you apply any changes to recipes in site-cookbooks/ folder.

To start chef-solo provisioner with vagrant, run
```bash
$ vagrant up
```

or, to reload configuration when modifying recipes or roles run
```bash
$ vagrant reload --provision
```

## Content
### Roles
This Chef repository includes a single role, _e2ee-server_, that automatically configures postgres, redis, RSA keys and paths for use with E2EE server.

You should change absolute paths included in attribute values according to your preferred configuration. 

### Recipes
*  _default.rb_ - Installs, tests and runs E2EE server
* _database.rb_ - Sets up and configures postgres backend for E2EE server
* _setup_keys.rb_ - Sets up RSA keypair for E2EE server and configures paths to keys in configuration files
* _setup_config.rb_ - Sets up config.json for E2EE server based on provided attributes

### Templates
This cookbook contains a single template file, that will replace the content of configuration file _config.json_ from e2ee-server repository.

## Issues
- Sometimes the first time Chef-solo provisioner is run with Vagrant, it fails at installation of E2EE server with "Errno::ENOENT: No such file or directory - go". The second time chef provisioning is run however, Chef run succeeds.
