#!/bin/bash

echo "Checking if SSH key exists..."
if test -f ./ssh/id_ed25519; then
    echo "SSH key exists!"
fi
if ! test -f ./ssh/id_ed25519; then
    echo "Creating SSH key..."
    mkdir -p ./ssh
    ssh-keygen -t ed25519 -f ./ssh/id_ed25519 -q -N ""
    ssh-add  ./ssh/id_ed25519
    echo "SSH key created!"
fi

read -p "User? " user
read -p "Target? " target
read -p "Password? " password

echo "export USER=${user}" >> ~/.profile
echo "export TARGET=${target}" >> ~/.profile
echo "export PASSWORD=${password}" >> ~/.profile

source ~/.profile

sshpass -p $PASSWORD ssh-copy-id $USER@$TARGET

echo 'alias ssht="ssh $user@$target"' >> ~/.profile

source ~/.profile
