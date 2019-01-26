#!/bin/bash
# https://developer.godaddy.com/keys
KEY=""
SECRET=""

domain="example.com"
name="a-record"

# Get public IP from ipinfo.io
public_ip="$(curl --silent ipinfo.io/ip)"

# get current record
current_record="$(dig @8.8.8.8 +short $name.$domain)"

if [ "$current_record" != "$public_ip" ]; then
    echo New public IP is $public_ip
    # Update godaddy A record - https://developer.godaddy.com/doc/endpoint/domains#/v1/recordReplaceTypeName
    curl -X PUT "https://api.godaddy.com/v1/domains/$domain/records/A/$name" \
        -H "Authorization: sso-key $KEY:$SECRET" \
        -H 'Content-Type: application/json' \
        --data '[{"type": "A", "name": "'"$name"'", "data": "'"$public_ip"'", "ttl": 3600}]'
else
    echo Nothing changed
fi
