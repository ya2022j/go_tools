expect -c "set timeout 2 ; spawn ssh root@192.168.0.123 ; expect yes { send 123\r } password { send 123\r } ; exit"


#  interact 停留在命令行状态
# exit 0
# 可以写一个完全裸机的状态，全自动安装python,docker,mysql等工具的脚本。活学活用啊
#!/bin/sh
timeout=10
password="your_password"
command="sudo service network-manager restart"
expect -c "
    set timeout ${timeout}
    spawn ${command}
    expect \"sudo\"
    send \"${password}\n\"
    expect \"$\"
    exit 0
    interact
    e
"
exit 0
