#!/bin/sh

# 加载字体
fc-cache -fv

/usr/local/bin/supervisord -c /etc/supervisor/conf.d/supervisord.conf