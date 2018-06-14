
## iOS restriction password recoverer

[![Build Status](https://travis-ci.com/brunomvsouza/iosrc.svg?branch=master)](https://travis-ci.com/brunomvsouza/iosrc) [![codecov](https://codecov.io/gh/brunomvsouza/iosrc/branch/master/graph/badge.svg)](https://codecov.io/gh/brunomvsouza/iosrc) [![Maintainability](https://api.codeclimate.com/v1/badges/4baddf0f6bb3c94f13e3/maintainability)](https://codeclimate.com/github/brunomvsouza/iosrc/maintainability)




### Step 1: Create a backup

Make sure you have an **unencrypted** backup of you device.

### Step 2: Find restriction `key` and `salt` on the created backup

You'll be able to find the values for `key` and `salt` on the file below
```
~/Library/Application\ Support/MobileSync/Backup/<BACKUP_HASH>/398bc9c2aeeab4cb0c12ada0f52eea12cf14f40b
```

It will look something like:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>RestrictionsPasswordKey</key>
	<data>
	aI9gkP3NEG+LLL8UPyAT1ehRd7g=
	</data>
	<key>RestrictionsPasswordSalt</key>
	<data>
	hpOLll==
	</data>
</dict>
</plist>
```

### Step 3: Execute the password cracker

```bash
# example for the example file below would be
./cli -key aI9gkP3NEG+LLL8UPyAT1ehRd7g= -salt hpOLll==
```

### Disclaimer

This repo is 100% inspired on http://ios7hash.derson.us. It was created because
I needed to recover an iPad restriction password and the implementation of the
original author wasn't fast enough (even for this small key space) due to running
on browsers.
