name             'e2ee'
maintainer       'XLAB'
maintainer_email 'manca.bizjak@xlab.si'
license          'All rights reserved'
description      'Installs/Configures E2EE (End-to-End Encryption) Server'
long_description IO.read(File.join(File.dirname(__FILE__), 'README.md'))
version          '0.1.0'

# attribute -> the list of attributes that are required to configure a cookbook
depends 'database'
depends 'postgresql'
depends 'golang'
depends 'redisio'