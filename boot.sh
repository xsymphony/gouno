#!/bin/sh

fc-cache -fv
circusd /etc/circus/circus.ini
#/usr/bin/supervisord -c /etc/supervisor/conf.d/supervisord.conf