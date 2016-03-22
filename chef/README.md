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

## Content
### Roles
This Chef repository includes a single role, _e2ee-server_, that automatically configures postgres, redis, RSA keys and paths for use with E2EE server.

You should change absolute paths included in attribute values according to your preferred configuration.

### Recipes
*  _default.rb_ - Installs, tests and runs E2EE server
* _database.rb_ - Sets up and configures postgres backend for E2EE server
* _setup_keys.rb_ - Sets up RSA keypair for E2EE server and configures paths to keys in configuration files

### Templates
Templates include 4 .erb files, that will replace the content of configuration files from e2ee-server repository.
