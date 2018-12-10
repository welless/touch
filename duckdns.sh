#!/bin/sh

#if [ ! -e /tmp/dnsip.txt ] ; then
    ping npe.duckdns.org -c1 | grep PING | awk '{ print $3 }' | sed 's/[()]//g' > /tmp/dnsip.txt
#fi

mydnsip=$(head -1 /tmp/dnsip.txt)
token=xxxx

curl -s http://icanhazip.com > /tmp/ip.txt
myip=$(head -1 /tmp/ip.txt)
#echo "current IP:"$myip

current=$(date "+%Y-%m-%d %H:%M:%S")

if [ "$mydnsip" = "$myip" ] ; then
    echo $current $myip 'DDNS IP is not change'
else
    http_code=$(curl -s -w "%{http_code}" "https://www.duckdns.org/update?domains=npe&token=$token&ip=$myip")
    if [ $http_code = 'OK200' ] ; then
        echo $current $mydnsip ' => ' $myip 'update DDNS success'
        echo $myip > /tmp/dnsip.txt
    else
        echo $current 'update DDNS fail:$http_code'
    fi
fi
