#!/bin/sh

if [ ! -e /tmp/dnsip.txt ] ; then
    ping mydomain.f3322.org -c1 | grep PING | awk '{ print $3 }' | sed 's/[()]//g' > /tmp/dnsip.txt
fi

mydnsip=$(head -1 /tmp/dnsip.txt)

curl -s http://icanhazip.com > /tmp/ip.txt
myip=$(head -1 /tmp/ip.txt)
echo "current IP:"$myip

if [ "$mydnsip" = "$myip" ] ; then
    echo 'IP is not change'
else
    http_code=`curl -o /dev/null -s -w %{http_code} --basic -u username:password "http://members.3322.net/dyndns/update?myip=$myip&hostname=mydomain.f3322.org"`
    if [ $http_code -eq 200 ] ; then
        echo 'update DDNS success'
        echo $myip > /tmp/dnsip.txt
    else
        echo 'update DDNS fail:$http_code'
    fi
fi
