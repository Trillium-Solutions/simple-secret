#!/bin/bash


# Test read entire file.
cat <<EOF
### Test reading entire vault. Output should be:
===
hello: world
a: 1
b: 2
===
EOF
./simple-secret -passwordFile ./testdata/password_file.txt -vaultFile ./testdata/password_is_secret.yml -view
echo ===
echo

# Test getting keys.
cat <<EOF

### Test getting keys. Output should be:
===
hello: world
a: 1
b: 2
===
EOF
./simple-secret -passwordFile ./testdata/password_file.txt -vaultFile ./testdata/password_is_secret.yml -get a
./simple-secret -passwordFile ./testdata/password_file.txt -vaultFile ./testdata/password_is_secret.yml -get hello
echo ===
echo

# Test modifiying a file.
cat <<EOF
### Test modifying a vault. Output should be:
===
a: "1"
b: "2"
c: "3"
hello: world
===
EOF
cp ./testdata/password_is_secret.yml ./testdata/password_is_secret_modified.yml ;
./simple-secret -passwordFile ./testdata/password_file.txt -vaultFile ./testdata/password_is_secret_modified.yml -put c -putval 3
./simple-secret -passwordFile ./testdata/password_file.txt -vaultFile ./testdata/password_is_secret_modified.yml -view
echo ===
echo
