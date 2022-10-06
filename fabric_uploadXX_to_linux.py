

from fabric import Connection
import os

def scaner_file (window_path):
    path_list = []
    file = os.listdir(window_path)
    for f in file:
        read_path = os.path.join(window_path,f)
        path_list.append(read_path)
    return path_list


# 1. 遍历目录下的所有文件
# 2. 再遍历上传

def from_Windows_upload_to_linux(window_path,username,Linux_ip,password,Linux_path):

    c = Connection(host=f'{username}@{Linux_ip}',connect_kwargs={'password': password})
    path_list = scaner_file(window_path)
    for one_file in path_list:
        c.put(one_file, remote=Linux_path)




if __name__ == "__main__":
    username = "root"
    password = "123456"
    Linux_ip = "192.168.56.129"
    window_path = r'C:\vue_t'
    Linux_path = r'/mnt'
    from_Windows_upload_to_linux(window_path, username, Linux_ip, password, Linux_path)