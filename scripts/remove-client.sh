#!/bin/bash

# This option could be documented a bit better and maybe even be simplified
# ...but what can I say, I want some sleep too

if [[ "$#" -ne 1 ]]
then 
    echo "Provide the name of the client to be removed" >&2
    exit 1
fi

number_of_clients=$(tail -n +2 /etc/openvpn/server/easy-rsa/pki/index.txt | grep -c "^V")
if [[ "$number_of_clients" = 0 ]]; then
    echo
    echo "There are no existing clients!" >&2
    exit
fi
echo
# ------------------------------------------------------------------------------------------------------
# echo "Select the client to revoke:"
# tail -n +2 /etc/openvpn/server/easy-rsa/pki/index.txt | grep "^V" | cut -d '=' -f 2 | nl -s ') '
# client_number="$1"
# if [[ "$client_number" =~ ^[0-9]+$ && "$client_number" -le "$number_of_clients" ]]
# then
#     echo "$client_number: invalid selection."
#     exit 2
# fi
# -------------------------------------------------------------------------------------------------------
client="$1"
# client=$(tail -n +2 /etc/openvpn/server/easy-rsa/pki/index.txt | grep "^V" | cut -d '=' -f 2 | sed -n "$client_number"p)
echo

n=$(tail -n +2 /etc/openvpn/server/easy-rsa/pki/index.txt | grep "^V" | cut -d '=' -f 2 | grep "$client" | wc -l)
if [[ "$n" -ne 1 ]]
then
    echo "Invalid client name" >&2
    exit 2
fi


# revoke a client
cd /etc/openvpn/server/easy-rsa/
./easyrsa --batch revoke "$client"
EASYRSA_CRL_DAYS=3650 ./easyrsa gen-crl
rm -f /etc/openvpn/server/crl.pem
cp /etc/openvpn/server/easy-rsa/pki/crl.pem /etc/openvpn/server/crl.pem
# CRL is read with each client connection, when OpenVPN is dropped to nobody
chown nobody:"$group_name" /etc/openvpn/server/crl.pem
echo
echo "$client revoked!" >&1

exit